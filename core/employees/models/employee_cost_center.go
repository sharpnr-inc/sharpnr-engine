package coreEmployeeModels

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

// EmployeeCostCenter is the employee's posting: the one cost center they work
// out of. An employee sits in exactly one at a time — enforce with a partial
// unique index on (employee_id) WHERE effective_to IS NULL AND deleted_at IS NULL.
//
// A transfer does not overwrite this row. It closes the current one by setting
// EffectiveTo and opens a new one at the destination, so a payroll run or cost
// report for a past period still resolves the office the employee was actually
// posted to at the time. Overwriting would silently rewrite history.
type EmployeeCostCenter struct {
	ID              uuid.UUID `gorm:"type:uuid;primaryKey"`
	EmployeeID      uuid.UUID `gorm:"type:uuid;not null;index"`
	OrganizationID  uuid.UUID `gorm:"type:uuid;not null;index"`
	OrgCostCenterID uuid.UUID `gorm:"type:uuid;not null;index"`

	EffectiveFrom time.Time  `gorm:"type:date;not null"`
	EffectiveTo   *time.Time `gorm:"type:date"` // nil = current posting

	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
}

func (t EmployeeCostCenter) TableName() string {
	return "employee.employee_cost_centers"
}

/* ---------------------------------- <***> --------------------------------- */

// EmployeeCostCenterDepartment assigns an employee to a department within the
// cost center they are posted to. Unlike the posting itself, this is not
// one-per-employee: someone can sit across two teams at the same site, so
// several live rows are allowed and exactly one carries IsPrimary.
//
// The cost center is implied by the department (CostCenterDepartment holds
// OrgCostCenterID) and deliberately not repeated here — storing it twice would
// let the two disagree after a transfer.
type EmployeeCostCenterDepartment struct {
	ID                     uuid.UUID `gorm:"type:uuid;primaryKey"`
	EmployeeID             uuid.UUID `gorm:"type:uuid;not null;index"`
	OrganizationID         uuid.UUID `gorm:"type:uuid;not null;index"`
	CostCenterDepartmentID uuid.UUID `gorm:"type:uuid;not null;index"`

	// IsPrimary marks the assignment that owns the employee's cost when they
	// sit across more than one department. Exactly one live row per employee
	// should carry it — enforce with a partial unique index on (employee_id)
	// WHERE is_primary AND effective_to IS NULL AND deleted_at IS NULL.
	IsPrimary bool `gorm:"default:true;not null"`

	EffectiveFrom time.Time  `gorm:"type:date;not null"`
	EffectiveTo   *time.Time `gorm:"type:date"` // nil = current assignment

	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
}

func (EmployeeCostCenterDepartment) TableName() string {
	return "employee.employee_cost_center_departments"
}

/* ---------------------------------- <***> --------------------------------- */
