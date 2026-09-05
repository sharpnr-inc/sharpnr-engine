package cortexEngineService

import (
	auditEngineService "sharpnr-engine/audit_log_engine/service"
	cortexEngineRepo "sharpnr-engine/cortex_engine/repo"
	dbtx "sharpnr-engine/pkg/dbTx"
)

type coretexService struct {

	// mailer
	cortexRepo cortexEngineRepo.CortexEngineRepoInterface

	// txManager wraps a write and the audit entry that records it in one
	// transaction, so a rolled-back change never leaves a log line claiming it
	// happened. Provisioning needs it for a second reason: the account and its
	// settings row are one unit.
	txManager dbtx.TxManager

	// auditLogEngine records the account lifecycle onto the org's history. Nil
	// is a working configuration — recordAudit becomes a no-op — so a build
	// without the audit engine wired still serves the Cortex API.
	auditLogEngine auditEngineService.AuditLogServiceInterface
}

func NewCortexEngineService(
	cortexRepo cortexEngineRepo.CortexEngineRepoInterface,
	txManager dbtx.TxManager,
	auditLogEngine auditEngineService.AuditLogServiceInterface,
) *coretexService {
	return &coretexService{
		cortexRepo:     cortexRepo,
		txManager:      txManager,
		auditLogEngine: auditLogEngine,
	}
}
