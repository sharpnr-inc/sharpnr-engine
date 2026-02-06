package org_property_models

import (
	"time"

	"github.com/google/uuid"
)

type OrganizationProperty struct {
	ID             uuid.UUID `gorm:"primaryKey"`
	OrganizationID uuid.UUID `gorm:"index;not null"`

	Code         string `gorm:"type:varchar(50);index;not null"` // internal short code
	Name         string `gorm:"type:varchar(150);not null"`
	IsConfigured bool   `gorm:"default:false"`

	CreatedAt time.Time
	UpdatedAt time.Time
}

func (OrganizationProperty) TableName() string {
	return "organization_property.organization_property"
}

// TODO we will do the Data Migration later
type OrganizationPropertyConfigurationStatus struct {
	ID                     uuid.UUID
	OrganizationPropertyID uuid.UUID
	IsConfigured           bool
}

func (OrganizationPropertyConfigurationStatus) TableName() string {
	return "organization_property.organization_property_configuration_statuses"
}

type OrganizationPropertyAddress struct {
	ID                     uuid.UUID `gorm:"primaryKey"`
	OrganizationID         uuid.UUID `gorm:"index;not null"`
	OrganizationPropertyID uuid.UUID `gorm:"index;not null"`

	AddressLine string `gorm:"type:text;not null"`
	City        string `gorm:"type:varchar(100);index"`
	CityCode    string `gorm:"type:varchar(20);index"`
	State       string `gorm:"type:varchar(100);index"`
	Country     string `gorm:"type:varchar(100);index"`
	Pincode     string `gorm:"type:varchar(20);index"`

	CreatedAt time.Time
	UpdatedAt time.Time
}

func (OrganizationPropertyAddress) TableName() string {
	return "organization_property.organization_property_address"
}

type OrganizationPropertyTimeZone struct {
	ID                     uuid.UUID `gorm:"primaryKey"`
	OrganizationPropertyID uuid.UUID `gorm:"uniqueIndex;not null"`

	Timezone string `gorm:"type:varchar(50);not null"` // e.g. Asia/Kolkata

	CreatedAt time.Time
	UpdatedAt time.Time
}

func (OrganizationPropertyTimeZone) TableName() string {
	return "organization_property.organization_property_timezone"
}
