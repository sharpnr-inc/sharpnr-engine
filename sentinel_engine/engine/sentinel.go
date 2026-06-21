package sentinelengine

import (
	"context"
	"fmt"
	"sync"

	"github.com/cedar-policy/cedar-go"
	"github.com/cedar-policy/cedar-go/types"
	"github.com/google/uuid"

	sentinel_actions "sharpnr-engine/sentinel_engine/actions"
	sentinelengine_models "sharpnr-engine/sentinel_engine/models"
)

// Sentinel is the authorization engine. It uses cedar-go (v1.8.0) and follows
// the model at https://docs.cedarpolicy.com/.
//
// Two things move at different speeds, so they're handled differently:
//   - Policies change rarely -> compiled into a per-org *cedar.PolicySet and cached.
//   - Membership (who is in which role/group) changes constantly -> the EntityMap
//     is rebuilt from the database on every authorization call.

// Cedar entity type names. These must match the types referenced in stored
// policies (e.g. `principal in Role::"owner"`).
const (
	entityOrganization = "Organization"
	entityUser         = "User"
	entityRole         = "Role"
	entityGroup        = "Group"
)

// GroupMembership is one node in a user's group closure: the group, its parent
// group (empty for a top-level group), and whether the user is a *direct*
// member. Only direct memberships make the user a member; ancestor groups are
// reached transitively via the ParentSlug edge.
type GroupMembership struct {
	Slug       string
	ParentSlug string
	Direct     bool
}

// Repository is the slice of the data layer the engine needs. Your
// sentinel_repo.SentinelRepository should satisfy this directly or via a thin
// adapter. Defining it here (rather than importing the full repo interface)
// keeps the engine decoupled and easy to test with a fake.
type Repository interface {
	// ListEnabledPolicies returns every enabled policy row for an org.
	ListEnabledPolicies(ctx context.Context, orgID uuid.UUID) ([]*sentinelengine_models.SentinelPolicy, error)
	// UserRoleSlugs returns the slugs of the roles the user is directly assigned.
	UserRoleSlugs(ctx context.Context, orgID, userID uuid.UUID) ([]string, error)
	// UserGroupClosure returns the user's direct groups plus all ancestor groups
	// (typically via a recursive CTE over groups.parent_group_id).
	UserGroupClosure(ctx context.Context, orgID, userID uuid.UUID) ([]GroupMembership, error)
}

// Sentinel is the engine. Safe for concurrent use.
type Sentinel struct {
	repo Repository

	mu    sync.RWMutex
	cache map[uuid.UUID]*cedar.PolicySet // compiled policies, keyed by org
}

// New constructs a Sentinel engine.
func New(repo Repository) *Sentinel {
	return &Sentinel{
		repo:  repo,
		cache: make(map[uuid.UUID]*cedar.PolicySet),
	}
}

// Resource identifies the thing being acted upon.
type Resource struct {
	Type    string                 // Cedar entity type, e.g. "Document"
	ID      string                 // entity id
	Attrs   map[string]types.Value // optional attributes; orgId is added automatically
	Parents []types.EntityUID      // optional hierarchy, e.g. a Document in a Folder
}

// AuthorizeInput is everything needed to make one decision.
type AuthorizeInput struct {
	OrgID    uuid.UUID
	UserID   uuid.UUID
	Action   string // a sentinel_actions constant, e.g. sentinel_actions.ActionEditDocument
	Resource Resource
	Context  map[string]types.Value // optional request context
}

// Result is the outcome of a check.
type Result struct {
	Allowed bool
	Reasons []string // ids of the policies that determined the decision
	Errors  []string // evaluation errors, if any (the decision is still authoritative)
}

// IsAuthorized answers a single request: can this user perform this action on
// this resource, within this org? It loads (or reuses) the org's compiled
// policies, builds the principal+resource entity graph from current membership,
// and evaluates. On any infrastructure error it returns (deny, err) — fail closed.
func (s *Sentinel) IsAuthorized(ctx context.Context, in AuthorizeInput) (Result, error) {
	ps, err := s.policySetFor(ctx, in.OrgID)
	if err != nil {
		return Result{Allowed: false}, fmt.Errorf("load policies: %w", err)
	}

	entities, err := s.buildEntities(ctx, in)
	if err != nil {
		return Result{Allowed: false}, fmt.Errorf("build entities: %w", err)
	}

	req := cedar.Request{
		Principal: cedar.NewEntityUID(entityUser, types.String(in.UserID.String())),
		Action:    sentinel_actions.Action(in.Action),
		Resource:  cedar.NewEntityUID(types.EntityType(in.Resource.Type), types.String(in.Resource.ID)),
		Context:   cedar.NewRecord(toRecordMap(in.Context)),
	}

	decision, diag := cedar.Authorize(ps, entities, req)

	res := Result{Allowed: decision == cedar.Allow}
	for _, r := range diag.Reasons {
		res.Reasons = append(res.Reasons, string(r.PolicyID))
	}
	for _, de := range diag.Errors {
		res.Errors = append(res.Errors, fmt.Sprintf("%s: %s", de.PolicyID, de.Message))
	}
	return res, nil
}

// InvalidateOrg drops an org's compiled PolicySet so the next authorization
// recompiles from the database. Call this after any create/update/delete of an
// org's policies. Role/group *membership* changes do NOT need this — entities
// are rebuilt on every request.
func (s *Sentinel) InvalidateOrg(orgID uuid.UUID) {
	s.mu.Lock()
	delete(s.cache, orgID)
	s.mu.Unlock()
}

// policySetFor returns the cached compiled set for an org, compiling on a miss.
func (s *Sentinel) policySetFor(ctx context.Context, orgID uuid.UUID) (*cedar.PolicySet, error) {
	s.mu.RLock()
	ps, ok := s.cache[orgID]
	s.mu.RUnlock()
	if ok {
		return ps, nil
	}

	ps, err := s.compile(ctx, orgID)
	if err != nil {
		return nil, err
	}

	s.mu.Lock()
	s.cache[orgID] = ps
	s.mu.Unlock()
	return ps, nil
}

// compile turns an org's stored policy rows into a *cedar.PolicySet. The row's
// ID is used as the Cedar PolicyID, so diagnostics point back at your table.
func (s *Sentinel) compile(ctx context.Context, orgID uuid.UUID) (*cedar.PolicySet, error) {
	rows, err := s.repo.ListEnabledPolicies(ctx, orgID)
	if err != nil {
		return nil, err
	}

	ps := cedar.NewPolicySet()
	for _, row := range rows {
		var p cedar.Policy
		if err := p.UnmarshalCedar([]byte(row.Policy)); err != nil {
			// Fail closed: a malformed stored policy is an integrity error.
			// Policies are validated on write, so this should not happen in
			// practice. Returning here denies the whole org rather than
			// silently dropping one policy and granting more than intended.
			return nil, fmt.Errorf("policy %s failed to parse: %w", row.ID, err)
		}
		ps.Add(cedar.PolicyID(row.ID.String()), &p)
	}
	return ps, nil
}

// buildEntities assembles the principal's membership graph and the target
// resource into an EntityMap. Both principal and resource carry an orgId
// attribute so the `resource.orgId == principal.orgId` tenant guard holds.
func (s *Sentinel) buildEntities(ctx context.Context, in AuthorizeInput) (types.EntityMap, error) {
	em := types.EntityMap{}

	orgID := in.OrgID.String()
	orgUID := cedar.NewEntityUID(entityOrganization, types.String(orgID))
	em[orgUID] = types.Entity{UID: orgUID}

	roleSlugs, err := s.repo.UserRoleSlugs(ctx, in.OrgID, in.UserID)
	if err != nil {
		return nil, fmt.Errorf("load roles: %w", err)
	}
	groups, err := s.repo.UserGroupClosure(ctx, in.OrgID, in.UserID)
	if err != nil {
		return nil, fmt.Errorf("load groups: %w", err)
	}

	userParents := []types.EntityUID{orgUID}

	for _, slug := range roleSlugs {
		ruid := cedar.NewEntityUID(entityRole, types.String(slug))
		em[ruid] = types.Entity{UID: ruid, Parents: cedar.NewEntityUIDSet(orgUID)}
		userParents = append(userParents, ruid)
	}

	for _, g := range groups {
		guid := cedar.NewEntityUID(entityGroup, types.String(g.Slug))
		parents := []types.EntityUID{orgUID}
		if g.ParentSlug != "" {
			parents = append(parents, cedar.NewEntityUID(entityGroup, types.String(g.ParentSlug)))
		}
		em[guid] = types.Entity{UID: guid, Parents: cedar.NewEntityUIDSet(parents...)}
		if g.Direct {
			userParents = append(userParents, guid)
		}
	}

	// Principal (the employee).
	userUID := cedar.NewEntityUID(entityUser, types.String(in.UserID.String()))
	em[userUID] = types.Entity{
		UID:     userUID,
		Parents: cedar.NewEntityUIDSet(userParents...),
		Attributes: cedar.NewRecord(cedar.RecordMap{
			"orgId": types.String(orgID),
		}),
	}

	// Resource, with orgId merged into any caller-supplied attributes.
	resAttrs := cedar.RecordMap{"orgId": types.String(orgID)}
	for k, v := range in.Resource.Attrs {
		resAttrs[types.String(k)] = v
	}
	resUID := cedar.NewEntityUID(types.EntityType(in.Resource.Type), types.String(in.Resource.ID))
	em[resUID] = types.Entity{
		UID:        resUID,
		Parents:    cedar.NewEntityUIDSet(in.Resource.Parents...),
		Attributes: cedar.NewRecord(resAttrs),
	}

	return em, nil
}

func toRecordMap(m map[string]types.Value) cedar.RecordMap {
	rm := cedar.RecordMap{}
	for k, v := range m {
		rm[types.String(k)] = v
	}
	return rm
}
