package org_property_models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type OrganizationProperty struct {
	ID uuid.UUID `gorm:"primaryKey"`

	OrganizationID uuid.UUID `gorm:"index;not null"`

	Code string `gorm:"type:varchar(50);index;not null"`
	Name string `gorm:"type:varchar(150);not null"`

	// Address
	AddressLine string `gorm:"type:text"`
	City        string `gorm:"type:varchar(100);index"`
	CityCode    string `gorm:"type:varchar(20);index"`
	State       string `gorm:"type:varchar(100);index"`
	Country     string `gorm:"type:varchar(100);index"`
	Pincode     string `gorm:"type:varchar(20);index"`

	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
}

func (OrganizationProperty) TableName() string {
	return "organization_property.organization_property"
}

/* ---------------------------------- <***> --------------------------------- */

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

/* ---------------------------------- <***> --------------------------------- */

type OrganizationPropertyGeoLocation struct {
	ID                     uuid.UUID
	OrganizationPropertyID uuid.UUID

	Latitude  float64 `gorm:"type:decimal(9,6);not null"`
	Longitude float64 `gorm:"type:decimal(9,6);not null"`

	AccuracyM *int   `gorm:"type:int"`
	Source    string `gorm:"type:varchar(20)"`

	CapturedAt time.Time
	UpdatedAt  time.Time
}

func (OrganizationPropertyGeoLocation) TableName() string {
	return "organization_property.organization_property_geo_location"
}

/* ---------------------------------- <***> --------------------------------- */
