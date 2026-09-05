package siloEngineService

import (
	auditEngineService "sharpnr-engine/audit_log_engine/service"
	dbtx "sharpnr-engine/pkg/dbTx"
	siloEngineRepo "sharpnr-engine/silo_engine/repo"
)

type siloService struct {
	siloRepo siloEngineRepo.SiloEngineRepoInterface

	// txManager wraps a write and the audit entry that records it in one
	// transaction, so a rolled-back change never leaves a log line claiming it
	// happened. Provisioning needs it for a second reason: the account and its
	// settings row are one unit.
	txManager dbtx.TxManager

	// auditLogEngine records the account lifecycle onto the org's history. Nil
	// is a working configuration — recordAudit becomes a no-op — so a build
	// without the audit engine wired still serves the Silo API.
	auditLogEngine auditEngineService.AuditLogServiceInterface
}

func NewSiloEngineService(
	siloRepo siloEngineRepo.SiloEngineRepoInterface,
	txManager dbtx.TxManager,
	auditLogEngine auditEngineService.AuditLogServiceInterface,
) *siloService {
	return &siloService{
		siloRepo:       siloRepo,
		txManager:      txManager,
		auditLogEngine: auditLogEngine,
	}
}
