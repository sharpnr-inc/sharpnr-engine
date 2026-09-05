package app

import (
	auditEngineService "sharpnr-engine/audit_log_engine/service"
	dbtx "sharpnr-engine/pkg/dbTx"
	siloEngineRepo "sharpnr-engine/silo_engine/repo"
	siloEngineService "sharpnr-engine/silo_engine/service"

	"gorm.io/gorm"
)

func bootstrapSiloEngine(
	db *gorm.DB,
	txManager dbtx.TxManager,
	auditLogService auditEngineService.AuditLogServiceInterface,
) siloEngineService.SiloEngineServiceInterface {
	siloRepo := siloEngineRepo.NewSiloEngineRepo(db)
	return siloEngineService.NewSiloEngineService(siloRepo, txManager, auditLogService)
}
