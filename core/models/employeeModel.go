package models

import (
	"time"

	"gorm.io/gorm"
)

type Employee struct {
	ID          string `gorm:"type:uuid;primaryKey;default:gen_random_uuid()"`
	FirstName   string `gorm:"size:50;not null"`
	LastName    string `gorm:"size:50"`
	Email       string `gorm:"not null;unique"`
	PhoneNumber string `gorm:"size:20"`

	OrganizationID string       `gorm:"type:uuid;not null;index"`
	Organization   Organization `gorm:"foreignKey:OrganizationID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`

	DepartmentID string     `gorm:"type:uuid;not null;index"`
	Department   Department `gorm:"foreignKey:DepartmentID;references:ID"`

	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`

	TerminalAccess []string `gorm:"type:text[]"`
}

// TableName specifies the table name
func (Employee) TableName() string {
	return "employees"
}
