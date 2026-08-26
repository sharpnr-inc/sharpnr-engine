package sentinelengine_models

import (
	"encoding/json"
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

// SentinelPolicy stores a single Cedar policy (human-readable text form) for one
// organization.
// ID doubles as the cedar.PolicyID when building the PolicySet.
type SentinelPolicy struct {
	ID             uuid.UUID       `gorm:"type:uuid;primaryKey"`
	OrganizationID uuid.UUID       `gorm:"type:uuid;not null;index"` // tenant boundary
	Name           string          `gorm:"type:text"`                // human label for admin UX (optional)
	Description    string          `gorm:"type:text"`                // optional
	Policy         string          `gorm:"type:text;not null"`       // Cedar policy text — the form the engine compiles
	PolicyJson     json.RawMessage `gorm:"type:jsonb"`               // optional Cedar JSON form, stored for the FE builder (engine never reads it)
	Enabled        bool            `gorm:"not null;default:true"`    // disable without deleting
	CreatedAt      time.Time
	UpdatedAt      time.Time
	DeletedAt      gorm.DeletedAt `gorm:"index"` // optional: soft delete
}

func (SentinelPolicy) TableName() string {
	return "sentinel_engine.sentinel_policies"
}
