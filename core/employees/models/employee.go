// models/employee.go
package coreEmployeeModels

import (
	coreEmployeeConstants "sharpnr-engine/core/employees/constants"
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

// -------------------------------
// 1️⃣ CORE EMPLOYEE
// -------------------------------
type Employee struct {
	ID             uuid.UUID `gorm:"type:uuid;primaryKey"`
	OrganizationID uuid.UUID `gorm:"type:uuid;not null;index"`

	FirstName  string `gorm:"type:varchar(50);not null"`
	MiddleName string `gorm:"type:varchar(50)"`
	LastName   string `gorm:"type:varchar(50)"`
	NickName   string `gorm:"type:varchar(50)"`

	Email       string `gorm:"type:varchar(150);not null;index"`
	PhoneNumber string `gorm:"type:varchar(20)"`

	Status       coreEmployeeConstants.EmployeeStatus `gorm:"type:varchar(20);default:'ACTIVE'"`
	EmployeeCode string                               `gorm:"type:varchar(50)" json:"employee_code"`

	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
}

func (Employee) TableName() string {
	return "employee.employee"
}

/* ---------------------------------- <***> --------------------------------- */
