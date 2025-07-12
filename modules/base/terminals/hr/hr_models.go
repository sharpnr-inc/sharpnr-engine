package hr

import (
	"time"

	"github.com/sharpnr-inc/sharpnr-engine/core/models"
)

type Employee struct {
	EmployeeID         uint   `gorm:"primaryKey;autoIncrement"`
	FirstName          string `gorm:"not null"`
	LastName           *string
	Email              string `gorm:"unique;not null"`
	Phone              *string
	OrganizationID     string              `gorm:"size:16;not null"` // FK to Organization.SharpnrOrgID
	Organization       models.Organization `gorm:"foreignKey:OrganizationID;references:SharpnrOrgID"`
	Department         string
	TerminalAccess     bool
	IdentificationCode string
	CreatedAt          time.Time
	UpdatedAt          time.Time
}
