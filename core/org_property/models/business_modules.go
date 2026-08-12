package org_property_models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

// PropertyBusinessOperationModules is the registry of every business module
// running inside a property, and — crucially — whose it is.
//
// Today the module list is implicit: each domain root (hospitality.restaurant,
// hospitality.hotel, hospitality.bar) carries its own PropertyID, and the
// orchestrator fans out per domain to assemble the list. That works only while
// a property has exactly one operator. Once a mall has forty of them, "list the
// modules I am allowed to see" has to be answerable without querying every
// domain table in the system and filtering afterwards.
//
// This table becomes that index. One row per module instance, written by the
// domain orchestrator at setup time.
type PropertyBusinessOperationModules struct {
	ID                     uuid.UUID `gorm:"type:uuid;primaryKey"`
	OrganizationPropertyID uuid.UUID `gorm:"type:uuid;not null;index"`

	// PropertyOccupancyID is the grant this module operates under, and the
	// authoritative answer to "whose restaurant is this".
	PropertyOccupancyID uuid.UUID `gorm:"type:uuid;not null;index"`

	// OperatingOrganizationID is denormalised from the occupancy so the hot
	// path — scoping a module list to the caller's org — is a single indexed
	// predicate rather than a join.
	OperatingOrganizationID uuid.UUID `gorm:"type:uuid;not null;index"`

	// BusinessModuleName is the domain discriminator (RESTAURANT, HOTEL, BAR)
	// and BusinessModulesID the primary key of the row in that domain's table.
	// The pair is a polymorphic reference, so it cannot carry a real FK — the
	// domain services own the integrity of their side.
	BusinessModuleName string    `gorm:"type:varchar(50);not null;index"`
	BusinessModulesID  uuid.UUID `gorm:"type:uuid;not null;index"`

	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
}

func (PropertyBusinessOperationModules) TableName() string {
	return "organization_property.property_business_operation_modules"
}

/* ---------------------------------- <***> --------------------------------- */
