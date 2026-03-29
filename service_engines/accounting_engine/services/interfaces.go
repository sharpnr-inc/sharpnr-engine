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
	AccountingEngineLedgerServiceInterface

	CreateDraftAccountingSetup(ctx context.Context, input *accounting_engine_contracts.DraftAccountingSetupInput) error
}

type AccountingBooksServiceInterface interface {
	GetAccountingBook(ctx context.Context, orgID, emplID, legalEntityID uuid.UUID) (accountingdtos.AccountingBook_Response_dto, error)
	CreateAccount(ctx context.Context, orgID, emplID, legalEntityID uuid.UUID, accountRequest accountingdtos.CreateAccountRequest) (accountID uuid.UUID, err error)
}

type AccountingEngineAccountGroupServiceInterface interface {
	GetAccountGroupDetails(ctx context.Context, orgID, emplID, accountGroupID uuid.UUID) (accountingdtos.AccountingGroup_Response_dto, error)
}

type AccountingEnginePaymentsServiceInteface interface {
	ListAllPaymentOrder(ctx context.Context, orgID, emplID, legalEntityID uuid.UUID) ([]accountingdtos.PaymentOrderResponseDTO, error)
}

type AccountingEngineJournalEntryServiceInterface interface {
	GetJournalEntryDetails(ctx context.Context, entryID uuid.UUID) (accountingdtos.JournalEntryDetailsDTO, error)
	CreateJournalEntry(ctx context.Context, orgID, emplID, legalEntityID uuid.UUID, journalEntryRequest accountingdtos.JournalEntryExternalRequestDto) (journalEntryID uuid.UUID, err error)
	GetLegalEntityLedgerEntry(ctx context.Context, orgID, emplID, legalEntityID uuid.UUID) (accountingdtos.LedgerPaginationResponse, error)
}

type AccountingEngineJounalLineServiceInterface interface {
	CreateJournalLine(ctx context.Context, orgID, emplID, legalEntityID uuid.UUID, journalLineRequest accountingdtos.JournalLineExternalRequestDto) (journalLineID uuid.UUID, err error)
}

type AccountingEngineLedgerServiceInterface interface {
	GetAccountLeger(ctx context.Context, orgID, emplID, legalEntityID, accountID uuid.UUID) (accountingdtos.AccountLedgerDto, error)
}

/* -------------------------------------------------------------------------- */
/*                          Resolvers for Accounting                          */
/* -------------------------------------------------------------------------- */

type AccountingEngineResolverInterface interface {
	// This is resovolver method reolve the AccountID
	// Given to Fetch the
	ResolveAccountIDByTaxComponent(ctx context.Context)
}
