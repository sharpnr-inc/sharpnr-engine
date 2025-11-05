package hr

import (
	"time"

	"gorm.io/gorm"
)

type Employee struct {
	ID             string
	OrganisationID string
	OrgAssignID    string

	Name            string
	Email           string
	IsEmailVerified bool
	Address         string
	Salary          int

	Position    string
	JoiningDate string
	Role        string

	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`

	IsIDVerfied      string
	EmergencyContact string
}
