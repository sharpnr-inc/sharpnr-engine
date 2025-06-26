package models

import (
	"time"
)

type Employee struct {
	EmployeeID     uint   `gorm:"primaryKey;autoIncrement"`
	FirstName      string `gorm:"not null"`
	LastName       *string
	Email          string `gorm:"unique;not null"`
	Phone          *string
	OrganizationID string       `gorm:"size:16;not null"` // FK to Organization.SharpnrOrgID
	Organization   Organization `gorm:"foreignKey:OrganizationID;references:SharpnrOrgID"`

	CreatedAt time.Time
	UpdatedAt time.Time
}
