package models

import (
	"time"
	"github.com/google/uuid"
)

type Employee struct {
	EmployeeID         uint      `gorm:"primaryKey;autoIncrement"`
	FirstName          string    `gorm:"not null"`
	LastName           *string
	Email              string    `gorm:"unique;not null"`
	Phone              *string
	OrganizationID     uuid.UUID `gorm:"type:uuid;not null"` // FK to Organization.SharpnrOrgUUID
	Organization       Organization `gorm:"foreignKey:OrganizationID;references:SharpnrOrgUUID"`
	Department         string
	TerminalAccess     bool
	IdentificationCode string
	CreatedAt          time.Time
	UpdatedAt          time.Time
}