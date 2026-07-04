package sentinelEngine

import (
	"context"
	"errors"
	"fmt"
	"sync"
	"time"

	cedar "github.com/cedar-policy/cedar-go"
	"github.com/google/uuid"
)

// Cedar entity types the engine emits. They must match what the frontend policy
// builder writes into policy text (see sentinel/dtos/principals.go).
const (
	EntityTypeUser  cedar.EntityType = "User"
	EntityTypeRole  cedar.EntityType = "Role"
	EntityTypeGroup cedar.EntityType = "Group"

	entityTypeAction cedar.EntityType = "Action"

	// anyResourceID stands in for "no particular instance" on type-level checks
	// (create/list). Policies scoped `resource is <Type>` still match it.
	anyResourceID = "*"
)

// policyCacheTTL bounds how stale a compiled policy set can get when a write
// slips past InvalidateOrganization (e.g. a manual DB edit).
const policyCacheTTL = time.Minute

type compiledPolicies struct {
	set      *cedar.PolicySet
	loadedAt time.Time
}

// SentinelEngine evaluates Cedar policies stored by the Sentinel module. It
// keeps one compiled PolicySet per organization; principal membership (roles,
// groups) is read fresh on every request so assignment changes apply instantly.
type SentinelEngine struct {
	store PolicyStore

	mu    sync.RWMutex
	cache map[uuid.UUID]compiledPolicies
}

// Compile-time guard: the engine must satisfy its public interface.
var _ SentinelEngineInterface = (*SentinelEngine)(nil)

func NewSentinelEngine(store PolicyStore) *SentinelEngine {
	return &SentinelEngine{
		store: store,
		cache: make(map[uuid.UUID]compiledPolicies),
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

	policies, err := e.policiesForOrg(ctx, req.OrgID)
	if err != nil {
		return false, err
	}

	parents, err := e.store.GetPrincipalParents(ctx, req.OrgID, req.UserID)
	if err != nil {
		return false, fmt.Errorf("sentinel: loading principal memberships: %w", err)
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
		buildEntities(principal, resource, parents, req.ResourceAttributes),
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

// policiesForOrg returns the org's compiled policy set, recompiling from the
// store when the cache is cold or past its TTL.
func (e *SentinelEngine) policiesForOrg(ctx context.Context, orgID uuid.UUID) (compiledPolicies, error) {
	e.mu.RLock()
	cached, ok := e.cache[orgID]
	e.mu.RUnlock()
	if ok && time.Since(cached.loadedAt) < policyCacheTTL {
		return cached, nil
	}

	records, err := e.store.ListEnabledPolicies(ctx, orgID)
	if err != nil {
		return compiledPolicies{}, fmt.Errorf("sentinel: loading policies: %w", err)
	}

	set, err := compilePolicySet(records)
	if err != nil {
		return compiledPolicies{}, err
	}

	compiled := compiledPolicies{set: set, loadedAt: time.Now()}
	e.mu.Lock()
	e.cache[orgID] = compiled
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
// with its role/group parent edges, those parents themselves, and the resource
// with its attributes.
func buildEntities(principal, resource cedar.EntityUID, parents PrincipalParents, resourceAttrs map[string]any) cedar.EntityMap {
	parentUIDs := make([]cedar.EntityUID, 0, len(parents.Roles)+len(parents.Groups))
	for _, role := range parents.Roles {
		parentUIDs = append(parentUIDs, cedar.NewEntityUID(EntityTypeRole, cedar.String(role)))
	}
	for _, group := range parents.Groups {
		parentUIDs = append(parentUIDs, cedar.NewEntityUID(EntityTypeGroup, cedar.String(group)))
	}

	entities := cedar.EntityMap{
		principal: cedar.Entity{
			UID:     principal,
			Parents: cedar.NewEntityUIDSet(parentUIDs...),
		},
		resource: cedar.Entity{
			UID:        resource,
			Attributes: toCedarRecord(resourceAttrs),
		},
	}
	for _, uid := range parentUIDs {
		entities[uid] = cedar.Entity{UID: uid}
	}
	return entities
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
