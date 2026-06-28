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

// A User(employee) can have multiple Roles
//
// Note: UserID is EmployeeID or emplID
type UserRole struct {
	ID             uuid.UUID `gorm:"type:uuid;primaryKey"`
	OrganizationID uuid.UUID `gorm:"type:uuid;not null;index"` // tenant scope; must equal roles.org_id
	UserID         uuid.UUID `gorm:"type:uuid;not null;index"` // = employee.id
	RoleID         uuid.UUID `gorm:"type:uuid;not null;index"` // -> sentinel_engine.roles.id
	AssignedBy     uuid.UUID `gorm:"type:uuid"`                // employee who made the assignment (audit; optional)
	CreatedAt      time.Time
}

func (UserRole) TableName() string { return "sentinel_engine.user_roles" }

/* ---------------------------------- <***> --------------------------------- */

// this UserPolicy Table here in this table
//
// A User or Employee can be in a Group, have Roles multiple Roles, But Some times they need only one Policy
// So organization Can give user a Policy or Multiple Policy,
// lets just Say organizaion only wants to give some permission/policy they dont want to give Roles and Add in the Group
// So thats when this Table Comes in
// A User Can have multiple Policy
type UserPolicy struct {
	ID               uuid.UUID `gorm:"type:uuid;primaryKey"`
	OrganizationID   uuid.UUID `gorm:"type:uuid;not null;index"`
	EmployeeID       uuid.UUID `gorm:"type:uuid;not null;index"`
	SentinelPolicyID uuid.UUID `gorm:"type:uuid;not null;index"`

	CreatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
}

func (UserPolicy) TableName() string {
	return "sentinel_engine.user_policies"
}

/* ---------------------------------- <***> --------------------------------- */

// UserGroup is the membership edge: a user (employee) is in a group.
type UserGroup struct {
	ID             uuid.UUID `gorm:"type:uuid;primaryKey"`
	OrganizationID uuid.UUID `gorm:"type:uuid;not null;index"` // tenant scope; must equal groups.org_id
	UserID         uuid.UUID `gorm:"type:uuid;not null;index"` // = employee.id
	GroupID        uuid.UUID `gorm:"type:uuid;not null;index"` // -> sentinel_engine.groups.id
	AssignedBy     uuid.UUID `gorm:"type:uuid"`                // employee who added the membership (audit; optional)
	CreatedAt      time.Time
}

func (UserGroup) TableName() string { return "sentinel_engine.user_groups" }
