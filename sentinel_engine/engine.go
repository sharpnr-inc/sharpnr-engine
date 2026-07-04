package sentinelEngine

import (
	"context"
	"errors"
	"fmt"
	"maps"
	"sync"
	"time"

	cedar "github.com/cedar-policy/cedar-go"
	"github.com/google/uuid"
)

// Cedar entity types the engine emits.
//
// Policy applicability is ATTACHMENT-driven: a policy applies to a user because
// it is org-scoped or attached to the user / one of their roles / groups — not
// because its text names a role or group. Policy text therefore targets the
// generic principal (`permit (principal, action, resource)`), and the engine
// builds no Role/Group parent entities.
const (
	EntityTypeUser cedar.EntityType = "User"

	entityTypeAction cedar.EntityType = "Action"

	// anyResourceID stands in for "no particular instance" on type-level checks
	// (create/list). Policies scoped `resource is <Type>` still match it.
	anyResourceID = "*"
)

// policyCacheTTL bounds how stale a user's compiled policy set can get for
// changes that bypass InvalidateOrganization (e.g. role/group membership edits
// or manual DB changes).
const policyCacheTTL = time.Minute

type compiledPolicies struct {
	set      *cedar.PolicySet
	loadedAt time.Time
}

// SentinelEngine evaluates Cedar policies stored by the Sentinel module. A
// user's effective policy set is the union of the org-scoped policies and the
// policies attached to them (directly, via roles, via groups) — fetched in one
// store call and compiled once per (org, user), with a short TTL.
type SentinelEngine struct {
	store PolicyStore

	mu sync.RWMutex
	// compiled policy sets keyed org -> user, so InvalidateOrganization can
	// drop everything an org's policy writes may have affected.
	cache map[uuid.UUID]map[uuid.UUID]compiledPolicies
}

// Compile-time guard: the engine must satisfy its public interface.
var _ SentinelEngineInterface = (*SentinelEngine)(nil)

func NewSentinelEngine(store PolicyStore) *SentinelEngine {
	return &SentinelEngine{
		store: store,
		cache: make(map[uuid.UUID]map[uuid.UUID]compiledPolicies),
	}
}

// IsAuthorized implements [SentinelEngineInterface].
func (e *SentinelEngine) IsAuthorized(ctx context.Context, req AccessRequest) (bool, error) {
	if req.OrgID == uuid.Nil || req.UserID == uuid.Nil {
		return false, errors.New("sentinel: request is missing orgID or userID")
	}
	if req.Action == "" || req.Resource.Type == "" {
		return false, errors.New("sentinel: request is missing action or resource type")
	}

	// Multi-tenant boundary: a resource known to belong to another org is
	// denied before any policy runs. Policies can only ever widen access
	// WITHIN the request's org — never across it.
	if req.Resource.OrgID != uuid.Nil && req.Resource.OrgID != req.OrgID {
		return false, nil
	}

	policies, err := e.policiesForUser(ctx, req.OrgID, req.UserID)
	if err != nil {
		return false, err
	}

	principal := cedar.NewEntityUID(EntityTypeUser, cedar.String(req.UserID.String()))
	action := cedar.NewEntityUID(entityTypeAction, cedar.String(req.Action))
	resourceID := req.Resource.ID
	if resourceID == "" {
		resourceID = anyResourceID
	}
	resource := cedar.NewEntityUID(cedar.EntityType(req.Resource.Type), cedar.String(resourceID))

	decision, _ := cedar.Authorize(
		policies.set,
		buildEntities(principal, resource, req),
		cedar.Request{
			Principal: principal,
			Action:    action,
			Resource:  resource,
			Context:   toCedarRecord(req.Context),
		},
	)
	return decision == cedar.Allow, nil
}

// Authorize implements [SentinelEngineInterface].
func (e *SentinelEngine) Authorize(ctx context.Context, req AccessRequest) error {
	allowed, err := e.IsAuthorized(ctx, req)
	if err != nil {
		return err
	}
	if !allowed {
		return fmt.Errorf("%w: user %s may not %s on %s", ErrAccessDenied, req.UserID, req.Action, req.Resource.Type)
	}
	return nil
}

// InvalidateOrganization implements [SentinelEngineInterface].
func (e *SentinelEngine) InvalidateOrganization(orgID uuid.UUID) {
	e.mu.Lock()
	delete(e.cache, orgID)
	e.mu.Unlock()
}

/* ---------------------------------- <***> --------------------------------- */

// policiesForUser returns the user's compiled effective policy set, recompiling
// from the store when the cache is cold or past its TTL.
func (e *SentinelEngine) policiesForUser(ctx context.Context, orgID, userID uuid.UUID) (compiledPolicies, error) {
	e.mu.RLock()
	cached, ok := e.cache[orgID][userID]
	e.mu.RUnlock()
	if ok && time.Since(cached.loadedAt) < policyCacheTTL {
		return cached, nil
	}

	records, err := e.store.GetAllPoliciesOfUser(ctx, orgID, userID)
	if err != nil {
		return compiledPolicies{}, fmt.Errorf("sentinel: loading policies of user %s: %w", userID, err)
	}

	set, err := compilePolicySet(records)
	if err != nil {
		return compiledPolicies{}, err
	}

	compiled := compiledPolicies{set: set, loadedAt: time.Now()}
	e.mu.Lock()
	if e.cache[orgID] == nil {
		e.cache[orgID] = make(map[uuid.UUID]compiledPolicies)
	}
	e.cache[orgID][userID] = compiled
	e.mu.Unlock()
	return compiled, nil
}

// compilePolicySet parses every stored policy text into one PolicySet. A policy
// that fails to parse fails the whole compile: silently skipping it could drop a
// forbid rule and grant more access than intended, so the engine fails closed.
// (Texts are validated at write time, so this signals data corruption.)
func compilePolicySet(records []PolicyRecord) (*cedar.PolicySet, error) {
	set := cedar.NewPolicySet()
	for _, rec := range records {
		list, err := cedar.NewPolicyListFromBytes(rec.ID.String()+".cedar", []byte(rec.Text))
		if err != nil {
			return nil, fmt.Errorf("sentinel: policy %s does not compile: %w", rec.ID, err)
		}
		// A single row may hold several statements; suffix the row id to keep
		// cedar.PolicyIDs unique while still tracing back to the row.
		for i, p := range list {
			set.Add(cedar.PolicyID(fmt.Sprintf("%s@%d", rec.ID, i)), p)
		}
	}
	return set, nil
}

// buildEntities assembles the entity graph for one evaluation: the principal
// and the resource with its attributes.
//
// The engine stamps the tenancy attributes itself so they are never missing or
// caller-forgeable: principal.orgId is the request's org, resource.orgId is the
// resource's owning org (falling back to the request's org when the caller
// didn't resolve it — the structural check above already handled the known-
// foreign case). Authored `when { resource.orgId == principal.orgId }` guards
// therefore evaluate correctly instead of erroring on a missing attribute.
func buildEntities(principal, resource cedar.EntityUID, req AccessRequest) cedar.EntityMap {
	resourceOrg := req.Resource.OrgID
	if resourceOrg == uuid.Nil {
		resourceOrg = req.OrgID
	}

	resourceAttrs := make(map[string]any, len(req.ResourceAttributes)+1)
	maps.Copy(resourceAttrs, req.ResourceAttributes)
	resourceAttrs["orgId"] = resourceOrg // engine-authoritative; overrides any caller value

	return cedar.EntityMap{
		principal: cedar.Entity{
			UID:        principal,
			Attributes: toCedarRecord(map[string]any{"orgId": req.OrgID}),
		},
		resource: cedar.Entity{
			UID:        resource,
			Attributes: toCedarRecord(resourceAttrs),
		},
	}
}

// toCedarRecord converts a plain Go map into a Cedar record, dropping values
// that have no sensible Cedar representation.
func toCedarRecord(m map[string]any) cedar.Record {
	if len(m) == 0 {
		return cedar.Record{}
	}
	rec := make(cedar.RecordMap, len(m))
	for k, v := range m {
		if cv, ok := toCedarValue(v); ok {
			rec[cedar.String(k)] = cv
		}
	}
	return cedar.NewRecord(rec)
}

func toCedarValue(v any) (cedar.Value, bool) {
	switch t := v.(type) {
	case bool:
		return cedar.Boolean(t), true
	case string:
		return cedar.String(t), true
	case int:
		return cedar.Long(t), true
	case int32:
		return cedar.Long(t), true
	case int64:
		return cedar.Long(t), true
	case float64:
		// JSON numbers arrive as float64; only integral values map cleanly.
		if t == float64(int64(t)) {
			return cedar.Long(int64(t)), true
		}
		return nil, false
	case uuid.UUID:
		return cedar.String(t.String()), true
	case time.Time:
		return cedar.NewDatetime(t), true
	case map[string]any:
		return toCedarRecord(t), true
	case []any:
		vals := make([]cedar.Value, 0, len(t))
		for _, item := range t {
			if cv, ok := toCedarValue(item); ok {
				vals = append(vals, cv)
			}
		}
		return cedar.NewSet(vals...), true
	case []string:
		vals := make([]cedar.Value, 0, len(t))
		for _, item := range t {
			vals = append(vals, cedar.String(item))
		}
		return cedar.NewSet(vals...), true
	default:
		return nil, false
	}
}
