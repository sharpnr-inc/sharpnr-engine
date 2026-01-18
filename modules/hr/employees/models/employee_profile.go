package employee_models

import (
	"time"

	"github.com/google/uuid"
)

type EmployeeProfile struct {
	ID         uuid.UUID `gorm:"type:uuid;primaryKey"`
	EmployeeID uuid.UUID `gorm:"type:uuid;not null;uniqueIndex"`

	Gender      string
	DateOfBirth *time.Time

	Address          string
	EmergencyContact string
	BloodGroup       string

	CreatedAt time.Time
	UpdatedAt time.Time
}

type EmployeeDocument struct {
	ID uuid.UUID `gorm:"type:uuid;primaryKey"`

	EmployeeID uuid.UUID

	Type string // ID_PROOF, CONTRACT
	URL  string

	Verified bool
}

type EmployeeBankAccount struct {
	ID uuid.UUID `gorm:"type:uuid;primaryKey"`

	EmployeeID uuid.UUID `gorm:"type:uuid;uniqueIndex"`

	BankName      string
	AccountNumber string
	IFSC          string
	UPI           string

	IsPrimary bool
}
