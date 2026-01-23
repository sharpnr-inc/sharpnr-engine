package employee_models

import (
	"time"

	"github.com/google/uuid"
)

// -------------------------------
// 1️⃣ EMPLOYMENT DETAILS (The Contract)
// -------------------------------
type EmployeeEmployment struct {
	ID         uuid.UUID `gorm:"type:uuid;primaryKey"`
	EmployeeID uuid.UUID `gorm:"type:uuid;uniqueIndex;not null"`

	// EmploymentType: FULL_TIME, PART_TIME, CONTRACT, INTERN
	EmploymentType string `gorm:"type:varchar(30);not null"`
	Designation    string `gorm:"type:varchar(100);not null"`

	ConfirmationDate *time.Time `gorm:"type:date"`
	JoinedAt         time.Time  `gorm:"type:date;not null"`

	CreatedAt time.Time
}

func (EmployeeEmployment) TableName() string {
	return "employee.employee_employment"
}

// -------------------------------
// 2️⃣ COMPENSATION (The Salary Header)
// -------------------------------
type EmployeeCompensation struct {
	ID         uuid.UUID `gorm:"type:uuid;primaryKey"`
	EmployeeID uuid.UUID `gorm:"type:uuid;index;not null"`

	// Using decimal for financial precision
	BaseSalary float64 `gorm:"type:decimal(15,2);not null"`
	Currency   string  `gorm:"type:varchar(3);default:'USD'"`

	// Effective dating logic
	EffectiveFrom time.Time  `gorm:"type:date;not null"`
	EffectiveTo   *time.Time `gorm:"type:date"` // Null if current/active

	CreatedAt time.Time
}

func (EmployeeCompensation) TableName() string {
	return "employee.employee_compensation"
}

// -------------------------------
// 3️⃣ SALARY COMPONENTS (The Breakdown)
// -------------------------------
type EmployeeSalaryComponent struct {
	ID             uuid.UUID `gorm:"type:uuid;primaryKey"`
	CompensationID uuid.UUID `gorm:"type:uuid;index;not null"`

	// Component: BASIC, HRA, BONUS, PF, SPECIAL_ALLOWANCE
	Component string  `gorm:"type:varchar(50);not null"`
	Amount    float64 `gorm:"type:decimal(15,2);not null"`

	CreatedAt time.Time
}

func (EmployeeSalaryComponent) TableName() string {
	return "employee.employee_salary_component"
}
