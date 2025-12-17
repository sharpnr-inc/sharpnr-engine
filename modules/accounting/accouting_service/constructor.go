package accoutingservice

import "sharpnr-engine/modules/accounting/accounting_repo"

type accountingServices struct {
	accRepo accounting_repo.AccountingRepositoryInterface
}

func NewAccountingService(accRepo accounting_repo.AccountingRepositoryInterface) *accountingServices {
	return &accountingServices{
		accRepo: accRepo,
	}

}
