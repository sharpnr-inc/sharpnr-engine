package sharphr_models

import (
	"sharpnr-engine/core/constants"
	"sharpnr-engine/core/models"
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
	Email       string `gorm:"not null"`
	PhoneNumber string `gorm:"size:20"`

	//This is a EmployeeID
	ReportingMangerID string

	SkipLevelMangerID string

	OrganizationID uuid.UUID           `gorm:"type:uuid;not null;index"`
	Organization   models.Organization `gorm:"foreignKey:OrganizationID"`

	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt          `gorm:"index"`
	Role      constants.EmployeeRoles `gorm:"type:varchar(50);not null"`

	IsEmailVerified  bool
	Address          string
	EmergencyContact string
	IsIDVerified     string
	Salary           string
}
