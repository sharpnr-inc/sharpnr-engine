package org_property_models

import (
	property_constants "sharpnr-engine/core/org_property/constants"
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

// PropertySharedService is a facility the landlord runs once for the whole
// property and every occupant draws on — the security desk, reception, parking,
// power backup. Defining them explicitly is what makes shared-facility cost
// allocable per occupancy instead of buried in a lump CAM figure.
type PropertySharedService struct {
	ID                     uuid.UUID `gorm:"type:uuid;primaryKey"`
	OrganizationPropertyID uuid.UUID `gorm:"type:uuid;not null;index"`

	ServiceType property_constants.SharedServiceType `gorm:"type:varchar(40);not null;index"`
	Name        string                               `gorm:"type:varchar(150);not null"`
	Description string                               `gorm:"type:text"`

	// IsMandatory services are billed to every occupancy whether or not a
	// subscription row exists — fire safety and core security aren't opt-out.
	IsMandatory bool `gorm:"default:false;not null"`
	IsActive    bool `gorm:"default:true;not null"`

	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
}

func (PropertySharedService) TableName() string {
	return "organization_property.property_shared_service"
}

/* ---------------------------------- <***> --------------------------------- */

// PropertySharedServiceSubscription opts one occupancy into one optional shared
// service. Mandatory services need no row here.
type PropertySharedServiceSubscription struct {
	ID                      uuid.UUID `gorm:"type:uuid;primaryKey"`
	PropertySharedServiceID uuid.UUID `gorm:"type:uuid;not null;index"`
	PropertyOccupancyID     uuid.UUID `gorm:"type:uuid;not null;index"`

	StartsOn time.Time `gorm:"not null"`
	EndsOn   *time.Time
	IsActive bool `gorm:"default:true;not null"`

	CreatedAt time.Time
	UpdatedAt time.Time
}

func (PropertySharedServiceSubscription) TableName() string {
	return "organization_property.property_shared_service_subscription"
}

/* ---------------------------------- <***> --------------------------------- */
