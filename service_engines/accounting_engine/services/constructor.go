package accounting_service

import (
	dbtx "sharpnr-engine/core/dbTx"
	payment_engine_service "sharpnr-engine/payment_engine/services"
	accounting_repo "sharpnr-engine/service_engines/accounting_engine/repository"
	invoice_engine_services "sharpnr-engine/service_engines/invoicing_engine/services"
	tax_engine_service "sharpnr-engine/service_engines/tax_engine/service"
)

type accountingServices struct {
	accRepo       accounting_repo.AccountingRepositoryInterface
	txManager     dbtx.TxManager
	taxEngine     tax_engine_service.TaxEngineServiceInterface
	invoiceEngine invoice_engine_services.InvoiceEngineServiceInterface
	paymentEngine payment_engine_service.PaymentServiceInterface
}

func NewAccountingService(
	accRepo accounting_repo.AccountingRepositoryInterface,
	txManager dbtx.TxManager,
	taxEngine tax_engine_service.TaxEngineServiceInterface,
	invoiceEngine invoice_engine_services.InvoiceEngineServiceInterface,
	paymentEngine payment_engine_service.PaymentServiceInterface,
) *accountingServices {
	return &accountingServices{
		accRepo:       accRepo,
		txManager:     txManager,
		taxEngine:     taxEngine,
		invoiceEngine: invoiceEngine,
		paymentEngine: paymentEngine,
	}
}
