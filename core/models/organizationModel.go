package models

import (
	"sharpnr-engine/core/constants"
	"time"

	"github.com/google/uuid"
	"gorm.io/datatypes"
	"gorm.io/gorm"
)

type OrganizationSignUpStaging struct {
	ID       uuid.UUID
	Name     string
	Email    string
	Phone    string
	Otp      string
	Password string
}

/***********************************************************
 * This Table is Master Organization                       *
 * this table should be considered as Root or Orchestrator *
 ***********************************************************/
type Organization struct {
	ID        uuid.UUID `gorm:"not null;primaryKey"`
	Name      string    `gorm:"size:100;not null"`
	SharpnrID string    `gorm:"not null;unique"`

	OrgOperationType constants.OrgType `gorm:"type:varchar(20);not null"`
	Status           string            // CREATED, ACTIVE, SUSPENDED
	RootEmail        string

	// Timestamps
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
}

type OrganizationOffice struct {
	ID uuid.UUID `gorm:"type:uuid;primaryKey"`

	OrganizationID uuid.UUID `gorm:"type:uuid;index;not null"`

	// Office metadata
	Code string `gorm:"type:varchar(30)"`           // HQ, BLR-01, MUM-02
	Name string `gorm:"type:varchar(100);not null"` // Head Office, Bangalore Branch

	// Address
	AddressLine string `gorm:"type:varchar(255);not null"`
	City        string `gorm:"type:varchar(100);not null"`
	State       string `gorm:"type:varchar(100);not null"`
	Pincode     string `gorm:"type:varchar(20);not null"`
	Country     string `gorm:"type:varchar(100);not null"`

	// Flags
	IsHQ     bool `gorm:"default:false"`
	IsActive bool `gorm:"default:true"`

	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
}

type OrganizationContact struct {
	ID uuid.UUID `gorm:"type:uuid;primaryKey"`

	OrganizationID uuid.UUID `gorm:"type:uuid;index"`

	Label string `gorm:"type:varchar(50);not null"` // SALES, SUPPORT, BILLING

	ContactPerson *string `gorm:"type:varchar(100)"`
	Email         *string `gorm:"type:varchar(100)"`
	Phone         *string `gorm:"type:varchar(20)"`

	IsPrimary bool `gorm:"default:false"`

	CreatedAt time.Time
}

type OrganizationSubscription struct {
	ID uuid.UUID `gorm:"type:uuid;primaryKey"`

	OrganizationID uuid.UUID `gorm:"index"`

	PlanCode string
	Status   string // ACTIVE, PAUSED, CANCELLED

	RenewsAt time.Time

	CreatedAt time.Time
	UpdatedAt time.Time
}

type OrganizationVerification struct {
	ID uuid.UUID `gorm:"type:uuid;primaryKey"`

	OrganizationID uuid.UUID `gorm:"type:uuid;uniqueIndex"`

	EmailVerified    bool `gorm:"default:false"`
	MobileVerified   bool `gorm:"default:false"`
	PhysicalVerified bool `gorm:"default:false"`
	GSTVerified      bool `gorm:"default:false"`
	CINVerified      bool `gorm:"default:false"`

	VerifiedAt *time.Time
	VerifiedBy *uuid.UUID // admin / system

	CreatedAt time.Time
	UpdatedAt time.Time
}

type OrganizationSetting struct {
	ID uuid.UUID `gorm:"type:uuid;primaryKey"`

	OrganizationID uuid.UUID `gorm:"type:uuid;uniqueIndex"`

	Settings datatypes.JSON `gorm:"type:jsonb;default:'{}'"`

	UpdatedAt time.Time
}

type OrganizationOnBoarding struct {
	ID        uuid.UUID
	TempOrgID uuid.UUID
	Status    string
}
