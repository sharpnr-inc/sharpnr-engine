package employee_models

import (
	hr_emp_constants "sharpnr-engine/modules/hr/employees/constants"
	"time"

	"github.com/google/uuid"
)

type EmployeeOnboarding struct {
	ID uuid.UUID `gorm:"type:uuid;primaryKey"`

	// Links to the core employee identity
	EmployeeID uuid.UUID `gorm:"type:uuid;uniqueIndex;not null"`

	// Status: e.g., PENDING_DOCUMENTS, TRAINING, READY_FOR_SHIFT, COMPLETED
	Status hr_emp_constants.EmployeeOnboardingStatus `gorm:"type:varchar(50);not null;index"`

	// Audit trail: Who pushed the status forward?
	LastUpdatedBy *uuid.UUID `gorm:"type:uuid"`
	Notes         *string    `gorm:"type:text"`

	CreatedAt time.Time
	UpdatedAt time.Time
}

func (EmployeeOnboarding) TableName() string {
	return "employee.employee_onboarding"
}
