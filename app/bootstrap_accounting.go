package app

import (
	dbtx "sharpnr-engine/core/dbTx"
	suppliers_service "sharpnr-engine/core/org_supplier/service"
	payment_engine_service "sharpnr-engine/payment_engine/services"
	"sharpnr-engine/sero"
	accounting_repo "sharpnr-engine/service_engines/accounting_engine/repository"
	accounting_service "sharpnr-engine/service_engines/accounting_engine/services"
	inventory_engine_services "sharpnr-engine/service_engines/inventory_engine/services"
	invoice_engine_services "sharpnr-engine/service_engines/invoicing_engine/services"
	tax_engine_service "sharpnr-engine/service_engines/tax_engine/service"

	"gorm.io/gorm"
)

func bootstrapAccounting(
	db *gorm.DB,
	txManager dbtx.TxManager,
	taxEngine tax_engine_service.TaxEngineServiceInterface,
	invoiceEngine invoice_engine_services.InvoiceEngineServiceInterface,
	paymentEngine payment_engine_service.PaymentServiceInterface,
	inventoryEngine inventory_engine_services.InventoryEngineServicesInterface,
	supplierService suppliers_service.SupplierServiceInterface,
	seroEngine *sero.SEROEngine,
) accounting_service.AccountingServiceInterface {
	repo := accounting_repo.NewAccountingRepository(db)
	return accounting_service.NewAccountingService(
		seroEngine,
		repo,
		txManager,
		taxEngine,
		invoiceEngine,
		paymentEngine,
		inventoryEngine,
		supplierService,
	)
}
