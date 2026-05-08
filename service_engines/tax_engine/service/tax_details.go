package tax_engine_service

import (
	"context"
	tax_engine_contracts "sharpnr-engine/service_engines/tax_engine/contracts"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

func (s *taxEngineService) GetTaxCategoryTaxBreakdown(
	ctx context.Context,
	taxCategoryID uuid.UUID,
	itemBasePrice float64,
) ([]tax_engine_contracts.TaxComponentBreakdown, error) {

	// Initialize as an empty slice to ensure JSON returns [] instead of null
	res := []tax_engine_contracts.TaxComponentBreakdown{}

	err := s.txManager.WithTx(ctx, func(tx *gorm.DB) error {
		// 1. Fetch the full hierarchy from the Legal Entity Service
		taxCategoryDetails, err := s.legalEntityService.GetTaxCategoryDetails(ctx, taxCategoryID)
		if err != nil {
			return err
		}

		// 2. Logic: For now, we take the first rule found.
		// Later, itemBasePrice will be used here to filter the correct rule based on Min/Max.
		if len(taxCategoryDetails.Rule) > 0 {
			// Take the primary/first rule
			activeRule := taxCategoryDetails.Rule[0]

			// 3. Map the nested TaxBreakdown (components) to the result slice
			for _, component := range activeRule.TaxBreakdown {
				res = append(res, tax_engine_contracts.TaxComponentBreakdown{
					Name: component.TaxType, // This contains the ComponentCode (e.g., "CGST")
					Rate: component.Rate,
				})
			}
		}

		return nil
	})

	if err != nil {
		return nil, err
	}

	return res, nil
}
