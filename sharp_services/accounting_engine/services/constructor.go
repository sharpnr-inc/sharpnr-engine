package accounting_service

import accounting_repo "sharpnr-engine/sharp_services/accounting_engine/repository"

type accountingServices struct {
	accRepo accounting_repo.AccountingRepositoryInterface
}

func NewAccountingService(accRepo accounting_repo.AccountingRepositoryInterface) *accountingServices {
	return &accountingServices{
		accRepo: accRepo,
	}

}
