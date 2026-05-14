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

	Rules []LegalEntityTaxCategoryTaxRuleResponseDto `json:"rules"`
}

/* ---------------------------------- <***> --------------------------------- */

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

type TaxCategoryReader struct {
	TaxCategoryID   uuid.UUID
	LegalEntityID   uuid.UUID
	TaxCategoryName string
	TaxCategoryCode string
	Description     string
	IsActive        bool
	Rule            []TaxCategoryRule
}

type TaxCategoryRule struct {
	TaxCategoryRuleID uuid.UUID
	TaxCategoryID     uuid.UUID
	TaxRegistrationID uuid.UUID
	TaxType           string
	Rate              float64
	MinValue          float64
	MaxValue          float64
	EffectiveFrom     time.Time
	EffectiveTo       *time.Time
	IsActive          bool
	TaxBreakdown      []TaxCategoryRule
}

type TaxCategoryTaxComponent struct {
	TaxComponentID    uuid.UUID
	TaxCategoryRuleID uuid.UUID
	Code              string
	TaxRate           float64
}

/* ---------------------------------- <***> --------------------------------- */

// Legal Enity General Response DTO,
//
// This can be General Slice
type LegalEntityResponseDto struct {
	ID          uuid.UUID
	Name        string
	LegalName   string
	EnitityType string
}

/* ---------------------------------- <***> --------------------------------- */
