package orchestrator_service

import (
	"context"
	"encoding/json"

	auditEngineDtos "sharpnr-engine/audit_log_engine/dto"
	auditEngineModels "sharpnr-engine/audit_log_engine/models"
	auditEngineService "sharpnr-engine/audit_log_engine/service"
	coreEmployeeServices "sharpnr-engine/core/employees/services"

	"github.com/google/uuid"
)

// auditOrchestraterService implements AuditLogOrchestrationInterface: the
// organization's history, as the history screen needs to read it.
//
// The employee service is here for one job — turning the actor id on every
// entry into the name of a person. Without it the screen renders a column of
// UUIDs.
//
// ⚠ No guard.Authorize calls yet (doc.go rule 1): the audit trail has no
// Sentinel resource registered, so there is no action id to check against.
// Reading an org's history is currently open to any authenticated employee of
// that org — wire the guard here the day the resource lands.
type auditOrchestraterService struct {
	auditLogService   auditEngineService.AuditLogServiceInterface
	employeeHRService coreEmployeeServices.EmployeeHRServiceInterface
}

/* ---------------------------------- <***> --------------------------------- */

// ListOrgAuditLogs serves one page of the organization's history, newest first.
//
// Limit and Offset are clamped here rather than only in the engine because the
// response reports them back: the frontend pages off these numbers, so they
// have to describe the page that was actually served, not the one that was
// asked for.
func (s *auditOrchestraterService) ListOrgAuditLogs(
	ctx context.Context,
	orgID, emplID uuid.UUID,
	query auditEngineDtos.ListLogsQuery,
) (auditEngineDtos.AuditLogPageResponseDto, error) {

	limit := query.Limit
	if limit <= 0 {
		limit = auditEngineService.DefaultPageSize
	}
	if limit > auditEngineService.MaxPageSize {
		limit = auditEngineService.MaxPageSize
	}

	offset := query.Offset
	if offset < 0 {
		offset = 0
	}

	query.Limit, query.Offset = limit, offset

	logs, total, err := s.auditLogService.ListOrgLogs(ctx, orgID, query)
	if err != nil {
		return auditEngineDtos.AuditLogPageResponseDto{}, err
	}

	return auditEngineDtos.AuditLogPageResponseDto{
		Logs: s.toEntries(ctx, orgID, emplID, logs),
		Pagination: auditEngineDtos.PaginationMeta{
			Total:   total,
			Limit:   limit,
			Offset:  offset,
			HasMore: int64(offset+len(logs)) < total,
		},
	}, nil
}

/* ---------------------------------- <***> --------------------------------- */

// GetResourceAuditHistory is the full trail of one entity — everything that
// ever happened to this department, this customer, this property.
//
// It is not paginated: one entity's trail is bounded by how often that entity
// was touched, and the screen it feeds is a timeline read top to bottom.
func (s *auditOrchestraterService) GetResourceAuditHistory(
	ctx context.Context,
	orgID, emplID uuid.UUID,
	resourceType string,
	resourceID uuid.UUID,
) ([]auditEngineDtos.AuditLogEntry, error) {

	logs, err := s.auditLogService.GetResourceHistory(ctx, orgID, resourceType, resourceID)
	if err != nil {
		return nil, err
	}

	return s.toEntries(ctx, orgID, emplID, logs), nil
}

/* ---------------------------------- <***> --------------------------------- */

// toEntries maps stored events onto the wire shape, resolving actor names in
// one pass over the org's employee list rather than one lookup per row.
func (s *auditOrchestraterService) toEntries(
	ctx context.Context,
	orgID, emplID uuid.UUID,
	logs []auditEngineModels.AuditLog,
) []auditEngineDtos.AuditLogEntry {

	names := s.actorNames(ctx, orgID, emplID, logs)

	entries := make([]auditEngineDtos.AuditLogEntry, 0, len(logs))
	for _, log := range logs {
		actor := auditEngineDtos.AuditLogActor{
			ID:        log.ActorID,
			Type:      log.ActorType,
			IPAddress: log.ActorIPAddress,
		}
		if log.ActorID != nil {
			actor.Name = names[*log.ActorID]
		}

		entry := auditEngineDtos.AuditLogEntry{
			ID:           log.ID,
			Actor:        actor,
			Module:       log.Module,
			Action:       log.Action,
			ResourceType: log.ResourceType,
			ResourceID:   log.ResourceID,
			Description:  log.Description,
			CreatedAt:    log.CreatedAt,
		}

		// datatypes.JSON is already the module's own payload; pass it through
		// untouched, and drop it when it is empty or not valid JSON so the
		// response never carries a field the frontend cannot parse.
		if len(log.Metadata) > 0 && json.Valid(log.Metadata) {
			entry.Metadata = json.RawMessage(log.Metadata)
		}

		entries = append(entries, entry)
	}

	return entries
}

/* ---------------------------------- <***> --------------------------------- */

// actorNames resolves the employee ids on a page of entries to names.
//
// Fail-soft: a history that cannot label its actors is still a usable history,
// so a failure here leaves the names empty instead of failing the read. The
// lookup is skipped entirely when every entry on the page is a system event.
func (s *auditOrchestraterService) actorNames(
	ctx context.Context,
	orgID, emplID uuid.UUID,
	logs []auditEngineModels.AuditLog,
) map[uuid.UUID]string {

	names := map[uuid.UUID]string{}
	if s.employeeHRService == nil {
		return names
	}

	hasEmployeeActor := false
	for _, log := range logs {
		if log.ActorID != nil {
			hasEmployeeActor = true
			break
		}
	}
	if !hasEmployeeActor {
		return names
	}

	employees, err := s.employeeHRService.ListOrgAllEmployees(ctx, orgID, emplID)
	if err != nil {
		return names
	}

	for _, e := range employees {
		if e == nil {
			continue
		}
		names[e.ID] = e.EmployeeName
	}

	return names
}

/* ---------------------------------- <***> --------------------------------- */
