package legal_entity_models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/datatypes"
)

/* ---------------------------------- <***> --------------------------------- */

type LegalEntityTaxRegistration struct {
	ID uuid.UUID `gorm:"type:uuid;primaryKey"`

	LegalEntityID uuid.UUID `gorm:"type:uuid;index;not null"`

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

func (LegalEntityTaxRegistration) TableName() string {
	return "legal_entity.legal_entity_tax_registration"
}

/* ---------------------------------- <***> --------------------------------- */

type LegalEntityTaxCategory struct {
	ID            uuid.UUID `gorm:"type:uuid;primaryKey"`
	LegalEntityID uuid.UUID `gorm:"type:uuid;index;not null"`
	CategoryName  string    `gorm:"type:uuid"`
	Code          string    // ROOM_RENT, FOOD, LIQUOR, SERVICE
	Description   string
	IsActive      bool
	CreatedAt     time.Time
	UpdatedAt     time.Time
}

func (LegalEntityTaxCategory) TableName() string {
	return "legal_entity.legal_entity_tax_category"
}

/* ---------------------------------- <***> --------------------------------- */

type LegalEntityTaxRule struct {
	ID uuid.UUID `gorm:"type:uuid;primaryKey"`

	LegalEntityID            uuid.UUID `gorm:"type:uuid;index;not null"`
	LegalEntityTaxCategoryID uuid.UUID `gorm:"type:uuid;index;not null"`

	TaxType string // GST

	Rate float64

	ConditionJSON datatypes.JSON
	// e.g. { "room_price_gt": 7500 }

	MinValue *float64
	MaxValue *float64

	EffectiveFrom time.Time
	EffectiveTo   *time.Time

	IsActive bool
}

func (LegalEntityTaxRule) TableName() string {
	return "legal_entity.legal_entity_tax_rule"
}

type LegalEntityTaxComponent struct {
	ID uuid.UUID `gorm:"type:uuid;primaryKey"`

	LegalEntityTaxRuleID uuid.UUID `gorm:"type:uuid;index;not null"`

	ComponentCode string `gorm:"size:20;not null"`
	// CGST, SGST, IGST, VAT

	Rate float64 `gorm:"not null"` // %

	CreatedAt time.Time
}

func (LegalEntityTaxComponent) TableName() string {
	return "legal_entity.legal_entity_tax_component"
}
