package accounting_service

import (
	dbtx "sharpnr-engine/core/dbTx"
	accounting_repo "sharpnr-engine/sharp_services/accounting_engine/repository"
)

type accountingServices struct {
	accRepo   accounting_repo.AccountingRepositoryInterface
	txManager dbtx.TxManager
}

func NewAccountingService(
	accRepo accounting_repo.AccountingRepositoryInterface,
	txManager dbtx.TxManager,
) *accountingServices {
	return &accountingServices{
		accRepo:   accRepo,
		txManager: txManager,
	}

}
