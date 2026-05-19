package legal_entity_models

import (
	"time"

	"github.com/google/uuid"
)

// LegalEntityEInvoiceConfig stores IRP (Invoice Registration Portal) credentials
// and e-invoicing settings per entity. Mandatory under GST for entities
// crossing the turnover threshold. Your invoice service reads this before
// generating any B2B invoice to decide whether IRN generation is needed.
type LegalEntityEInvoiceConfig struct {
	ID            uuid.UUID `gorm:"type:uuid;primaryKey"`
	LegalEntityID uuid.UUID `gorm:"type:uuid;not null;uniqueIndex"`

	IsEnabled bool `gorm:"default:false;not null"`

	// IRPUsername / IRPPasswordHash: credentials for the NIC IRP portal.
	// Store password as bcrypt hash; your service fetches a session token
	// using these, then caches the token with a TTL.
	IRPUsername     string `gorm:"size:100"`
	IRPPasswordHash string `gorm:"size:255"`

	// GSTINForEInvoice: the specific GSTIN used as supplier GSTIN on IRN.
	// Usually the entity's primary GSTIN.
	GSTINForEInvoice string `gorm:"size:15"`

	// TurnoverThresholdExceeded: set true once your finance team confirms
	// the entity has crossed the mandate threshold for the current FY.
	TurnoverThresholdExceeded bool `gorm:"default:false;not null"`

	// LastIRNGeneratedAt: useful for rate-limit monitoring on the IRP API.
	LastIRNGeneratedAt *time.Time

	CreatedAt time.Time
	UpdatedAt time.Time
}

func (LegalEntityEInvoiceConfig) TableName() string {
	return "legal_entity.legal_entity_einvoice_config"
}
