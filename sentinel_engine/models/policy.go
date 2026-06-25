package sentinelengine_models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

// SentinelPolicy stores a single Cedar policy (human-readable text form) for one
// organization.
// ID doubles as the cedar.PolicyID when building the PolicySet.
type SentinelPolicy struct {
	ID             uuid.UUID `gorm:"type:uuid;primaryKey"`
	OrganizationID uuid.UUID `gorm:"type:uuid;not null;index"` // tenant boundary
	Name           string    `gorm:"type:text"`                // human label for admin UX (optional)
	Description    string    `gorm:"type:text"`                // optional
	Policy         string    `gorm:"type:text;not null"`       // Cedar policy,
	Enabled        bool      `gorm:"not null;default:true"`    // disable without deleting

	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"` // optional: soft delete
}

func (SentinelPolicy) TableName() string {
	return "sentinel_engine.sentinel_policies"
}
