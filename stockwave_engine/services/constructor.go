package stockwaveEngineServices

import (
	auditEngineService "sharpnr-engine/audit_log_engine/service"
	dbtx "sharpnr-engine/pkg/dbTx"
	"sharpnr-engine/sero"
	stockwaveEngineRepo "sharpnr-engine/stockwave_engine/repo"
)

type stockwaveEngine struct {
	repo stockwaveEngineRepo.StockwaveEngineRepoInterface

	// txManager wraps a write and the audit entry that records it in one
	// transaction, so a rolled-back change never leaves a log line claiming it
	// happened. Provisioning needs it for a second reason: the account and its
	// settings row are one unit.
	txManager dbtx.TxManager

	// seroEngine carries the cross-engine events stockwave both emits and
	// reacts to.
	seroEngine *sero.SEROEngine

	// auditLogEngine records the account lifecycle onto the org's history. Nil
	// is a working configuration, so a build without the audit engine wired
	// still serves the stockwave API.
	auditLogEngine auditEngineService.AuditLogServiceInterface
}

func NewStockwaveEngineService(
	repo stockwaveEngineRepo.StockwaveEngineRepoInterface,
	txManager dbtx.TxManager,
	seroEngine *sero.SEROEngine,
	auditLogEngine auditEngineService.AuditLogServiceInterface,
) *stockwaveEngine {
	return &stockwaveEngine{
		repo:           repo,
		txManager:      txManager,
		seroEngine:     seroEngine,
		auditLogEngine: auditLogEngine,
	}
}
