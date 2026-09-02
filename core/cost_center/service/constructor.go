package orgCostCenterService

import (
	auditEngineService "sharpnr-engine/audit_log_engine/service"
	orgCostCenterRepo "sharpnr-engine/core/cost_center/repo"
	dbtx "sharpnr-engine/pkg/dbTx"
	"sharpnr-engine/sero"
	"sharpnr-engine/sharp_mailer_engine"
)

type orgCostCenterServiceConstructor struct {
	repo orgCostCenterRepo.OrgCostCenterRepoInterface

	// txManager wraps a write and the audit entry that records it in one
	// transaction, so a rolled-back change never leaves a log line claiming it
	// happened.
	txManager   dbtx.TxManager
	auditEngine auditEngineService.AuditLogServiceInterface

	// Parked: wired here but not consumed by any department flow yet. A
	// department change notifies nobody and publishes no event today; both
	// land on this struct the day one does.
	mailerEngine *sharp_mailer_engine.SharpMailerEngine
	seroEngine   *sero.SEROEngine
}

func NewOrgCostCenterService(
	repo orgCostCenterRepo.OrgCostCenterRepoInterface,
	txManager dbtx.TxManager,
	auditEngine auditEngineService.AuditLogServiceInterface,
	mailerEngine *sharp_mailer_engine.SharpMailerEngine,
	seroEngine *sero.SEROEngine,
) *orgCostCenterServiceConstructor {
	return &orgCostCenterServiceConstructor{
		repo:         repo,
		txManager:    txManager,
		auditEngine:  auditEngine,
		mailerEngine: mailerEngine,
		seroEngine:   seroEngine,
	}
}
