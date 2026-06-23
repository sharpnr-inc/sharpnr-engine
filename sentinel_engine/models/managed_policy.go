package sentinelengine_models

import (
	"time"

	"github.com/google/uuid"
	"github.com/lib/pq"
	"gorm.io/gorm"
)

// Effect is the Cedar policy effect a managed policy compiles to.
type Effect string

const (
	EffectPermit Effect = "permit"
	EffectForbid Effect = "forbid"
)

// ManagedPolicy is a reusable, named bundle of permissions — the AWS "managed
// policy" analogue. It is the source of truth the UI renders and edits; the
// generator compiles it (per role it's attached to) into a concrete
// SentinelPolicy the engine evaluates.
//
// OrganizationID is nil for a *system* policy shipped to every org (e.g. the
// "Full Access" owner bundle); set for an org-specific policy authored by a
// founding employee.
type ManagedPolicy struct {
	ID             uuid.UUID      `gorm:"type:uuid;primaryKey"`
	OrganizationID *uuid.UUID     `gorm:"type:uuid;index"`        // nil = system policy (all orgs)
	Name           string         `gorm:"type:text;not null"`     // human label, unique per org
	Description    string         `gorm:"type:text"`              //
	Actions        pq.StringArray `gorm:"type:text[];not null"`   // Cedar action ids (catalog Keys)
	Effect         Effect         `gorm:"type:text;not null;default:permit"`
	ResourceType   string         `gorm:"type:text"`              // optional Cedar resource type; "" = any
	// AllActions, when true, grants EVERY action (a wildcard `action` scope),
	// including actions added later. Actions is ignored when this is set. Used by
	// the "Full Access" owner policy.
	AllActions bool `gorm:"not null;default:false"`
	System     bool `gorm:"not null;default:false"` // platform-provided; FE cannot edit/delete
	CreatedAt      time.Time
	UpdatedAt      time.Time
	DeletedAt      gorm.DeletedAt `gorm:"index"`
}

func (ManagedPolicy) TableName() string { return "sentinel_engine.managed_policies" }

// RolePolicy is the attachment edge: a role has a managed policy attached.
// Hard-deleted on detach (no DeletedAt), mirroring UserRole/UserGroup, so the
// same (role, policy) pair can be re-attached later.
type RolePolicy struct {
	OrganizationID  uuid.UUID `gorm:"type:uuid;not null;index"`
	RoleID          uuid.UUID `gorm:"type:uuid;primaryKey"`
	ManagedPolicyID uuid.UUID `gorm:"type:uuid;primaryKey;index"`
	AttachedBy      uuid.UUID `gorm:"type:uuid"` // employee who attached (audit; optional)
	CreatedAt       time.Time
}

func (RolePolicy) TableName() string { return "sentinel_engine.role_policies" }
