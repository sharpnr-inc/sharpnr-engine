package app

import (
	auditEngineService "sharpnr-engine/audit_log_engine/service"
	legal_entity_service "sharpnr-engine/core/legal_entity/service"
	suppliers_service "sharpnr-engine/core/org_supplier/service"
	payment_engine_service "sharpnr-engine/payment_engine/services"
	dbtx "sharpnr-engine/pkg/dbTx"
	"sharpnr-engine/sero"
	accounting_repo "sharpnr-engine/sharp_accounting_engine/accounting_engine/repository"
	accounting_service "sharpnr-engine/sharp_accounting_engine/accounting_engine/services"
	stockwaveEngineServices "sharpnr-engine/stockwave_engine/services"

	invoice_engine_services "sharpnr-engine/sharp_accounting_engine/invoicing_engine/services"

	"gorm.io/gorm"
)

func bootstrapAccounting(
	db *gorm.DB,
	txManager dbtx.TxManager,

	invoiceEngine invoice_engine_services.InvoiceEngineServiceInterface,
	paymentEngine payment_engine_service.PaymentServiceInterface,
	stockwaveEngine stockwaveEngineServices.StockwaveEngineServicesInterface,
	supplierService suppliers_service.SupplierServiceInterface,
	seroEngine *sero.SEROEngine,
	legalEntity legal_entity_service.OrgLegalEntityServiceInterface,
	auditLogService auditEngineService.AuditLogServiceInterface,
) accounting_service.AccountingServiceInterface {
	repo := accounting_repo.NewAccountingRepository(db)
	return accounting_service.NewAccountingService(
		seroEngine,
		repo,
		txManager,
		invoiceEngine,
		paymentEngine,
		stockwaveEngine,
		supplierService,
		legalEntity,
		auditLogService,
	)
}
