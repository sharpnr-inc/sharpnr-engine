package employee_models

import (
	"time"

	"github.com/google/uuid"
)

type EmployeePolicy struct {
	ID uuid.UUID `gorm:"type:uuid;primaryKey"`

	OrganizationID uuid.UUID `gorm:"type:uuid;index;not null"`

	Code string `gorm:"type:varchar(100);uniqueIndex;not null"`
	Name string

	Description string

	// SYSTEM, CUSTOM
	Type string `gorm:"type:varchar(30)"`

	CreatedAt time.Time
	UpdatedAt time.Time
}

type EmployeePrivilegePolicy struct {
	ID uuid.UUID `gorm:"type:uuid;primaryKey"`

	EmployeeID       uuid.UUID `gorm:"type:uuid;not null;index"`
	EmployeePolicyID uuid.UUID `gorm:"type:uuid;not null;index"`

	// Optional but VERY useful
	OrganizationID uuid.UUID `gorm:"type:uuid;not null;index"`
	Domain         string    `gorm:"type:varchar(50)"` // HOSPITALITY, HR, etc.

	AssignedAt time.Time
	AssignedBy uuid.UUID // HR employee
}

type EmployeePolicyRule struct {
	ID uuid.UUID `gorm:"type:uuid;primaryKey"`

	EmployeePolicyID uuid.UUID `gorm:"index;not null"`

	// Scope
	Module   string // BAR, RESTAURANT, HOTEL, HR
	Resource string // ORDER, MENU, BOOKING
	Action   string // VIEW, CREATE, UPDATE, DELETE

	// Effect
	Effect string // ALLOW or DENY

	CreatedAt time.Time
}
