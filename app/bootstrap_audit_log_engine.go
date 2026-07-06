package app

import (
	auditEngineRepo "sharpnr-engine/audit_log_engine/repo"
	auditEngineService "sharpnr-engine/audit_log_engine/service"

	"gorm.io/gorm"
)

// bootstrapAuditLogEngine wires the tenant event history recorder. The
// returned service is injected into every other engine so each module can
// call auditLogService.Record(...) and the org keeps a full history of
// what happened.
func bootstrapAuditLogEngine(
	db *gorm.DB,
) auditEngineService.AuditLogServiceInterface {
	repo := auditEngineRepo.NewAuditLogRepository(db)
	return auditEngineService.NewAuditLogService(repo)
}
