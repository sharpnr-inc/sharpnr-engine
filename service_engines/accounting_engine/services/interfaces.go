package accounting_service

import (
	"context"
	accounting_engine_contracts "sharpnr-engine/service_engines/accounting_engine/contracts"
	accountingdtos "sharpnr-engine/service_engines/accounting_engine/dto"

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

	CreateDraftAccountingSetup(ctx context.Context, input *accounting_engine_contracts.DraftAccountingSetupInput) error
}

type AccountingBooksServiceInterface interface {
	GetAccountingBook(ctx context.Context, orgID, emplID, legalEntityID uuid.UUID) (accountingdtos.AccountingBook_Response_dto, error)
}

type AccountingEngineAccountGroupServiceInterface interface {
	GetAccountGroupDetails(ctx context.Context, orgID, emplID, accountGroupID uuid.UUID) (accountingdtos.AccountingGroup_Response_dto, error)
}

type AccountingEngineAccountServiceInterface interface {
	CreateAccount(ctx context.Context, orgID, emplID, legalEntityID uuid.UUID, accountRequest accountingdtos.CreateAccountRequest) (accountID uuid.UUID, err error)
	GetAccountLeger(ctx context.Context, orgID, emplID, legalEntityID, accountID uuid.UUID) (accountingdtos.AccountLedgerDto, error)
}

type AccountingEngineJournalEntryServiceInterface interface {
	CreateJournalEntry(ctx context.Context, orgID, emplID, legalEntityID uuid.UUID, journalEntryRequest accountingdtos.JournalEntryExternalRequestDto) (journalEntryID uuid.UUID, err error)
	GetJournalEntries(ctx context.Context, orgID, emplID, legalEntityID uuid.UUID) ([]accountingdtos.JournalResponseDto, error)
}

type AccountingEngineJounalLineServiceInterface interface {
	CreateJournalLine(ctx context.Context, orgID, emplID, legalEntityID uuid.UUID, journalLineRequest accountingdtos.JournalLineExternalRequestDto) (journalLineID uuid.UUID, err error)
}

type AccountingEnginePaymentsServiceInteface interface {
	ListAllPaymentOrder(ctx context.Context, orgID, emplID, legalEntityID uuid.UUID) ([]accountingdtos.PaymentOrderResponseDTO, error)
}

/* -------------------------------------------------------------------------- */
/*                          Resolvers for Accounting                          */
/* -------------------------------------------------------------------------- */

type AccountingEngineResolverInterface interface {
	// This is resovolver method reolve the AccountID
	// Given to Fetch the
	ResolveAccountIDByTaxComponent(ctx context.Context)
}

type AccountingEngineReportServiceInterface interface {
	GenerateProfitAndLossReport(ctx context.Context, orgID, emplID, legalEntityID uuid.UUID, periodFrom, periodTo string) (accountingdtos.ProfitAndLossReportResponseDto, error)
}

type AccountingSuppliersServiceInterface interface {
	GetLegalEntitySuppliers(ctx context.Context, orgID, emplID, legalEntityID uuid.UUID) ([]accounting_engine_contracts.LegalEntitySuppliersResponseDto, error)
	AddNewSuppliers(ctx context.Context, orgID, emplID, legalEntityID uuid.UUID, request *accounting_engine_contracts.AddNewSuppliersForLegalEntityRequestDto) error
}
