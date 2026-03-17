package accounting_service

import (
	dbtx "sharpnr-engine/core/dbTx"
	accounting_repo "sharpnr-engine/service_engines/accounting_engine/repository"
	tax_engine_service "sharpnr-engine/service_engines/tax_engine/service"
)

type accountingServices struct {
	accRepo   accounting_repo.AccountingRepositoryInterface
	txManager dbtx.TxManager
	taxEngine tax_engine_service.TaxEngineServiceInterface
}

func NewAccountingService(
	accRepo accounting_repo.AccountingRepositoryInterface,
	txManager dbtx.TxManager,
	taxEngine tax_engine_service.TaxEngineServiceInterface,
) *accountingServices {
	return &accountingServices{
		accRepo:   accRepo,
		txManager: txManager,
		taxEngine: taxEngine,
	}

}
