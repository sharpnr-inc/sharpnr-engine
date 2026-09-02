package auditEngineService

import (
	"context"
	"encoding/json"
	"errors"
	auditEngineDtos "sharpnr-engine/audit_log_engine/dto"
	auditEngineModels "sharpnr-engine/audit_log_engine/models"

	"github.com/google/uuid"
)

// Page-size policy for reading an org's history. Exported because the
// orchestrator clamps with the same numbers before it reports the page it
// served back to the frontend — one source of truth, so the pagination
// metadata can never describe a page different from the one the engine built.
const (
	DefaultPageSize = 50
	MaxPageSize     = 200
)

func (s *auditLogService) Record(
	ctx context.Context,
	input auditEngineDtos.RecordEventInput,
) error {
	// 1. Minimal validation — an audit entry without a tenant or an
	// action name is useless.
	if input.OrganizationID == uuid.Nil {
		return errors.New("audit log: organization id is required")
	}
	if input.Action == "" {
		return errors.New("audit log: action is required")
	}

	// 2. Default the actor type so callers only set it for system events
	actorType := input.ActorType
	if actorType == "" {
		actorType = auditEngineModels.ActorTypeEmployee
	}

	// 3. Serialize the optional metadata payload to JSONB
	var metadata []byte
	if len(input.Metadata) > 0 {
		encoded, err := json.Marshal(input.Metadata)
		if err != nil {
			return err
		}
		metadata = encoded
	}

	log := &auditEngineModels.AuditLog{
		ID:             uuid.New(),
		OrganizationID: input.OrganizationID,
		ActorID:        input.ActorID,
		ActorType:      actorType,
		ActorIPAddress: input.ActorIPAddress,
		Module:         input.Module,
		Action:         input.Action,
		ResourceType:   input.ResourceType,
		ResourceID:     input.ResourceID,
		Description:    input.Description,
		Metadata:       metadata,
	}

	return s.auditLogRepo.CreateLog(ctx, log)
}

/* ---------------------------------- <***> --------------------------------- */

func (s *auditLogService) ListOrgLogs(
	ctx context.Context,
	orgID uuid.UUID,
	query auditEngineDtos.ListLogsQuery,
) ([]auditEngineModels.AuditLog, int64, error) {
	// Clamp pagination so a bad caller can't pull the whole history at once
	if query.Limit <= 0 {
		query.Limit = DefaultPageSize
	}
	if query.Limit > MaxPageSize {
		query.Limit = MaxPageSize
	}
	if query.Offset < 0 {
		query.Offset = 0
	}

	return s.auditLogRepo.GetLogsByOrgID(ctx, orgID, query)
}

/* ---------------------------------- <***> --------------------------------- */

func (s *auditLogService) GetResourceHistory(
	ctx context.Context,
	orgID uuid.UUID,
	resourceType string,
	resourceID uuid.UUID,
) ([]auditEngineModels.AuditLog, error) {
	return s.auditLogRepo.GetLogsByResource(ctx, orgID, resourceType, resourceID)
}
