package tax_engine_service

import (
	"context"
	tax_engine_contracts "sharpnr-engine/service_engines/tax_engine/contracts"

	"github.com/google/uuid"
)

type TaxEngineServiceInterface interface {
	GetTaxCategoryTaxBreakdown(ctx context.Context, taxCategoryID uuid.UUID, itemBasePrice float64) ([]tax_engine_contracts.TaxComponentBreakdown, error)
}
