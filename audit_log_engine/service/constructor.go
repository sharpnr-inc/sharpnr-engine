package auditEngineService

import (
	auditEngineRepo "sharpnr-engine/audit_log_engine/repo"
)

type auditLogService struct {
	auditLogRepo auditEngineRepo.AuditLogRepoInterface
}

func NewAuditLogService(
	auditLogRepo auditEngineRepo.AuditLogRepoInterface,
) *auditLogService {
	return &auditLogService{
		auditLogRepo: auditLogRepo,
	}
}
