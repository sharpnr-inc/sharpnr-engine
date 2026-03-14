package legal_entity_contracts

import (
	"time"

	"github.com/google/uuid"
)

// Request Body for the Creating the Tax Registration For the Legal Entity
// LegalEntityTaxRegistration
type CreateLegalEntityTaxRegistrationRequestDto struct {
	LegalEntityID      uuid.UUID `json:"legal_entity_id"`
	TaxType            string    `json:"tax_type"`
	RegistrationNumber string    `json:"registration_number"`

	Country string `json:"country"`
	State   string `json:"state"`

	EffectiveFrom time.Time  `json:"effective_from"`
	EffectiveTo   *time.Time `json:"effective_to,omitempty"`
}

/* ---------------------------------- <***> --------------------------------- */

type LegalEntityTaxRegistrationResponseDto struct {
	ID                 uuid.UUID `json:"id"`
	TaxType            string    `json:"tax_type"`
	RegistrationNumber string    `json:"registration_number"`

	Country string `json:"country"`
	State   string `json:"state"`

	EffectiveFrom time.Time  `json:"effective_from"`
	EffectiveTo   *time.Time `json:"effective_to,omitempty"`
	IsActive      bool       `json:"is_active"`
	LastUpdatedAt time.Time  `json:"last_updated_at"`
}

/* ---------------------------------- <***> --------------------------------- */

type CreateLegalEntityTaxCategory struct {
	LegalEntityID uuid.UUID `json:"legal_entity_id" validate:"required"`
	CategoryName  string    `json:"category_name" validate:"required"`
	Code          string    `json:"code" validate:"required,uppercase"` // e.g., ROOM_RENT
	Description   string    `json:"description,omitempty"`

	AccountingBookID    uuid.UUID `json:"accounting_book_id" validate:"required"`
	ReceivableAccountID uuid.UUID `json:"receivable_account_id" validate:"required"`
	IncomeAccountID     uuid.UUID `json:"income_account_id" validate:"required"`

	// Nested Rule creation
	Rule CreateLegalEntityTaxCategoryTaxRule `json:"rule" validate:"required"`
}

type CreateLegalEntityTaxCategoryTaxRule struct {
	// Note: CategoryID will be injected by the Service after the Category is saved
	LegalEntityTaxRegistrationID uuid.UUID `json:"legal_entity_tax_registration_id" validate:"required"`

	TaxType       string     `json:"tax_type" validate:"required"` // GST, VAT
	Rate          float64    `json:"rate" validate:"required"`
	MinValue      *float64   `json:"min_value,omitempty"`
	MaxValue      *float64   `json:"max_value,omitempty"`
	EffectiveFrom time.Time  `json:"effective_from" validate:"required"`
	EffectiveTo   *time.Time `json:"effective_to,omitempty"`

	// The individual rows from your UI table
	Components []TaxComponentRequestDto `json:"components" validate:"required,dive,min=1"`
}

type TaxComponentRequestDto struct {
	ComponentCode      string    `json:"component_code" validate:"required"`
	Rate               float64   `json:"rate" validate:"required"`
	LiabilityAccountID uuid.UUID `json:"liability_account_id" validate:"required"`
}

/* ---------------------------------- <***> --------------------------------- */


