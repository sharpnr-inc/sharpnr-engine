// Package orchestrator_authz holds the orchestrator's authorization guard.
//
// It lives in its own package (rather than beside the orchestrators) so that
// per-domain orchestrator sub-packages — orchestrator/service/property and the
// domain packages under it — can share one guard without importing
// orchestrator_service, which would be an import cycle.
package orchestrator_authz

import (
	"context"
	"errors"

	sentinelEngine "sharpnr-engine/sentinel_engine"

	"github.com/google/uuid"
)

// Guard is the orchestrator's guard clause, shared by every domain
// orchestrator: may emplID perform action on the given resource inside orgID?
// resourceID is empty for type-level checks (create/list). Returns
// sentinelEngine.ErrAccessDenied on a deny, which the controllers translate
// to 403.
type Guard struct {
	engine sentinelEngine.SentinelEngineInterface
}

// New builds the single Guard shared by all domain orchestrators.
func New(engine sentinelEngine.SentinelEngineInterface) *Guard {
	return &Guard{engine: engine}
}

func (g *Guard) Authorize(ctx context.Context, orgID, emplID uuid.UUID, action, resourceType, resourceID string) error {
	if g == nil || g.engine == nil {
		// Fail closed: a missing engine must never mean "everything allowed".
		return errors.New("sentinel engine is not wired into the orchestrator")
	}
	return g.engine.Authorize(ctx, sentinelEngine.AccessRequest{
		OrgID:  orgID,
		UserID: emplID,
		Action: action,
		Resource: sentinelEngine.Resource{
			Type: resourceType,
			ID:   resourceID,
		},
	})
}
