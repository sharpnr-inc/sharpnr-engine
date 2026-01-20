package org_legal_entity_models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/datatypes"
)

type LegalEntityTaxRegistration struct {
	ID uuid.UUID `gorm:"type:uuid;primaryKey"`

	OrganizationLegalEntityID uuid.UUID `gorm:"type:uuid;index;not null"`

	TaxType string
	// GST, VAT, SERVICE_TAX (future)

	RegistrationNumber string
	// GSTIN

	State   string
	Country string

	EffectiveFrom time.Time
	EffectiveTo   *time.Time

	IsActive bool

	CreatedAt time.Time
}

type LegalEntityTaxCategory struct {
	ID uuid.UUID `gorm:"type:uuid;primaryKey"`

	OrganizationLegalEntityID uuid.UUID `gorm:"type:uuid;index;not null"`

	Code string
	// ROOM_RENT, FOOD, LIQUOR, SERVICE

	Description string

	IsActive bool
}

type LegalEntityTaxRule struct {
	ID uuid.UUID `gorm:"type:uuid;primaryKey"`

	OrganizationLegalEntityID uuid.UUID `gorm:"type:uuid;index;not null"`
	TaxCategoryID             uuid.UUID `gorm:"type:uuid;index;not null"`

	TaxType string // GST

	Rate float64

	ConditionJSON datatypes.JSON
	// e.g. { "room_price_gt": 7500 }

	EffectiveFrom time.Time
	EffectiveTo   *time.Time

	IsActive bool
}
