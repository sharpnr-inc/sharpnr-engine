package legal_entity_contracts

import (
	"time"

	"github.com/google/uuid"
)

/* ---------------------------------- <***> --------------------------------- */

type LegalEntityTaxCategoryResponseDto struct {
	ID           uuid.UUID `json:"id"`
	CategoryName string    `json:"category_name"`
	Code         string    `json:"code"` // e.g., ROOM_RENT
	Description  string    `json:"description,omitempty"`

	AccountingBookID    uuid.UUID `json:"accounting_book_id"`
	ReceivableAccountID uuid.UUID `json:"receivable_account_id"`
	IncomeAccountID     uuid.UUID `json:"income_account_id"`

	// Multiple rules
	Rules []LegalEntityTaxCategoryTaxRuleResponseDto `json:"rules"`
}

type LegalEntityTaxCategoryTaxRuleResponseDto struct {
	LegalEntityTaxRegistrationID uuid.UUID `json:"legal_entity_tax_registration_id"`

	TaxType string  `json:"tax_type"` // GST, VAT
	Rate    float64 `json:"rate"`

	MinValue *float64 `json:"min_value,omitempty"`
	MaxValue *float64 `json:"max_value,omitempty"`

	EffectiveFrom time.Time  `json:"effective_from"`
	EffectiveTo   *time.Time `json:"effective_to,omitempty"`

	Components []TaxComponentResponseDto `json:"components"`
}

type TaxComponentResponseDto struct {
	ComponentCode      string    `json:"component_code"`
	Rate               float64   `json:"rate"`
	LiabilityAccountID uuid.UUID `json:"liability_account_id"`
}

/* ---------------------------------- <***> --------------------------------- */
