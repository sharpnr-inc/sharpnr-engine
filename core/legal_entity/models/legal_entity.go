package legal_entity_models

import (
	legal_entity_constants "sharpnr-engine/core/legal_entity/constants"
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

// =============================================================================
// CORE ENTITY
// =============================================================================

// LegalEntity is the root aggregate for a registered business entity.
// Identifiers like PAN and CIN live in LegalEntityIdentifier — adding
// new ID types (DPIIT, LLPIN, etc.) never requires a migration here.
type LegalEntity struct {
	ID             uuid.UUID `gorm:"type:uuid;primaryKey"`
	OrganizationID uuid.UUID `gorm:"type:uuid;not null;index"`

	LegalName   string `gorm:"size:255;not null"`
	DisplayName string `gorm:"size:255"`

	// EntityType drives which identifiers are expected and which
	// compliance rules apply. e.g. PVT_LTD requires CIN, LLP requires LLPIN.
	EntityType legal_entity_constants.LegalEntityType `gorm:"size:20;not null"` // see EntityType constants

	// IncorporationDate is needed for MCA filings, ROC compliance,
	// and age-gating checks in certain regulated sectors.
	IncorporationDate *time.Time

	// FYStartMonth is 1-12 (default 4 = April for Indian FY).
	// Used by accounting and tax modules to align reporting periods.
	FYStartMonth int `gorm:"default:4;not null"`

	IsActive  bool `gorm:"default:true;not null"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
}

func (LegalEntity) TableName() string {
	return "legal_entity.legal_entity"
}

/* ---------------------------------- <***> --------------------------------- */

// =============================================================================
// IDENTIFIERS
// =============================================================================

// LegalEntityIdentifier stores all government-issued IDs for an entity.
// One row per identifier type — PAN, CIN, TAN, UDYAM, IEC, etc.
// IsVerified flags cross-checking with the issuing authority (MCA, GSTN, etc.).
type LegalEntityIdentifier struct {
	ID            uuid.UUID `gorm:"type:uuid;primaryKey"`
	LegalEntityID uuid.UUID `gorm:"type:uuid;not null;index"`

	Type legal_entity_constants.LegalEntityIdentifierTypes `gorm:"size:30;not null"` // see Identifier constants

	Value      string `gorm:"size:100;not null"`
	IsVerified bool   `gorm:"default:false;not null"`
	VerifiedAt *time.Time

	// ExpiresAt is null for permanent IDs (PAN, CIN).
	// Set for time-bound registrations (IEC renewal, etc.).
	ExpiresAt *time.Time

	CreatedAt time.Time
	UpdatedAt time.Time
}

func (LegalEntityIdentifier) TableName() string {
	return "legal_entity.legal_entity_identifier"
}

/* ---------------------------------- <***> --------------------------------- */

// =============================================================================
// ADDRESSES
// =============================================================================

// LegalEntityAddress stores physical addresses.
// GSTStateCode is the critical field — it drives CGST+SGST vs IGST
// determination during tax resolution. Must match the 2-digit state
// prefix in a GSTIN (e.g. "06" for Haryana, "27" for Maharashtra).
//
// PropertyID optionally anchors an address to a physical property in another
// domain service (hotel, warehouse, outlet). Acts as cost-center key.
type LegalEntityAddress struct {
	ID            uuid.UUID `gorm:"type:uuid;primaryKey"`
	LegalEntityID uuid.UUID `gorm:"type:uuid;not null;index"`

	Type legal_entity_constants.LegalEntityAddressType `gorm:"size:20;not null"`

	AddressLine1 string `gorm:"size:255;not null"`
	AddressLine2 string `gorm:"size:255"`
	Landmark     string `gorm:"size:255"`
	City         string `gorm:"size:100;not null"`
	State        string `gorm:"size:100;not null"`

	// GSTStateCode: 2-digit code per GST council state list.
	// Used as join key between addresses and tax registrations.
	GSTStateCode string `gorm:"size:2"`

	Pincode string `gorm:"size:10;not null"`
	Country string `gorm:"size:2;not null;default:'IN'"` // ISO 3166-1 alpha-2

	// IsPrimary marks default address for correspondence and documents.
	// Enforce single-primary-per-entity in the service layer, not the DB.
	IsPrimary bool `gorm:"default:false;not null"`

	// PropertyID links address to a property record in another service.
	// Nullable — entity-level addresses (registered office) won't have one.
	PropertyID *uuid.UUID `gorm:"type:uuid"`

	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
}

func (LegalEntityAddress) TableName() string {
	return "legal_entity.legal_entity_address"
}

/* ---------------------------------- <***> --------------------------------- */

// =============================================================================
// COMPLIANCE / LICENSES
// =============================================================================

// LegalEntityCompliance stores all regulatory licenses and certifications.
//
// PropertyID scopes a license to one location (a PVT LTD running 3 hotels
// needs 3 FSSAI certificates). Module scopes it to a business vertical.
//
// The notification scheduler reads RenewalReminderDays to fire alerts
// before ValidTo. Set to 0 to disable reminders for a specific record.
type LegalEntityCompliance struct {
	ID            uuid.UUID `gorm:"type:uuid;primaryKey"`
	LegalEntityID uuid.UUID `gorm:"type:uuid;not null;index"`

	Type legal_entity_constants.LegalEntityLicenseType `gorm:"size:30;not null"`

	LicenseNumber    string `gorm:"size:100;not null"`
	IssuingAuthority string `gorm:"size:255"`

	ValidFrom time.Time
	ValidTo   *time.Time // null = perpetual / no expiry

	RenewalReminderDays int `gorm:"default:30;not null"`

	// PropertyID: null = entity-level license (e.g. company-wide trade license).
	// Set = location-specific license (e.g. FSSAI for one outlet).
	PropertyID *uuid.UUID `gorm:"type:uuid;index"`

	// Module: optional vertical scoping. Used by HasValidCompliance() when a
	// module checks its own license before serving (HOTEL | RESTAURANT | BAR).
	Module *string `gorm:"size:30"`

	DocumentURL *string `gorm:"size:512"` // S3 / GCS object path

	IsActive  bool `gorm:"default:true;not null"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
}

func (LegalEntityCompliance) TableName() string {
	return "legal_entity.legal_entity_compliance"
}
