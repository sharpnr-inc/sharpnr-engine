package orchestrator_service

import auditEngineService "sharpnr-engine/audit_log_engine/service"

type auditOrchestraterService struct {
	auditLogService auditEngineService.AuditLogServiceInterface
}


