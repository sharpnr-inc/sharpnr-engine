package org_property_models

import (
	property_constants "sharpnr-engine/core/org_property/constants"
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

// PropertyOccupancy is the grant that lets one organization operate inside a
// property owned by another. It is the single source of truth for the whole
// multi-tenant story:
//
//   - it is what Sentinel reads to allow a cross-org request. Today the engine
//     denies any request whose resource org differs from the request org
//     (sentinel_engine/types.go). A live occupancy is the only thing that may
//     lift that, and only for the property subtree it names.
//   - it is what every business module hangs off, so "whose restaurant is this"
//     has an answer.
//   - it is what rent, CAM and shared-service charges bill against.
//
// The owning org's own operations are modelled as an occupancy too (tenure
// OWNER_OPERATED, nil unit). Keeping one path for both means the authorization
// and billing code never needs an "is this the landlord?" branch.
type PropertyOccupancy struct {
	ID                     uuid.UUID `gorm:"type:uuid;primaryKey"`
	OrganizationPropertyID uuid.UUID `gorm:"type:uuid;not null;index"`

	// PropertyUnitID is nil when the occupancy covers the whole property rather
	// than a demarcated unit — the single-occupant case.
	PropertyUnitID *uuid.UUID `gorm:"type:uuid;index"`

	// OccupantOrganizationID is the org that actually operates here. It equals
	// the property's OrganizationID for OWNER_OPERATED occupancies and differs
	// for every real tenant.
	OccupantOrganizationID uuid.UUID `gorm:"type:uuid;not null;index"`

	Tenure property_constants.OccupancyTenure `gorm:"type:varchar(30);not null"`
	Status property_constants.OccupancyStatus `gorm:"type:varchar(30);not null;index"`

	StartsOn time.Time  `gorm:"not null"`
	EndsOn   *time.Time // nil for OWNED / OWNER_OPERATED — no expiry

	// Billing counterparties. Charges raised against this occupancy are invoiced
	// from the landlord's legal entity to the occupant's, which is a B2B
	// inter-org flow — not the same shape as the existing customer invoicing.
	LandlordLegalEntityID *uuid.UUID `gorm:"type:uuid;index"`
	OccupantLegalEntityID *uuid.UUID `gorm:"type:uuid;index"`

	TerminatedAt      *time.Time
	TerminatedBy      *uuid.UUID `gorm:"type:uuid"`
	TerminationReason string     `gorm:"type:text"`

	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
}

func (PropertyOccupancy) TableName() string {
	return "organization_property.property_occupancy"
}

/* ---------------------------------- <***> --------------------------------- */

// PropertyJoinRequest is the handshake that precedes an occupancy. An outside
// org finds a PUBLIC property in search and raises a request; an employee on the
// landlord side approves it, which is what creates the PropertyOccupancy. Orgs
// must never be able to attach themselves to a property directly — that would
// be self-service access to another tenant's building.
type PropertyJoinRequest struct {
	ID                     uuid.UUID  `gorm:"type:uuid;primaryKey"`
	OrganizationPropertyID uuid.UUID  `gorm:"type:uuid;not null;index"`
	PropertyUnitID         *uuid.UUID `gorm:"type:uuid;index"` // requested unit, if the requester picked one

	RequestingOrganizationID uuid.UUID `gorm:"type:uuid;not null;index"`
	RequestedByUserID        uuid.UUID `gorm:"type:uuid;not null"` // employee of the requesting org

	Status  property_constants.JoinRequestStatus `gorm:"type:varchar(30);not null;index"`
	Message string                               `gorm:"type:text"`

	DecidedByUserID *uuid.UUID `gorm:"type:uuid"` // employee on the landlord side
	DecidedAt       *time.Time
	DecisionNote    string `gorm:"type:text"`

	// ResultingOccupancyID is set when the request is approved, so the audit
	// trail from "who asked" to "what was granted" stays intact.
	ResultingOccupancyID *uuid.UUID `gorm:"type:uuid;index"`

	CreatedAt time.Time
	UpdatedAt time.Time
}

func (PropertyJoinRequest) TableName() string {
	return "organization_property.property_join_request"
}

/* ---------------------------------- <***> --------------------------------- */
