package auditEngineRepo

import (
	"context"
	auditEngineDtos "sharpnr-engine/audit_log_engine/dto"
	auditEngineModels "sharpnr-engine/audit_log_engine/models"

	"github.com/google/uuid"
)

type AuditLogRepoInterface interface {
	CreateLog(ctx context.Context, log *auditEngineModels.AuditLog) error

	// GetLogsByOrgID returns a page of the org's history (newest first)
	// plus the total count for pagination.
	GetLogsByOrgID(
		ctx context.Context,
		orgID uuid.UUID,
		query auditEngineDtos.ListLogsQuery,
	) ([]auditEngineModels.AuditLog, int64, error)

	// GetLogsByResource returns the full trail of a single entity
	// (e.g. everything that ever happened to one customer), newest first.
	GetLogsByResource(
		ctx context.Context,
		orgID uuid.UUID,
		resourceType string,
		resourceID uuid.UUID,
	) ([]auditEngineModels.AuditLog, error)
}
