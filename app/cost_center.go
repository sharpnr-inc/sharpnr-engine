package app

import (
	auditEngineService "sharpnr-engine/audit_log_engine/service"
	orgCostCenterRepo "sharpnr-engine/core/cost_center/repo"
	orgCostCenterService "sharpnr-engine/core/cost_center/service"
	dbtx "sharpnr-engine/pkg/dbTx"
	"sharpnr-engine/sero"
	sharp_mailer_engine "sharpnr-engine/sharp_mailer_engine"

	"gorm.io/gorm"
)

func bootstrapCostCenterService(
	db *gorm.DB,
	txManager dbtx.TxManager,
	auditEngine auditEngineService.AuditLogServiceInterface,
	mailerEngine *sharp_mailer_engine.SharpMailerEngine,
	seroEngine *sero.SEROEngine) orgCostCenterService.OrgCostCenterServiceInterface {
	costCenterRepo := orgCostCenterRepo.NewOrgCostCenterRepo(db)
	return orgCostCenterService.NewOrgCostCenterService(
		costCenterRepo,
		txManager,
		auditEngine,
		mailerEngine,
		seroEngine,
	)
}
