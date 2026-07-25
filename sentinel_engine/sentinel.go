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
	// IsAuthorized evaluates the request against the user's effective policies
	// and reports the decision. It errs only on infrastructure problems (store
	// failure, uncompilable stored policy) — a plain "no" is (false, nil).
	IsAuthorized(ctx context.Context, req AccessRequest) (bool, error)

	// Authorize is IsAuthorized for guard clauses: nil on allow, ErrAccessDenied
	// on deny, the underlying error otherwise.
	Authorize(ctx context.Context, req AccessRequest) error

	// InvalidateOrganization drops the org's compiled policy caches (every user
	// of that org). The Sentinel authoring service calls it after any policy or
	// attachment write so changes take effect immediately instead of after the
	// cache TTL.
	InvalidateOrganization(orgID uuid.UUID)
}

// PolicyStore is the engine's read surface over persisted authorization data.
// The Sentinel repo implements it (see sentinel/repo/engine_store.go); the
// engine deliberately depends on this narrow interface rather than the full
// authoring repository.
type PolicyStore interface {
	// GetAllPoliciesOfUser collects every enabled policy that applies to a user,
	// from all four sources:
	//   1. policies scoped to the organization (apply to everyone in the org)
	//   2. policies attached to the user's roles      (role_policies)
	//   3. policies attached to the user's groups     (groups_sentinel_policies)
	//   4. policies attached directly to the user     (user_policies)
	// The result is deduplicated (the same policy may arrive via several
	// attachments) and each record carries the row id so the engine can trace a
	// decision or compile failure back to the stored policy.
	GetAllPoliciesOfUser(ctx context.Context, orgID, userID uuid.UUID) ([]PolicyRecord, error)
}
