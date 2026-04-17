package tax_engine_service

import (
	"context"
	tax_engine_contracts "sharpnr-engine/service_engines/tax_engine/contracts"
)

func (s *taxEngineService) CalculateTax(ctx context.Context, req tax_engine_contracts.TaxCalculationRequest) (*tax_engine_contracts.TaxCalculationResponse, error) {

	taxCategory, err := s.legalEntityService.GetTaxCategoryByID(ctx, req.TaxCategoryID)
	if err != nil {
		return nil, err
	}

	// 1. Basic Validation: Ensure the category is active and has rules
	if !taxCategory.IsActive || len(taxCategory.Rule) == 0 {
		return &tax_engine_contracts.TaxCalculationResponse{
			BaseAmount: req.BaseAmount,
			TotalTax:   0.0,
			GrandTotal: req.BaseAmount,
			Components: []tax_engine_contracts.TaxComponentBreakdown{},
		}, nil
	}

	// 2. Extract the primary rule (as per "one rule only for now")
	rule := taxCategory.Rule[0]

	// 3. Perform Calculations
	// Tax Amount = (Base * Rate) / 100
	taxAmount := req.BaseAmount * (rule.Rate / 100)

	// 4. Apply Min/Max constraints if defined
	if rule.MinValue > 0 && taxAmount < rule.MinValue {
		taxAmount = rule.MinValue
	}
	if rule.MaxValue > 0 && taxAmount > rule.MaxValue {
		taxAmount = rule.MaxValue
	}
	// 5. Build the Response
	response := &tax_engine_contracts.TaxCalculationResponse{
		BaseAmount: req.BaseAmount,
		TotalTax:   taxAmount,
		GrandTotal: req.BaseAmount + taxAmount,
		Components: []tax_engine_contracts.TaxComponentBreakdown{
			{
				Name: rule.TaxType,
				Rate: rule.Rate,
			},
		},
	}

	return response, nil
}
