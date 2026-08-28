package sentinelengine_models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

// Group is a named entity users belong to; groups can nest via ParentGroupID.
// Identified by name within an org (slug was removed in migration 000039).
type Group struct {
	ID             uuid.UUID  `gorm:"type:uuid;primaryKey"`
	OrganizationID uuid.UUID  `gorm:"type:uuid;not null;index"`
	Name           string     `gorm:"type:text;not null"`
	Description    string     `gorm:"type:text"`       // optional human-facing description
	ParentGroupID  *uuid.UUID `gorm:"type:uuid;index"` // nullable: top-level groups have no parent
	CreatedAt      time.Time
	UpdatedAt      time.Time
	DeletedAt      gorm.DeletedAt `gorm:"index"`
}

func (Group) TableName() string { return "sentinel_engine.groups" }

/* ---------------------------------- <***> --------------------------------- */

type GroupsSentinelPolicies struct {
	ID               uuid.UUID `gorm:"type:uuid;primaryKey"`
	OrganizationID   uuid.UUID `gorm:"type:uuid;not null;index"`
	GroupID          uuid.UUID `gorm:"type:uuid;not null;index"`
	SentinelPolicyID uuid.UUID `gorm:"type:uuid;not null;index"`
	CreatedAt        time.Time
	DeletedAt        gorm.DeletedAt `gorm:"index"`
}

func (GroupsSentinelPolicies) TableName() string {
	return "sentinel_engine.groups_sentinel_policies"
}
