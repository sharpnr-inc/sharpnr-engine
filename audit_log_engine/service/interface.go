package auditEngineService

import (
	"context"
	auditEngineDtos "sharpnr-engine/audit_log_engine/dto"
	auditEngineModels "sharpnr-engine/audit_log_engine/models"

	"github.com/google/uuid"
)

// AuditLogServiceInterface is what every other engine injects to record
// tenant events. Record is transaction-aware: called with a dbtx context
// the entry commits/rolls back together with the business action.
type AuditLogServiceInterface interface {
	Record(ctx context.Context, input auditEngineDtos.RecordEventInput) error

	ListOrgLogs(
		ctx context.Context,
		orgID uuid.UUID,
		query auditEngineDtos.ListLogsQuery,
	) ([]auditEngineModels.AuditLog, int64, error)

	GetResourceHistory(
		ctx context.Context,
		orgID uuid.UUID,
		resourceType string,
		resourceID uuid.UUID,
	) ([]auditEngineModels.AuditLog, error)
}
