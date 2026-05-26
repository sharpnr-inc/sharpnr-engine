package accounting_service

import (
	legal_entity_service "sharpnr-engine/core/legal_entity/service"
	suppliers_service "sharpnr-engine/core/org_supplier/service"
	payment_engine_service "sharpnr-engine/payment_engine/services"
	dbtx "sharpnr-engine/pkg/dbTx"
	"sharpnr-engine/sero"
	accounting_repo "sharpnr-engine/service_engines/accounting_engine/repository"
	inventory_engine_services "sharpnr-engine/service_engines/inventory_engine/services"
	invoice_engine_services "sharpnr-engine/service_engines/invoicing_engine/services"
	tax_engine_service "sharpnr-engine/service_engines/tax_engine/service"
)

type accountingServices struct {
	seroEngine      *sero.SEROEngine
	accRepo         accounting_repo.AccountingRepositoryInterface
	txManager       dbtx.TxManager
	taxEngine       tax_engine_service.TaxEngineServiceInterface
	invoiceEngine   invoice_engine_services.InvoiceEngineServiceInterface
	paymentEngine   payment_engine_service.PaymentServiceInterface
	inventoryEngine inventory_engine_services.InventoryEngineServicesInterface
	supplierService suppliers_service.SupplierServiceInterface
	legalEntity     legal_entity_service.OrgLegalEntityServiceInterface
}

func NewAccountingService(
	seroEngine *sero.SEROEngine,
	accRepo accounting_repo.AccountingRepositoryInterface,
	txManager dbtx.TxManager,
	taxEngine tax_engine_service.TaxEngineServiceInterface,
	invoiceEngine invoice_engine_services.InvoiceEngineServiceInterface,
	paymentEngine payment_engine_service.PaymentServiceInterface,
	inventoryEngine inventory_engine_services.InventoryEngineServicesInterface,
	supplierService suppliers_service.SupplierServiceInterface,
	legalEntity legal_entity_service.OrgLegalEntityServiceInterface,
) *accountingServices {
	return &accountingServices{
		seroEngine:      seroEngine,
		accRepo:         accRepo,
		txManager:       txManager,
		taxEngine:       taxEngine,
		invoiceEngine:   invoiceEngine,
		paymentEngine:   paymentEngine,
		inventoryEngine: inventoryEngine,
		supplierService: supplierService,
		legalEntity:     legalEntity,
	}
}
