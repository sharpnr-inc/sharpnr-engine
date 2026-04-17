package tax_engine_contracts

import "github.com/google/uuid"

type TaxCalculationRequest struct {
	LegalEntityID uuid.UUID
	TaxCategoryID uuid.UUID
	BaseAmount    float64
}

type TaxCalculationResponse struct {
	BaseAmount float64                 `json:"base_amount"`
	TotalTax   float64                 `json:"total_tax"`
	GrandTotal float64                 `json:"grand_total"`
	Components []TaxComponentBreakdown `json:"components"`
}

type TaxComponentBreakdown struct {
	Name string  `json:"name"` // e.g., "CGST"
	Rate float64 `json:"rate"` // e.g., 9.0
}
