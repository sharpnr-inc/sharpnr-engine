// models/employee_leaves.go
package coreEmployeeModels

import (
	"time"

	"github.com/google/uuid"
)

// -------------------------------
// 1️⃣ LEAVE TYPES (Policy Definition)
// -------------------------------
type EmployeeLeaveType struct {
	ID             uuid.UUID `gorm:"type:uuid;primaryKey"`
	OrganizationID uuid.UUID `gorm:"type:uuid;not null;index"`

	// e.g., "Sick Leave", "Paid Time Off", "Bereavement"
	Name string `gorm:"type:varchar(50);not null"`

	CreatedAt time.Time
}

func (EmployeeLeaveType) TableName() string {
	return "employee.employee_leave_type"
}

// -------------------------------
// 2️⃣ LEAVE BALANCES (Individual Quotas)
// -------------------------------
type EmployeeLeaveBalance struct {
	ID          uuid.UUID `gorm:"type:uuid;primaryKey"`
	EmployeeID  uuid.UUID `gorm:"type:uuid;not null;index"`
	LeaveTypeID uuid.UUID `gorm:"type:uuid;not null;index"`

	// Using float64 to support half-day increments
	Balance float64 `gorm:"type:decimal(5,2);default:0"`

	UpdatedAt time.Time
}

func (EmployeeLeaveBalance) TableName() string {
	return "employee.employee_leave_balance"
}

// -------------------------------
// 3️⃣ LEAVE REQUESTS (The Workflow)
// -------------------------------
type EmployeeLeaveRequest struct {
	ID          uuid.UUID `gorm:"type:uuid;primaryKey"`
	EmployeeID  uuid.UUID `gorm:"type:uuid;not null;index"`
	LeaveTypeID uuid.UUID `gorm:"type:uuid;not null;index"`

	FromDate time.Time `gorm:"type:date;not null"`
	ToDate   time.Time `gorm:"type:date;not null"`

	// Status: PENDING, APPROVED, REJECTED, CANCELLED
	Status string `gorm:"type:varchar(20);default:'PENDING';index"`

	// Audit Trail
	ApprovedBy *uuid.UUID `gorm:"type:uuid"`
	ApprovedAt *time.Time

	Reason string `gorm:"type:text"`

	CreatedAt time.Time
	UpdatedAt time.Time
}

func (EmployeeLeaveRequest) TableName() string {
	return "employee.employee_leave_request"
}
