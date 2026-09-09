package accounting_service

import (
	"context"
	auditEngineService "sharpnr-engine/audit_log_engine/service"
	legal_entity_service "sharpnr-engine/core/legal_entity/service"
	suppliers_service "sharpnr-engine/core/org_supplier/service"
	payment_engine_service "sharpnr-engine/payment_engine/services"
	dbtx "sharpnr-engine/pkg/dbTx"
	"sharpnr-engine/sero"
	accountingdtos "sharpnr-engine/sharp_accounting_engine/accounting_engine/dto"
	accounting_repo "sharpnr-engine/sharp_accounting_engine/accounting_engine/repository"
	invoice_engine_services "sharpnr-engine/sharp_accounting_engine/invoicing_engine/services"
	stockwaveEngineServices "sharpnr-engine/stockwave_engine/services"

	"github.com/google/uuid"
)

type accountingServices struct {
	seroEngine *sero.SEROEngine
	accRepo    accounting_repo.AccountingRepositoryInterface
	txManager  dbtx.TxManager

	invoiceEngine          invoice_engine_services.InvoiceEngineServiceInterface
	paymentEngine          payment_engine_service.PaymentServiceInterface
	stockwaveEngineService stockwaveEngineServices.StockwaveEngineServicesInterface
	supplierService        suppliers_service.SupplierServiceInterface
	legalEntity            legal_entity_service.OrgLegalEntityServiceInterface

	// auditLogEngine records the org-level account lifecycle onto the org's
	// history. Nil is a working configuration — recordEngineAccountAudit becomes
	// a no-op — so a build without the audit engine wired still serves the API.
	auditLogEngine auditEngineService.AuditLogServiceInterface
}

// GetAccountGroupDetails implements [AccountingServiceInterface].
func (s *accountingServices) GetAccountGroupDetails(ctx context.Context, orgID uuid.UUID, emplID uuid.UUID, accountGroupID uuid.UUID) (accountingdtos.AccountingGroup_Response_dto, error) {
	panic("unimplemented")
}

func NewAccountingService(
	seroEngine *sero.SEROEngine,
	accRepo accounting_repo.AccountingRepositoryInterface,
	txManager dbtx.TxManager,
	invoiceEngine invoice_engine_services.InvoiceEngineServiceInterface,
	paymentEngine payment_engine_service.PaymentServiceInterface,
	stockwaveEngineService stockwaveEngineServices.StockwaveEngineServicesInterface,
	supplierService suppliers_service.SupplierServiceInterface,
	legalEntity legal_entity_service.OrgLegalEntityServiceInterface,
	auditLogEngine auditEngineService.AuditLogServiceInterface,
) *accountingServices {
	return &accountingServices{
		seroEngine:             seroEngine,
		accRepo:                accRepo,
		txManager:              txManager,
		invoiceEngine:          invoiceEngine,
		paymentEngine:          paymentEngine,
		stockwaveEngineService: stockwaveEngineService,
		supplierService:        supplierService,
		legalEntity:            legalEntity,
		auditLogEngine:         auditLogEngine,
	}
}
