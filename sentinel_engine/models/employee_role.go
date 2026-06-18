package sentinelengine_models

import (
	"time"

	"github.com/google/uuid"
)

// EmployeeRole assigns a Role to an Employee. It's a many-to-many link: one
// employee can hold several roles and receives the UNION of their permissions.
//
// EmployeeID points at the existing employee.employee table. The engine treats
// each assigned role as a Cedar "parent" of the employee, so a policy written
// for Role::"X" applies to every employee assigned that role.
type EmployeeRole struct {
	ID uuid.UUID `gorm:"type:uuid;primaryKey"`

	EmployeeID uuid.UUID `gorm:"type:uuid;not null;index;uniqueIndex:uq_employee_role"`
	RoleID     uuid.UUID `gorm:"type:uuid;not null;index;uniqueIndex:uq_employee_role"`

	// OrganizationID is denormalized from the employee/role for fast
	// "all assignments in this org" lookups and as a tenancy guard.
	OrganizationID uuid.UUID `gorm:"type:uuid;not null;index"`

	CreatedAt time.Time
	UpdatedAt time.Time
}

func (EmployeeRole) TableName() string {
	return "policy_engine.employee_role"
}
