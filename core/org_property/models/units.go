package org_property_models

import (
	property_constants "sharpnr-engine/core/org_property/constants"
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

// PropertyUnit is a divisible physical space inside a property — shop G-101,
// the third floor, a kiosk by the atrium. It is the thing that gets leased or
// sold, and the thing an occupancy points at.
//
// A single-occupant property (one org running its own hotel) needs no units at
// all: its occupancy carries a nil PropertyUnitID and covers the whole property.
type PropertyUnit struct {
	ID                     uuid.UUID `gorm:"type:uuid;primaryKey"`
	OrganizationPropertyID uuid.UUID `gorm:"type:uuid;not null;index"`

	// Code is the landlord-facing label ("G-101"). Unique per property —
	// enforce with a composite unique index in the migration.
	Code string `gorm:"type:varchar(50);not null;index"`
	Name string `gorm:"type:varchar(150)"`

	UnitType property_constants.PropertyUnitType `gorm:"type:varchar(30);not null"`

	FloorLabel string `gorm:"type:varchar(30);index"` // "G", "1", "B2"
	Block      string `gorm:"type:varchar(30)"`

	// CarpetAreaSqFt is the usable area; ChargeableAreaSqFt is carpet plus the
	// unit's share of common area, and is what PER_SQ_FT charge schedules
	// multiply against. They are rarely equal — CAM is billed on the latter.
	CarpetAreaSqFt     *float64 `gorm:"type:decimal(12,2)"`
	ChargeableAreaSqFt *float64 `gorm:"type:decimal(12,2)"`

	Status property_constants.PropertyUnitStatus `gorm:"type:varchar(30);not null;index"`

	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
}

func (PropertyUnit) TableName() string {
	return "organization_property.property_unit"
}

/* ---------------------------------- <***> --------------------------------- */
