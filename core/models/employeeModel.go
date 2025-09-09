package models

import (
	"sharpnr-engine/core/constants"
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Employee struct {
	ID        uuid.UUID `gorm:"type:uuid;primaryKey"`
	FirstName string    `gorm:"size:50;not null"`
	// Must be In the Org Scope
	NickName    string
	LastName    string `gorm:"size:50"`
	Email       string `gorm:"not null;unique"`
	PhoneNumber string `gorm:"size:20"`

	//This is a EmployeeID
	ReportingMangerID string

	SkipLevelMangerID string

	OrganizationID uuid.UUID    `gorm:"type:uuid;not null;index"`
	Organization   Organization `gorm:"foreignKey:OrganizationID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`

	DepartmentID string     `gorm:"type:uuid;not null;index"`
	Department   Department `gorm:"foreignKey:DepartmentID;references:ID"`

	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
	Role      constants.EmployeeRoles
}

// TableName specifies the table name
func (Employee) TableName() string {
	return "employees"
}
