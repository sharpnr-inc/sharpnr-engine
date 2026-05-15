package legal_entity_models

import (
	legal_entity_constants "sharpnr-engine/core/legal_entity/constants"
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

// =============================================================================
// CONTACTS
// =============================================================================

// LegalEntityContact stores key personnel for an entity.
// Used for statutory filings, compliance communication, and audit trails.
// A Director contact is mandatory for PVT_LTD / LTD entity types.
// AuthorizedSignatory is the person who signs GST returns and invoices.
type LegalEntityContact struct {
	ID            uuid.UUID `gorm:"type:uuid;primaryKey"`
	LegalEntityID uuid.UUID `gorm:"type:uuid;not null;index"`

	Type legal_entity_constants.LegalEntityContactType `gorm:"size:30;not null"` // see Contact constants

	FirstName   string `gorm:"size:100;not null"`
	LastName    string `gorm:"size:100"`
	Designation string `gorm:"size:100"`

	Email  string `gorm:"size:255"`
	Mobile string `gorm:"size:15"`

	// DIN: Director Identification Number — required for directors of
	// companies registered under the Companies Act.
	DIN *string `gorm:"size:10"`

	// IsPrimary marks the main point of contact for a given type.
	// e.g. primary ACCOUNTS contact receives all payment reminders.
	IsPrimary bool `gorm:"default:false;not null"`

	IsActive  bool `gorm:"default:true;not null"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
}

func (LegalEntityContact) TableName() string {
	return "legal_entity.legal_entity_contact"
}
