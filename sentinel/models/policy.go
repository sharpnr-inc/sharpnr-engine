package sentinelengine_models

import (
	"encoding/json"
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

// PolicyScope classifies which bucket a policy belongs to, so the UI can group
// policies as org-wide, role, group, or user. For now it is a classification
// label only (the engine still loads every enabled org policy); it will also
// drive the per-employee policy load once attachment-driven evaluation lands.
type PolicyScope string

const (
	PolicyScopeOrganization PolicyScope = "organization" // applies to everyone in the org (e.g. forbid gates)
	PolicyScopeRole         PolicyScope = "role"         // attached to a role
	PolicyScopeGroup        PolicyScope = "group"        // attached to a group
	PolicyScopeUser         PolicyScope = "user"         // attached to a single employee
)

// Valid reports whether s is a known scope.
func (s PolicyScope) Valid() bool {
	switch s {
	case PolicyScopeOrganization, PolicyScopeRole, PolicyScopeGroup, PolicyScopeUser:
		return true
	}
	return false
}

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
	Scope          PolicyScope     `gorm:"type:text;not null;index"` // UI classification bucket
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"` // optional: soft delete
}

func (SentinelPolicy) TableName() string {
	return "sentinel_engine.sentinel_policies"
}
