package app

import (
	dbtx "sharpnr-engine/core/dbTx"
	accounting_repo "sharpnr-engine/service_engines/accounting_engine/repository"
	accounting_service "sharpnr-engine/service_engines/accounting_engine/services"
	invoice_engine_services "sharpnr-engine/service_engines/invoicing_engine/services"
	tax_engine_service "sharpnr-engine/service_engines/tax_engine/service"

	"gorm.io/gorm"
)

func bootstrapAccounting(
	db *gorm.DB,
	txManager dbtx.TxManager,
	taxEngine tax_engine_service.TaxEngineServiceInterface,
	invoiceEngine invoice_engine_services.InvoiceEngineServiceInterface,
) accounting_service.AccountingServiceInterface {
	repo := accounting_repo.NewAccountingRepository(db)
	return accounting_service.NewAccountingService(repo, txManager, taxEngine, invoiceEngine)
}

