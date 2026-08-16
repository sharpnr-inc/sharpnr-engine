package coreEmployeeModels

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

// EmployeeDepartments assigns an employee to an ORG-LEVEL department
// (core.departments) — the track used by organizations that want departments
// without managing cost centers.
//
// Organizations that do manage cost centers use EmployeeCostCenterDepartment
// instead, which points at core.cost_center_departments. The two tracks never
// mix, which is what keeps DepartmentID here unambiguous about the table it
// refers to.
type EmployeeDepartments struct {
	ID           uuid.UUID `gorm:"type:uuid;primaryKey"`
	EmployeeID   uuid.UUID `gorm:"type:uuid;not null;index"`
	DepartmentID uuid.UUID `gorm:"type:uuid;not null;index"`

	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
}

func (EmployeeDepartments) TableName() string {
	return "employee.employee_departments"
}

/* ---------------------------------- <***> --------------------------------- */
