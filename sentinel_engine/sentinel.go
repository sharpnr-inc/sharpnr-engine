package sentinelEngine

import (
	"context"
	"errors"

	"github.com/google/uuid"
)

// ErrAccessDenied is returned by Authorize when policy evaluation denies the
// request. Callers (controllers) match it with errors.Is to answer 403 instead
// of 500.
var ErrAccessDenied = errors.New("sentinel: access denied")

// SentinelEngineInterface is the authorization decision point injected into
// services (and, later, route middleware). It is safe for concurrent use; one
// instance is shared across the app.
type SentinelEngineInterface interface {
	// IsAuthorized evaluates the request against the org's enabled policies and
	// reports the decision. It errs only on infrastructure problems (store
	// failure, uncompilable stored policy) — a plain "no" is (false, nil).
	IsAuthorized(ctx context.Context, req AccessRequest) (bool, error)

	// Authorize is IsAuthorized for guard clauses: nil on allow, ErrAccessDenied
	// on deny, the underlying error otherwise.
	Authorize(ctx context.Context, req AccessRequest) error

	// InvalidateOrganization drops the org's compiled policy cache. The Sentinel
	// authoring service calls it after any policy write so changes take effect
	// immediately instead of after the cache TTL.
	InvalidateOrganization(orgID uuid.UUID)
}

// PolicyStore is the engine's read surface over persisted authorization data.
// The Sentinel repo implements it (see sentinel/repo/engine_store.go); the
// engine deliberately depends on this narrow interface rather than the full
// authoring repository.
type PolicyStore interface {
	// ListEnabledPolicies returns every enabled, live policy of an org — the
	// engine compiles them all; attachment tables only classify policies for
	// the UI (a role-scoped policy already names its role in the Cedar text).
	ListEnabledPolicies(ctx context.Context, orgID uuid.UUID) ([]PolicyRecord, error)

	// GetPrincipalParents returns the role and group names a user belongs to
	// within an org — the parent edges of the Cedar principal.
	GetPrincipalParents(ctx context.Context, orgID, userID uuid.UUID) (PrincipalParents, error)
}
