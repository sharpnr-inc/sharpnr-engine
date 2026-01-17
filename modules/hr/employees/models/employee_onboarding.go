package employee_models

import (
	hr_emp_constants "sharpnr-engine/modules/hr/employees/constants"
	"time"

	"github.com/google/uuid"
)

type EmployeeOnboarding struct {
	ID uuid.UUID `gorm:"type:uuid;primaryKey"`

	EmployeeID uuid.UUID `gorm:"type:uuid;uniqueIndex"`

	Status hr_emp_constants.EmployeeOnboardingStatus `gorm:"type:varchar(50);not null"`

	LastUpdatedBy *uuid.UUID
	Notes         *string

	CreatedAt time.Time
	UpdatedAt time.Time
}
