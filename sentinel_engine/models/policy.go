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
	Policy         string    `gorm:"type:text;not null"`       // Cedar policy, human-readable text form
	Enabled        bool      `gorm:"not null;default:true"`    // disable without deleting

	// Provenance of a GENERATED policy. Both set => this row was compiled by the
	// generator from a (role, managed policy) attachment, and the generator owns
	// it (delete-by-source on regeneration). Both nil => hand-authored; the
	// generator must never touch it.
	SourceRoleID          *uuid.UUID `gorm:"type:uuid;index"`
	SourceManagedPolicyID *uuid.UUID `gorm:"type:uuid"`

	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"` // optional: soft delete
}

func (SentinelPolicy) TableName() string {
	return "sentinel_engine.sentinel_policies"
}
