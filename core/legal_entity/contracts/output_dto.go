package legal_entity_contracts

import (
	"time"

	"github.com/google/uuid"
)

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
