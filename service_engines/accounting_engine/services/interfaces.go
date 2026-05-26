package accounting_service

import (
	"context"
	legal_entity_contracts "sharpnr-engine/core/legal_entity/contracts"
	accounting_engine_contracts "sharpnr-engine/service_engines/accounting_engine/contracts"
	accountingdtos "sharpnr-engine/service_engines/accounting_engine/dto"
	inventory_engine_models "sharpnr-engine/service_engines/inventory_engine/models"

	"github.com/google/uuid"
)

type AccountingServiceInterface interface {
	AccountingBooksServiceInterface
	AccountingEngineJournalEntryServiceInterface
	AccountingEngineJounalLineServiceInterface
	AccountingEngineResolverInterface
	AccountingEnginePaymentsServiceInteface
	AccountingEngineAccountServiceInterface
	AccountingEngineReportServiceInterface
	AccountingSuppliersServiceInterface
	PurchasesServiceInterfaces
	AccountingInventoryServiceInterface
	AccountingLegalEntityServiceInterface

	CreateDraftAccountingSetup(ctx context.Context, input *accounting_engine_contracts.DraftAccountingSetupInput) error
}

type AccountingLegalEntityServiceInterface interface {
	GetLegalEntities(ctx context.Context, orgID, emplID uuid.UUID) ([]legal_entity_contracts.LegalEntityResponseDto, error)
}

type AccountingBooksServiceInterface interface {
	GetAccountingBook(ctx context.Context, orgID, emplID, legalEntityID uuid.UUID) (accountingdtos.AccountingBook_Response_dto, error)
}

/* --------------------------------- <****> --------------------------------- */

type AccountingEngineAccountGroupServiceInterface interface {
	GetAccountGroupDetails(ctx context.Context, orgID, emplID, accountGroupID uuid.UUID) (accountingdtos.AccountingGroup_Response_dto, error)
}

/* --------------------------------- <****> --------------------------------- */

type AccountingEngineAccountServiceInterface interface {
	CreateAccount(ctx context.Context, orgID, emplID, legalEntityID uuid.UUID, accountRequest accountingdtos.CreateAccountRequest) (accountID uuid.UUID, err error)
	GetAccountLeger(ctx context.Context, orgID, emplID, legalEntityID, accountID uuid.UUID) (accountingdtos.AccountLedgerDto, error)
}

/* --------------------------------- <****> --------------------------------- */

type AccountingEngineJournalEntryServiceInterface interface {
	CreateJournalEntry(ctx context.Context, orgID, emplID, legalEntityID uuid.UUID, journalEntryRequest accountingdtos.JournalEntryExternalRequestDto) (journalEntryID uuid.UUID, err error)
	GetJournalEntries(ctx context.Context, orgID, emplID, legalEntityID uuid.UUID) ([]accountingdtos.JournalResponseDto, error)
}

/* --------------------------------- <****> --------------------------------- */

type AccountingEngineJounalLineServiceInterface interface {
	CreateJournalLine(ctx context.Context, orgID, emplID, legalEntityID uuid.UUID, journalLineRequest accountingdtos.JournalLineExternalRequestDto) (journalLineID uuid.UUID, err error)
}

/* --------------------------------- <****> --------------------------------- */

type AccountingEnginePaymentsServiceInteface interface {
	ListAllPaymentOrder(ctx context.Context, orgID, emplID, legalEntityID uuid.UUID) ([]accountingdtos.PaymentOrderResponseDTO, error)
}

/* --------------------------------- <****> --------------------------------- */

/* -------------------------------------------------------------------------- */
/*                          Resolvers for Accounting                          */
/* -------------------------------------------------------------------------- */

type AccountingEngineResolverInterface interface {
	// This is resovolver method reolve the AccountID
	ResolveAccountIDByTaxComponent(ctx context.Context)
}

type AccountingEngineReportServiceInterface interface {
	GenerateProfitAndLossReport(ctx context.Context, orgID, emplID, legalEntityID uuid.UUID, periodFrom, periodTo string) (accountingdtos.ProfitAndLossReportResponseDto, error)
}

/* -------------------------------------------------------------------------- */
/*                          Supplier Service Methods                          */
/* -------------------------------------------------------------------------- */

type AccountingSuppliersServiceInterface interface {
	GetLegalEntitySuppliers(ctx context.Context, orgID, emplID, legalEntityID uuid.UUID) ([]accounting_engine_contracts.LegalEntitySuppliersResponseDto, error)
	AddNewSuppliers(ctx context.Context, orgID, emplID, legalEntityID uuid.UUID, request *accounting_engine_contracts.AddNewSuppliersForLegalEntityRequestDto) error
}

type PurchasesServiceInterfaces interface {
	AddNewPurchaseInvoice(ctx context.Context, orgID, emplID, legalEntityID uuid.UUID, request *accountingdtos.PurchaseInvoiceRequestDto) error
	// GetPurchaseInvoice()
	// AddBatchPurchaseInvoiceLine()
	// AddPurchaseInvoiceLine()
	// GetPurchaseLineByPurchaseInvoiceID()
}

type AccountingInventoryServiceInterface interface {
	GetAllInventoriesAccounts(ctx context.Context, orgID, emplID, legalEntity uuid.UUID) ([]inventory_engine_models.Inventory, error)
}
