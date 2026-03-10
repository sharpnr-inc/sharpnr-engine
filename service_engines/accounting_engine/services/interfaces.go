package accounting_service

import (
	"context"
	accounting_engine_contracts "sharpnr-engine/service_engines/accounting_engine/contracts"
	accountingdtos "sharpnr-engine/service_engines/accounting_engine/dto"

	"github.com/google/uuid"
)

type AccountingServiceInterface interface {
	CreateDraftAccountingSetup(ctx context.Context, input *accounting_engine_contracts.DraftAccountingSetupInput) error
	AccountingBooksServiceInterface
}

type AccountingBooksServiceInterface interface {
	GetAccountingBook(ctx context.Context, orgID, emplID, legalEntityID uuid.UUID) (accountingdtos.AccountingBook_Response_dto, error)
}
