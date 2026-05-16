package models

import (
	organization_constants "sharpnr-engine/core/organization/constants"
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

/***********************************************************
 * This Table is Master Organization                       *
 * this table should be considered as Root or Orchestrator *
 ***********************************************************/
type Organization struct {
	ID        uuid.UUID                                 `gorm:"type:uuid;primaryKey;not null"`
	Name      string                                    `gorm:"size:100;not null"`
	SharpnrID string                                    `gorm:"size:50;not null;uniqueIndex"`
	Status    organization_constants.OrgSetupStatusType `gorm:"type:varchar(100)"`
	RootEmail string                                    `gorm:"size:255;index"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
}

func (Organization) TableName() string {
	return "core.organization"
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

func (OrganizationContact) TableName() string {
	return "core.organization_contact"
}
