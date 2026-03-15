package tax_engine_service

import (
	"context"
	tax_engine_contracts "sharpnr-engine/service_engines/tax_engine/contracts"

	"github.com/google/uuid"
)

type TaxEngineServiceInterface interface {
	CalculateTax(str string) string
	GetTaxDetails(ctx context.Context, taxCategoryID uuid.UUID, legalEntityID uuid.UUID) ([]tax_engine_contracts.TaxComponentBreakdown, error)
}
