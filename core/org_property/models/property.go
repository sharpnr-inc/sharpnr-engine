package org_property_models

import (
	"time"

	"github.com/google/uuid"
)

type OrganizationProperty struct {
	ID             uuid.UUID
	OrganizationID uuid.UUID
	Code           string
	Name           string
	CreatedAt      time.Time
	UpdatedAt      time.Time
}

// this table connects the Tables of Property and Domain
type OrganizationPropertyDomain struct {
	ID                     uuid.UUID `gorm:"type:uuid;primaryKey"`
	OrganizationPropertyID uuid.UUID `gorm:"type:uuid;index"`
	Domain                 string    // Hospitality / Medical / Retail
	Status                 string    // ENABLED, DISABLED
}

type OrganizationPropertyAddress struct {
	ID                     uuid.UUID
	OrganizationID         uuid.UUID
	OrganizationPropertyID uuid.UUID

	AddressLine string
	City        string
	CityCode    string
	State       string
	Country     string
	Pincode     string
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

type OrganizationPropertyTimeZone struct {
}
