package sentinelengine_models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

// Group is a named entity users belong to; groups can nest via ParentGroupID.
type Group struct {
	ID             uuid.UUID  `gorm:"type:uuid;primaryKey"`
	OrganizationID uuid.UUID  `gorm:"type:uuid;not null;uniqueIndex:idx_group_org_slug"`
	Slug           string     `gorm:"type:text;not null;uniqueIndex:idx_group_org_slug"`
	Name           string     `gorm:"type:text;not null"`
	Description    string     `gorm:"type:text"`       // optional human-facing description
	ParentGroupID  *uuid.UUID `gorm:"type:uuid;index"` // nullable: top-level groups have no parent
	CreatedAt      time.Time
	UpdatedAt      time.Time
	DeletedAt      gorm.DeletedAt `gorm:"index"`
}

func (Group) TableName() string { return "sentinel_engine.groups" }

// UserRole is the membership edge: a user (employee) is in a role.
//
// Hard-deleted on un-assignment. If you need an audit trail of who was
// granted/revoked what, write a separate audit log row — do NOT soft-delete
// this junction, because a soft-deleted (user_id, role_id) pair would block
// ever re-assigning that same role to that user.
type UserRole struct {
	OrganizationID uuid.UUID `gorm:"type:uuid;not null;index"`   // tenant scope; must equal roles.org_id
	UserID         uuid.UUID `gorm:"type:uuid;primaryKey"`       // = employee.id
	RoleID         uuid.UUID `gorm:"type:uuid;primaryKey;index"` // -> sentinel_engine.roles.id
	AssignedBy     uuid.UUID `gorm:"type:uuid"`                  // employee who made the assignment (audit; optional)
	CreatedAt      time.Time
}

func (UserRole) TableName() string { return "sentinel_engine.user_roles" }

// UserGroup is the membership edge: a user (employee) is in a group.
type UserGroup struct {
	OrganizationID uuid.UUID `gorm:"type:uuid;not null;index"`   // tenant scope; must equal groups.org_id
	UserID         uuid.UUID `gorm:"type:uuid;primaryKey"`       // = employee.id
	GroupID        uuid.UUID `gorm:"type:uuid;primaryKey;index"` // -> sentinel_engine.groups.id
	AssignedBy     uuid.UUID `gorm:"type:uuid"`                  // employee who added the membership (audit; optional)
	CreatedAt      time.Time
}

func (UserGroup) TableName() string { return "sentinel_engine.user_groups" }
