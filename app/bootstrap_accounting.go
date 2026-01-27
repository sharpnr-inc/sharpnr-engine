package app

import (
	dbtx "sharpnr-engine/core/dbTx"
	accounting_repo "sharpnr-engine/sharp_services/accounting_engine/repository"
	accounting_service "sharpnr-engine/sharp_services/accounting_engine/services"

	"gorm.io/gorm"
)

func bootstrapAccounting(db *gorm.DB, txManager dbtx.TxManager) accounting_service.AccountingServiceInterface {
	repo := accounting_repo.NewAccountingRepository(db)
	return accounting_service.NewAccountingService(repo, txManager)
}
