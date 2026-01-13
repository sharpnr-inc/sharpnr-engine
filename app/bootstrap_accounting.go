package app

import (
	"sharpnr-engine/modules/accounting/accounting_repo"
	"sharpnr-engine/modules/accounting/accounting_service"

	"gorm.io/gorm"
)

type accountingBootstrap struct {
	Repo    accounting_repo.AccountingRepositoryInterface
	Service accounting_service.AccountingServiceInterface
}

func bootstrapAccounting(db *gorm.DB) *accountingBootstrap {
	repo := accounting_repo.NewAccountingRepository(db)
	service := accounting_service.NewAccountingService(repo)

	return &accountingBootstrap{
		Repo:    repo,
		Service: service,
	}
}
