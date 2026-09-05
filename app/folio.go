package app

import (
	auditEngineService "sharpnr-engine/audit_log_engine/service"
	folioEngineRepo "sharpnr-engine/folio_engine/repo"
	folioEngineService "sharpnr-engine/folio_engine/service"
	dbtx "sharpnr-engine/pkg/dbTx"

	"gorm.io/gorm"
)

func bootstrapFolioEngine(
	db *gorm.DB,
	txManager dbtx.TxManager,
	auditLogService auditEngineService.AuditLogServiceInterface,
) folioEngineService.FolioEngineServiceInterface {
	folioRepo := folioEngineRepo.NewFolioEngineRepo(db)
	return folioEngineService.NewFolioEngineService(folioRepo, txManager, auditLogService)
}
