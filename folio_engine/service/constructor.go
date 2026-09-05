package folioEngineService

import (
	auditEngineService "sharpnr-engine/audit_log_engine/service"
	folioEngineRepo "sharpnr-engine/folio_engine/repo"
	dbtx "sharpnr-engine/pkg/dbTx"
)

type folioService struct {
	folioRepo folioEngineRepo.FolioEngineRepoInterface

	// txManager wraps a write and the audit entry that records it in one
	// transaction, so a rolled-back change never leaves a log line claiming it
	// happened. Provisioning needs it for a second reason: the account and its
	// settings row are one unit.
	txManager dbtx.TxManager

	// auditLogEngine records the account lifecycle onto the org's history. Nil
	// is a working configuration — recordAudit becomes a no-op — so a build
	// without the audit engine wired still serves the Folio API.
	auditLogEngine auditEngineService.AuditLogServiceInterface
}

func NewFolioEngineService(
	folioRepo folioEngineRepo.FolioEngineRepoInterface,
	txManager dbtx.TxManager,
	auditLogEngine auditEngineService.AuditLogServiceInterface,
) *folioService {
	return &folioService{
		folioRepo:      folioRepo,
		txManager:      txManager,
		auditLogEngine: auditLogEngine,
	}
}
