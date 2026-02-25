package org_property_models

import (
	property_constants "sharpnr-engine/core/org_property/constants"
	"time"

	"github.com/google/uuid"
)

/*
This table connects Property with a Domain
Example domains: Hospitality / Medical / Retail
*/
type OrganizationPropertyDomain struct {
	ID uuid.UUID `gorm:"type:uuid;primaryKey"`

	OrganizationPropertyID uuid.UUID `gorm:"type:uuid;index;not null"`

	Domain property_constants.PropertyDomainTypes   `gorm:"type:varchar(50);index;not null"` // Hospitality / Medical / Retail
	Status property_constants.PropertyRunningStatus `gorm:"type:varchar(20);index;not null"` // ENABLED / DISABLED

	CreatedAt time.Time
	UpdatedAt time.Time
}

func (OrganizationPropertyDomain) TableName() string {
	return "organization_property.organization_property_domain"
}

// this is the register in this domain what modules they have
// a Property can arrays of **OrganizationPropertyDomainModules**

type OrganizationPropertyDomainModules struct {
	ID                           uuid.UUID
	OrganizationPropertyID       uuid.UUID
	OrganizationPropertyDomainID uuid.UUID
	Module                       string
	// Hotel, Bar, Restaurant, Night Club
	ModuleID uuid.UUID
	// HotelID, BraID, RestaurantID
}

func (OrganizationPropertyDomainModules) TableName() string {
	return "organization_property.organization_property_domain_modules"
}

type OrganizationPropertyDomainInfra struct {
	ID                           uuid.UUID
	OrganizationPropertyID       uuid.UUID
	OrganizationPropertyDomainID uuid.UUID
	InfraName                    string
	InfraType                    property_constants.PropertyInfraTypes
	IsShared                     *bool
	// KitchenID or Etc
	InfraModuleID uuid.UUID
}

func (OrganizationPropertyDomainInfra) TableName() string {
	return "organization_property.organization_property_domain_infra"
}
