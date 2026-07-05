// models/employee_profile.go
package coreEmployeeModels

import (
	"time"

	"github.com/google/uuid"
)

// -------------------------------
// 1️⃣ PERSONAL PROFILE
// -------------------------------
type EmployeeProfile struct {
	ID         uuid.UUID `gorm:"type:uuid;primaryKey"`
	EmployeeID uuid.UUID `gorm:"type:uuid;not null;uniqueIndex"`

	Gender      string     `gorm:"type:varchar(10)"`
	DateOfBirth *time.Time `gorm:"type:date"`

	Address          string `gorm:"type:text"`
	EmergencyContact string `gorm:"type:varchar(100)"`
	BloodGroup       string `gorm:"type:varchar(5)"`

	CreatedAt time.Time
	UpdatedAt time.Time
}

func (EmployeeProfile) TableName() string {
	return "employee.employee_profile"
}

// -------------------------------
// 2️⃣ DOCUMENTS & COMPLIANCE
// -------------------------------
type EmployeeDocument struct {
	ID         uuid.UUID `gorm:"type:uuid;primaryKey"`
	EmployeeID uuid.UUID `gorm:"type:uuid;index;not null"`

	// Type: ID_PROOF, CONTRACT, VISA, HEALTH_CERTIFICATE
	Type string `gorm:"type:varchar(50)"`
	URL  string `gorm:"type:text"`

	Verified bool `gorm:"default:false"`

	CreatedAt time.Time
}

func (EmployeeDocument) TableName() string {
	return "employee.employee_document"
}

// -------------------------------
// 3️⃣ BANKING & PAYOUTS
// -------------------------------
type EmployeeBankAccount struct {
	ID         uuid.UUID `gorm:"type:uuid;primaryKey"`
	EmployeeID uuid.UUID `gorm:"type:uuid;uniqueIndex;not null"`

	BankName      string `gorm:"type:varchar(100)"`
	AccountNumber string `gorm:"type:varchar(50)"`
	IFSC          string `gorm:"type:varchar(20)"`
	UPI           string `gorm:"type:varchar(100)"`

	IsPrimary bool `gorm:"default:true"`

	CreatedAt time.Time
}

func (EmployeeBankAccount) TableName() string {
	return "employee.employee_bank_account"
}
