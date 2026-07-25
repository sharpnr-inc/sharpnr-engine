package sentinelengine_models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

// Role is a named entity that policies reference (principal in Role::"<name>").
// Name is the Cedar entity id; it is unique per org and must be stable because
// policies embed it as text.
type Role struct {
	ID             uuid.UUID `gorm:"type:uuid;primaryKey"`
	OrganizationID uuid.UUID `gorm:"type:uuid;not null;uniqueIndex:idx_role_org_name"`
	Name           string    `gorm:"type:text;not null;uniqueIndex:idx_role_org_name"` // human label + Cedar entity id
	Description    string    `gorm:"type:text"`
	CreatedAt      time.Time
	UpdatedAt      time.Time
	DeletedAt      gorm.DeletedAt `gorm:"index"`
}

func (Role) TableName() string { return "sentinel_engine.roles" }

/* ---------------------------------- <***> --------------------------------- */

// A role Can have Multiple Policies

type RolePolicies struct {
	ID             uuid.UUID `gorm:"type:uuid;primaryKey"`
	OrganizationID uuid.UUID `gorm:"type:uuid;not null;index"`

	RoleID           uuid.UUID `gorm:"type:uuid;not null;index"`
	SentinelPolicyID uuid.UUID `gorm:"type:uuid;not null;index"`

	CreatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
}

func (RolePolicies) TableName() string {
	return "sentinel_engine.role_policies"
}
