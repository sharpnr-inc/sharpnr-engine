package accounting_service

import (
	"context"
	accounting_engine_contracts "sharpnr-engine/sharp_services/accounting_engine/contracts"
)

type AccountingServiceInterface interface {

	// Setup Related stuffs
	CreateDraftAccountingSetup(ctx context.Context, input *accounting_engine_contracts.DraftAccountingSetupInput) error
	SetupAccountingForLegalEntity()
}
