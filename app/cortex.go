package app

import (
	auditEngineService "sharpnr-engine/audit_log_engine/service"
	cortexEngineRepo "sharpnr-engine/cortex_engine/repo"
	cortexEngineService "sharpnr-engine/cortex_engine/service"
	dbtx "sharpnr-engine/pkg/dbTx"

	"gorm.io/gorm"
)

func bootstrapCortexEngine(
	db *gorm.DB,
	txManager dbtx.TxManager,
	auditLogService auditEngineService.AuditLogServiceInterface,
) cortexEngineService.CortexEngineServiceInterface {
	cortexRepo := cortexEngineRepo.NewCortexEngineRepo(db)
	return cortexEngineService.NewCortexEngineService(cortexRepo, txManager, auditLogService)
}
