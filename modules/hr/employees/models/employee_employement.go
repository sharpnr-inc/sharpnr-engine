package employee_models

import (
	"time"

	"github.com/google/uuid"
)

type EmployeeEmployment struct {
	ID uuid.UUID `gorm:"type:uuid;primaryKey"`

	EmployeeID uuid.UUID `gorm:"type:uuid;uniqueIndex"`

	EmploymentType string // FULL_TIME, PART_TIME, CONTRACT
	Designation    string

	ConfirmationDate *time.Time

	CreatedAt time.Time
}

type EmployeeCompensation struct {
	ID uuid.UUID `gorm:"type:uuid;primaryKey"`

	EmployeeID uuid.UUID `gorm:"type:uuid;index"`

	BaseSalary float64
	Currency   string

	EffectiveFrom time.Time
	EffectiveTo   *time.Time
}

type EmployeeSalaryComponent struct {
	ID uuid.UUID `gorm:"type:uuid;primaryKey"`

	CompensationID uuid.UUID

	Component string // BASIC, HRA, BONUS, PF
	Amount    float64
}
