package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

/***********************************************************
 * This Table is Master Organization                       *
 * this table should be considered as Root or Orchestrator *
 ***********************************************************/
type Organization struct {
	ID        uuid.UUID `gorm:"not null;primaryKey;index"`
	Name      string    `gorm:"size:100;not null"`
	SharpnrID string    `gorm:"not null;unique"`

	Status    string // CREATED, ACTIVE, SUSPENDED
	RootEmail string

	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
}

type OrganizationOnBoardingStage struct {
	ID             uuid.UUID
	OrganizationID uuid.UUID
	Status         string // SEED, // DOMAIN_ADDED // MODULE_ADDED
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
