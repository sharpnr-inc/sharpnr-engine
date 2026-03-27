package accounting_service

import (
	"context"
	accounting_engine_contracts "sharpnr-engine/service_engines/accounting_engine/contracts"
	accountingdtos "sharpnr-engine/service_engines/accounting_engine/dto"

	"github.com/google/uuid"
)

type AccountingServiceInterface interface {
	AccountingBooksServiceInterface
	AccountingJournalServiceInterface
	AccountingEngineResolverInterface
	AccountingEnginePaymentsServiceInteface

	CreateDraftAccountingSetup(ctx context.Context, input *accounting_engine_contracts.DraftAccountingSetupInput) error
	GetLegalEntityLedgerEntry(ctx context.Context, orgID, emplID, legalEntityID uuid.UUID) (accountingdtos.LedgerPaginationResponse, error)
	GetJournalEntryDetails(ctx context.Context, entryID uuid.UUID) (accountingdtos.JournalEntryDetailsDTO, error)
}

type AccountingBooksServiceInterface interface {
	GetAccountingBook(ctx context.Context, orgID, emplID, legalEntityID uuid.UUID) (accountingdtos.AccountingBook_Response_dto, error)
}

type AccountingJournalServiceInterface interface {
	AddJournalEntry(ctx context.Context, orgID, emplID uuid.UUID, entryRequest accountingdtos.JournalEntryRequestDto) error
	AddJournalLines(ctx context.Context, orgID, emplID uuid.UUID, line accountingdtos.JournalLineRequestDto) error
}

type AccountingEngineResolverInterface interface {
	// This is resovolver method reolve the AccountID
	// Given to Fetch the
	ResolveAccountIDByTaxComponent(ctx context.Context)
}

type AccountingEnginePaymentsServiceInteface interface {
	ListAllPaymentOrder(ctx context.Context, orgID, emplID, legalEntityID uuid.UUID) ([]accountingdtos.PaymentOrderResponseDTO, error)
}
