package app

import (
	"sharpnr-engine/modules/accounting/accounting_repo"
	"sharpnr-engine/modules/accounting/accounting_service"

	"gorm.io/gorm"
)

func bootstrapAccounting(db *gorm.DB) accounting_service.AccountingServiceInterface {
	repo := accounting_repo.NewAccountingRepository(db)
	return accounting_service.NewAccountingService(repo)
}
