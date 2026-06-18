package sentinelengine_models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

// Role is a reusable, named bundle of access that belongs to ONE organization.
// On its own a role grants nothing — its power comes from the Permission rows
// attached to it (see permission.go). Employees receive a role via EmployeeRole
// (see employee_role.go).
//
// Cedar refers to a role as Role::"<ID>".
type Role struct {
	ID             uuid.UUID `gorm:"type:uuid;primaryKey"`
	OrganizationID uuid.UUID `gorm:"type:uuid;not null;index"` // multi-tenancy: a role lives in one org

	Name        string `gorm:"type:varchar(100);not null"` // "Super Admin", "Mumbai Manager"
	Description string `gorm:"type:varchar(255)"`

	// IsSystem marks roles the platform creates and manages (the per-org
	// "Super Admin"). The UI should show these as locked — not editable or
	// deletable by users.
	IsSystem bool `gorm:"not null;default:false"`

	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
}

func (Role) TableName() string {
	return "policy_engine.role"
}
