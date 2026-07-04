package orchestrator_service

import (
	"context"
	"errors"

	sentinelEngine "sharpnr-engine/sentinel_engine"

	"github.com/google/uuid"
)

// authorize is the orchestrator's guard clause: may emplID perform action on
// the given resource inside orgID? resourceID is empty for type-level checks
// (create/list). Returns sentinelEngine.ErrAccessDenied on a deny, which the
// controllers translate to 403.
func (s *orchestratorService) authorize(ctx context.Context, orgID, emplID uuid.UUID, action, resourceType, resourceID string) error {
	if s.sentinelEngine == nil {
		// Fail closed: a missing engine must never mean "everything allowed".
		return errors.New("sentinel engine is not wired into the orchestrator")
	}
	return s.sentinelEngine.Authorize(ctx, sentinelEngine.AccessRequest{
		OrgID:  orgID,
		UserID: emplID,
		Action: action,
		Resource: sentinelEngine.Resource{
			Type: resourceType,
			ID:   resourceID,
		},
	})
}
