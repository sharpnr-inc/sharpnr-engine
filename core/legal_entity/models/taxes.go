package legal_entity_models

import (
	legal_entity_constants "sharpnr-engine/core/legal_entity/constants"
	"time"

	"github.com/google/uuid"
	"gorm.io/datatypes"
	"gorm.io/gorm"
)

// =============================================================================
// TAX — REGISTRATION
// =============================================================================

// LegalEntityTaxRegistration stores each GSTIN or VAT number.
// An entity can have multiple GSTINs — one per state of operation.
//
// GSTStateCode here must match LegalEntityAddress.GSTStateCode so the
// tax resolver can join them: given a transaction in state "27", find
// the GSTIN whose GSTStateCode is "27", then decide CGST+SGST vs IGST
// based on whether supplier and buyer states match.
type LegalEntityTaxRegistration struct {
	ID            uuid.UUID `gorm:"type:uuid;primaryKey"`
	LegalEntityID uuid.UUID `gorm:"type:uuid;not null;index"`

	TaxType legal_entity_constants.LegalEntityTaxType `gorm:"size:10;not null"` // GST | VAT

	RegistrationNumber string `gorm:"size:20;not null"` // GSTIN is always 15 chars

	// GSTRegistrationType affects rate applicability and return forms.
	// COMPOSITION entities pay at a flat rate and cannot charge GST to customers.
	GSTRegistrationType legal_entity_constants.LegalEntityGSTRegistrationType `gorm:"size:20;not null;default:'REGULAR'"` // see GSTType constants

	State        string `gorm:"size:100"`
	GSTStateCode string `gorm:"size:2"` // "06", "27" — join key for tax resolution
	Country      string `gorm:"size:2;not null;default:'IN'"`

	EffectiveFrom time.Time
	EffectiveTo   *time.Time // null = currently active

	IsActive  bool `gorm:"default:true;not null"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
}

func (LegalEntityTaxRegistration) TableName() string {
	return "legal_entity.legal_entity_tax_registration"
}

/* ---------------------------------- <***> --------------------------------- */

// =============================================================================
// TAX — CATEGORY
// =============================================================================

// LegalEntityTaxCategory is a named bucket for taxable transaction types.
// Code is the stable key used by all other modules (e.g. "ROOM_RENT",
// "FOOD_BEV", "BANQUET"). Never change Code after creation — other services
// depend on it as a hardcoded constant.
//
// SACCode (Service Accounting Code) is mandatory on B2B GST invoices.
// HSNCode applies to goods-based categories (F&B items, retail, etc.).
type LegalEntityTaxCategory struct {
	ID            uuid.UUID `gorm:"type:uuid;primaryKey"`
	LegalEntityID uuid.UUID `gorm:"type:uuid;not null;index"`

	Name        string `gorm:"size:100;not null"`
	Code        string `gorm:"size:50;not null"` // stable key — never rename
	Description string `gorm:"size:255"`

	// SACCode: Service Accounting Code under GST.
	// e.g. "996311" for hotel accommodation, "996331" for restaurant.
	SACCode *string `gorm:"size:10"`

	// HSNCode: Harmonized System of Nomenclature for goods.
	// Required when selling physical goods (minibar, retail, etc.).
	HSNCode *string `gorm:"size:10"`

	IsActive  bool `gorm:"default:true;not null"`
	CreatedAt time.Time
	UpdatedAt time.Time
}

func (LegalEntityTaxCategory) TableName() string {
	return "legal_entity.legal_entity_tax_category"
}

/* ---------------------------------- <***> --------------------------------- */

// =============================================================================
// TAX — RULE
// =============================================================================

// LegalEntityTaxRule defines when a tax rate applies for a category.
// Multiple rules per category support slab-based taxation:
//
//	Priority 1: room rent >= 7500 → 18% GST
//	Priority 2: room rent < 7500  → 12% GST  (catch-all at priority 99)
//
// The resolver evaluates rules in ascending Priority order and stops at
// the first match. ConditionJSON encodes the match condition.
//
// ConditionJSON schema (all keys optional, combined with AND):
//
//	{"min_value": 7500}              — transaction value >= 7500
//	{"max_value": 7499}              — transaction value <= 7499
//	{"min_value": 1000, "max_value": 5000}
//	{}  or null                      — catch-all, always matches
type LegalEntityTaxRule struct {
	ID                           uuid.UUID `gorm:"type:uuid;primaryKey"`
	LegalEntityID                uuid.UUID `gorm:"type:uuid;not null;index"`
	LegalEntityTaxCategoryID     uuid.UUID `gorm:"type:uuid;not null;index"`
	LegalEntityTaxRegistrationID uuid.UUID `gorm:"type:uuid;not null;index"`

	TaxType legal_entity_constants.LegalEntityTaxType `gorm:"size:10;not null"` // GST | VAT

	Rate float64 `gorm:"not null"` // aggregate rate e.g. 18.0

	ConditionJSON datatypes.JSON // see schema above
	MinValue      *float64
	MaxValue      *float64

	// Priority: lower number evaluated first.
	// Use 1–10 for specific slabs, 99 for catch-all.
	Priority int `gorm:"default:99;not null"`

	EffectiveFrom time.Time
	EffectiveTo   *time.Time // null = currently active

	IsActive  bool `gorm:"default:true;not null"`
	CreatedAt time.Time
	UpdatedAt time.Time
}

func (LegalEntityTaxRule) TableName() string {
	return "legal_entity.legal_entity_tax_rule"
}

/* ---------------------------------- <***> --------------------------------- */

// =============================================================================
// TAX — COMPONENT
// =============================================================================

// LegalEntityTaxComponent breaks an aggregate rule rate into individual
// ledger-postable components.
//
// Example for 18% GST intra-state:
//
//	CGST → 9%  → LiabilityAccountID = <CGST Payable ledger UUID>
//	SGST → 9%  → LiabilityAccountID = <SGST Payable ledger UUID>
//
// Example for 18% GST inter-state:
//
//	IGST → 18% → LiabilityAccountID = <IGST Payable ledger UUID>
//
// The tax resolver picks CGST+SGST vs IGST by comparing the supplier's
// GSTStateCode with the buyer's state — that logic lives in the service layer,
// not here. Both sets of components exist; the wrong set is simply skipped.
type LegalEntityTaxComponent struct {
	ID                   uuid.UUID `gorm:"type:uuid;primaryKey"`
	LegalEntityTaxRuleID uuid.UUID `gorm:"type:uuid;not null;index"`

	// ComponentCode: CGST | SGST | IGST | UGST | CESS | VAT
	ComponentCode legal_entity_constants.LegalEntityTaxComponentType `gorm:"size:10;not null"`

	Rate float64 `gorm:"not null"` // individual % e.g. 9.0

	// LiabilityAccountID points to one specific Tax Payable ledger
	// in your Chart of Accounts (CGST Payable, SGST Payable, etc.).
	LiabilityAccountID uuid.UUID `gorm:"type:uuid;not null"`

	CreatedAt time.Time
	UpdatedAt time.Time
}

func (LegalEntityTaxComponent) TableName() string {
	return "legal_entity.legal_entity_tax_component"
}

/* ---------------------------------- <***> --------------------------------- */

// =============================================================================
// TAX — ACCOUNT MAPPING
// =============================================================================

// TaxCategoryAccountMapping links a tax category to its revenue and AR ledgers.
// This is the accounting backbone of every transaction:
//
//	Debit  → ReceivableAccountID  (Guest Ledger / Accounts Receivable)
//	Credit → IncomeAccountID      (Room Revenue / Food Sales / Bar Revenue)
//	Credit → LiabilityAccountID   (per component, from TaxComponent table)
//
// AccountingBookID future-proofs multi-book setups (branch books, consolidated
// books). Default it to your primary book ID in the service layer.
type TaxCategoryAccountMapping struct {
	ID                       uuid.UUID `gorm:"type:uuid;primaryKey"`
	LegalEntityID            uuid.UUID `gorm:"type:uuid;not null;index"`
	LegalEntityTaxCategoryID uuid.UUID `gorm:"type:uuid;not null;index"`
	AccountingBookID         uuid.UUID `gorm:"type:uuid;not null;index"`

	// ReceivableAccountID: Accounts Receivable / Guest Ledger
	ReceivableAccountID uuid.UUID `gorm:"type:uuid;not null"`

	// IncomeAccountID: Revenue ledger (Room Revenue, Food Sales, etc.)
	IncomeAccountID uuid.UUID `gorm:"type:uuid;not null"`

	IsActive  bool `gorm:"default:true;not null"`
	CreatedAt time.Time
	UpdatedAt time.Time
}

func (TaxCategoryAccountMapping) TableName() string {
	return "legal_entity.tax_category_account_mapping"
}
