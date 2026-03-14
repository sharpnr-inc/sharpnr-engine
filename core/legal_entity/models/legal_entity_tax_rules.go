package legal_entity_models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/datatypes"
)

/* ----------------------- Tax Registration -------------------- */

type LegalEntityTaxRegistration struct {
	ID            uuid.UUID `gorm:"type:uuid;primaryKey"`
	LegalEntityID uuid.UUID `gorm:"type:uuid;index;not null"`

	TaxType            string // GST, VAT
	RegistrationNumber string // GSTIN / VAT Number

	State   string
	Country string

	EffectiveFrom time.Time
	EffectiveTo   *time.Time

	IsActive  bool `gorm:"default:true"`
	CreatedAt time.Time
}

func (LegalEntityTaxRegistration) TableName() string {
	return "legal_entity.legal_entity_tax_registration"
}

/* ----------------------- Tax Category -------------------- */

type LegalEntityTaxCategory struct {
	ID            uuid.UUID `gorm:"type:uuid;primaryKey"`
	LegalEntityID uuid.UUID `gorm:"type:uuid;index;not null"`

	// Fixed: Changed from uuid type to string
	Name        string `gorm:"size:100"`
	Code        string `gorm:"size:50;not null;index"` // ROOM_RENT, FOOD_BEV
	Description string `gorm:"size:255"`

	IsActive  bool `gorm:"default:true"`
	CreatedAt time.Time
	UpdatedAt time.Time
}

func (LegalEntityTaxCategory) TableName() string {
	return "legal_entity.legal_entity_tax_category"
}

/* ----------------------- Tax Rules & Components -------------------- */

type LegalEntityTaxRule struct {
	ID                           uuid.UUID `gorm:"type:uuid;primaryKey"`
	LegalEntityID                uuid.UUID `gorm:"type:uuid;index;not null"`
	LegalEntityTaxCategoryID     uuid.UUID `gorm:"type:uuid;index;not null"`
	LegalEntityTaxRegistrationID uuid.UUID `gorm:"type:uuid;index;not null"`

	TaxType string  // GST, VAT
	Rate    float64 // Total aggregate rate (e.g., 12.0)

	ConditionJSON datatypes.JSON // e.g. {"room_price_gt": 7500}

	MinValue *float64
	MaxValue *float64

	EffectiveFrom time.Time
	EffectiveTo   *time.Time

	IsActive bool `gorm:"default:true"`
}

func (LegalEntityTaxRule) TableName() string {
	return "legal_entity.legal_entity_tax_rule"
}

type LegalEntityTaxComponent struct {
	ID                   uuid.UUID `gorm:"type:uuid;primaryKey"`
	LegalEntityTaxRuleID uuid.UUID `gorm:"type:uuid;index;not null"`

	ComponentCode string  `gorm:"size:20;not null"` // CGST, SGST, VAT
	Rate          float64 `gorm:"not null"`         // Individual % (e.g. 6.0)

	// RELATIONAL MAPPING: Matches the "Liability Account" column in UI
	// Points to the specific Tax Payable Ledger (UUID)
	LiabilityAccountID uuid.UUID `gorm:"type:uuid;not null"`

	CreatedAt time.Time
}

func (LegalEntityTaxComponent) TableName() string {
	return "legal_entity.legal_entity_tax_component"
}

/* ----------------------- Accounting Mapping -------------------- */

// Renamed for clarity: Maps a Category to its Ledger Destinations
type TaxCategoryAccountMapping struct {
	ID                       uuid.UUID `gorm:"type:uuid;primaryKey"`
	LegalEntityID            uuid.UUID `gorm:"type:uuid;index;not null"`
	LegalEntityTaxCategoryID uuid.UUID `gorm:"type:uuid;index;not null"`

	// Even with only one book, we keep this to define the context.
	// You can default this to your "Primary Book" ID in the backend.
	AccountingBookID uuid.UUID `gorm:"type:uuid;index;not null"`

	// Where the money is owed (e.g., Guest Ledger / Accounts Receivable)
	ReceivableAccountID uuid.UUID `gorm:"type:uuid;not null"`

	// Where the money is earned (e.g., Room Revenue / Food Sales)
	IncomeAccountID uuid.UUID `gorm:"type:uuid;not null"`

	IsActive  bool `gorm:"default:true"`
	CreatedAt time.Time
	UpdatedAt time.Time
}

func (TaxCategoryAccountMapping) TableName() string {
	return "legal_entity.tax_category_account_mapping"
}
