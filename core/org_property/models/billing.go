package org_property_models

import (
	property_constants "sharpnr-engine/core/org_property/constants"
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

// PropertyChargeSchedule is a recurring money obligation attached to an
// occupancy — monthly rent, quarterly CAM, the parking fee. It is a definition,
// not a posting: the run job reads active schedules and emits PropertyChargeLine
// rows per period, which are then invoiced.
type PropertyChargeSchedule struct {
	ID                  uuid.UUID `gorm:"type:uuid;primaryKey"`
	PropertyOccupancyID uuid.UUID `gorm:"type:uuid;not null;index"`

	// PropertySharedServiceID is set only when ChargeType is SHARED_SERVICE, so
	// the resulting invoice line can name the facility being paid for.
	PropertySharedServiceID *uuid.UUID `gorm:"type:uuid;index"`

	ChargeType property_constants.ChargeType      `gorm:"type:varchar(30);not null;index"`
	Basis      property_constants.ChargeBasis     `gorm:"type:varchar(20);not null"`
	Frequency  property_constants.ChargeFrequency `gorm:"type:varchar(20);not null"`

	// Amount is read according to Basis: the flat figure for FIXED, or the rate
	// for PER_SQ_FT and METERED.
	Amount   float64 `gorm:"type:decimal(15,2);not null"`
	Currency string  `gorm:"type:varchar(3);default:'INR';not null"`

	// AreaSqFtSnapshot freezes the unit's chargeable area at the time the
	// schedule was agreed, so a later remeasurement of the unit doesn't silently
	// reprice a signed lease. Set for PER_SQ_FT only.
	AreaSqFtSnapshot *float64 `gorm:"type:decimal(12,2)"`

	StartsOn time.Time `gorm:"not null"`
	EndsOn   *time.Time

	// NextRunOn is what the generator queries on. Advanced by Frequency each
	// time a line is emitted.
	NextRunOn *time.Time `gorm:"index"`

	IsActive bool `gorm:"default:true;not null"`

	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
}

func (PropertyChargeSchedule) TableName() string {
	return "organization_property.property_charge_schedule"
}

/* ---------------------------------- <***> --------------------------------- */

// PropertyChargeLine is one generated period's charge — "CAM for unit G-101,
// March 2026". It exists as its own row before invoicing so charges can be
// reviewed, adjusted and disputed without touching a raised invoice, and so a
// re-run of the generator is idempotent against (schedule, period).
type PropertyChargeLine struct {
	ID                       uuid.UUID `gorm:"type:uuid;primaryKey"`
	PropertyChargeScheduleID uuid.UUID `gorm:"type:uuid;not null;index"`

	// Denormalised from the schedule so the common "all charges for this
	// occupancy" query doesn't need a join.
	PropertyOccupancyID uuid.UUID `gorm:"type:uuid;not null;index"`

	PeriodStart time.Time `gorm:"not null"`
	PeriodEnd   time.Time `gorm:"not null"`

	// Quantity carries the metered reading for METERED charges and the area for
	// PER_SQ_FT, so the line can show its own arithmetic. Nil for FIXED.
	Quantity *float64 `gorm:"type:decimal(15,4)"`
	Amount   float64  `gorm:"type:decimal(15,2);not null"`
	Currency string   `gorm:"type:varchar(3);default:'INR';not null"`

	Status property_constants.ChargeLineStatus `gorm:"type:varchar(20);not null;index"`

	// InvoiceID references the invoice raised in the accounting engine. Nil
	// until the line is billed.
	InvoiceID *uuid.UUID `gorm:"type:uuid;index"`

	CreatedAt time.Time
	UpdatedAt time.Time
}

func (PropertyChargeLine) TableName() string {
	return "organization_property.property_charge_line"
}

/* ---------------------------------- <***> --------------------------------- */
