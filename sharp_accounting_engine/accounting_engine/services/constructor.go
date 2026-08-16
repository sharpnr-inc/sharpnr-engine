package accounting_service

import (
	"context"
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

	invoiceEngine   invoice_engine_services.InvoiceEngineServiceInterface
	paymentEngine   payment_engine_service.PaymentServiceInterface
	inventoryEngine stockwaveEngineServices.InventoryEngineServicesInterface
	supplierService suppliers_service.SupplierServiceInterface
	legalEntity     legal_entity_service.OrgLegalEntityServiceInterface
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
	inventoryEngine stockwaveEngineServices.InventoryEngineServicesInterface,
	supplierService suppliers_service.SupplierServiceInterface,
	legalEntity legal_entity_service.OrgLegalEntityServiceInterface,
) *accountingServices {
	return &accountingServices{
		seroEngine:      seroEngine,
		accRepo:         accRepo,
		txManager:       txManager,
		invoiceEngine:   invoiceEngine,
		paymentEngine:   paymentEngine,
		inventoryEngine: inventoryEngine,
		supplierService: supplierService,
		legalEntity:     legalEntity,
	}
}
