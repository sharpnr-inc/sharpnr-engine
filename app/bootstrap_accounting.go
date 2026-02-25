package app

import (
	dbtx "sharpnr-engine/core/dbTx"
	accounting_repo "sharpnr-engine/service_engines/accounting_engine/repository"
	accounting_service "sharpnr-engine/service_engines/accounting_engine/services"

	"gorm.io/gorm"
)

func bootstrapAccounting(db *gorm.DB, txManager dbtx.TxManager) accounting_service.AccountingServiceInterface {
	repo := accounting_repo.NewAccountingRepository(db)
	return accounting_service.NewAccountingService(repo, txManager)
}
