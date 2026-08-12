package accounting_service

import (
	"context"
	legal_entity_contracts "sharpnr-engine/core/legal_entity/contracts"
	accounting_engine_contracts "sharpnr-engine/sharp_accounting_engine/accounting_engine/contracts"
	accountingdtos "sharpnr-engine/sharp_accounting_engine/accounting_engine/dto"
	stockwaveEngineModels "sharpnr-engine/stockwave_engine/models"

	"github.com/google/uuid"
)

type AccountingServiceInterface interface {
	ListAllPaymentOrder(ctx context.Context, orgID, emplID, legalEntityID uuid.UUID) ([]accountingdtos.PaymentOrderResponseDTO, error)
	GetAccountingBook(ctx context.Context, orgID, emplID, legalEntityID uuid.UUID) (accountingdtos.AccountingBook_Response_dto, error)
	CreateAccount(ctx context.Context, orgID, emplID, legalEntityID uuid.UUID, accountRequest accountingdtos.CreateAccountRequest) (accountID uuid.UUID, err error)
	GetAccountLeger(ctx context.Context, orgID, emplID, legalEntityID, accountID uuid.UUID) (accountingdtos.AccountLedgerDto, error)
	GetAllInventoriesAccounts(ctx context.Context, orgID, emplID, legalEntity uuid.UUID) ([]stockwaveEngineModels.Inventory, error)
	AddNewPurchaseInvoice(ctx context.Context, orgID, emplID, legalEntityID uuid.UUID, request *accountingdtos.PurchaseInvoiceRequestDto) error
	GetLegalEntitySuppliers(ctx context.Context, orgID, emplID, legalEntityID uuid.UUID) ([]accounting_engine_contracts.LegalEntitySuppliersResponseDto, error)
	AddNewSuppliers(ctx context.Context, orgID, emplID, legalEntityID uuid.UUID, request *accounting_engine_contracts.AddNewSuppliersForLegalEntityRequestDto) error
	CreateJournalEntry(ctx context.Context, orgID, emplID, legalEntityID uuid.UUID, journalEntryRequest accountingdtos.JournalEntryExternalRequestDto) (journalEntryID uuid.UUID, err error)
	CreateJournalLine(ctx context.Context, orgID, emplID, legalEntityID uuid.UUID, journalLineRequest accountingdtos.JournalLineExternalRequestDto) (journalLineID uuid.UUID, err error)
	GetJournalEntries(ctx context.Context, orgID, emplID, legalEntityID uuid.UUID) ([]accountingdtos.JournalResponseDto, error)
	CreateDraftAccountingSetup(ctx context.Context, input *accounting_engine_contracts.DraftAccountingSetupInput) error

	legalEntityservice
	// Reports
	GenerateProfitAndLossReport(ctx context.Context, orgID, emplID, legalEntityID uuid.UUID, periodFrom, periodTo string) (accountingdtos.ProfitAndLossReportResponseDto, error)

	// New Temporary APIs
	GetAccountsSummaryOfLegalEntity(ctx context.Context, orgID, emplID, legalEntityID uuid.UUID)
}

type legalEntityservice interface {
	GetLegalEntities(ctx context.Context, orgID, emplID uuid.UUID) ([]legal_entity_contracts.LegalEntityResponseDto, error)
}
