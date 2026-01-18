package employee_models

import (
	"time"

	"github.com/google/uuid"
)

/*********************************************
 * All the Tables Related to Employee Leaves *
 *********************************************/

type EmployeeLeaveType struct {
	ID uuid.UUID `gorm:"type:uuid;primaryKey"`

	OrganizationID uuid.UUID
	Name           string // Casual, Sick, Paid
}

type EmployeeLeaveBalance struct {
	ID uuid.UUID `gorm:"type:uuid;primaryKey"`

	EmployeeID  uuid.UUID
	LeaveTypeID uuid.UUID

	Balance float64
}

type EmployeeLeaveRequest struct {
	ID uuid.UUID `gorm:"type:uuid;primaryKey"`

	EmployeeID  uuid.UUID
	LeaveTypeID uuid.UUID

	FromDate time.Time
	ToDate   time.Time

	Status string // PENDING, APPROVED, REJECTED

	ApprovedBy *uuid.UUID // Manager
	ApprovedAt *time.Time
}
