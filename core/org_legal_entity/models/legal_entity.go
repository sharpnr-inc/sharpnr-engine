package org_legal_entity_models

import (
	"time"

	"github.com/google/uuid"
)

type OrganizationLegalEntity struct {
	ID uuid.UUID `gorm:"type:uuid;primaryKey"`

	OrganizationID uuid.UUID `gorm:"type:uuid;index;not null"`

	LegalName   string
	DisplayName string

	EntityType string
	// PVT_LTD, LLP, PROPRIETORSHIP, PARTNERSHIP

	GSTIN *string
	CIN   *string
	PAN   *string

	IsActive bool

	CreatedAt time.Time
	UpdatedAt time.Time
}

type LegalEntityAddress struct {
	ID uuid.UUID `gorm:"type:uuid;primaryKey"`

	OrganizationLegalEntityID uuid.UUID `gorm:"type:uuid;index"`

	Type string
	// REGISTERED, OPERATING

	AddressLine string
	City        string
	State       string
	Pincode     string
	Country     string

	IsPrimary bool
}

type PropertyLegalEntity struct {
	ID uuid.UUID `gorm:"type:uuid;primaryKey"`

	PropertyID                uuid.UUID `gorm:"type:uuid;index"`
	OrganizationLegalEntityID uuid.UUID `gorm:"type:uuid;index"`

	Role string
	// OWNER, OPERATOR, FRANCHISEE

	IsActive bool

	CreatedAt time.Time
}

type ModuleLegalEntity struct {
	ID uuid.UUID `gorm:"type:uuid;primaryKey"`

	PropertyID                uuid.UUID `gorm:"type:uuid;index"`
	Module                    string    // HOTEL, RESTAURANT, BAR
	OrganizationLegalEntityID uuid.UUID `gorm:"type:uuid;index"`

	CreatedAt time.Time
}

type LegalEntityCompliance struct {
	ID uuid.UUID `gorm:"type:uuid;primaryKey"`

	OrganizationLegalEntityID uuid.UUID `gorm:"type:uuid;index"`

	Type string // FSSAI, LIQUOR, FIRE, HEALTH, TRADE

	LicenseNumber    string
	IssuingAuthority string

	ValidFrom time.Time
	ValidTo   *time.Time

	Scope string
	// PROPERTY, MODULE, GLOBAL

	PropertyID *uuid.UUID
	Module     *string // HOTEL, RESTAURANT, BAR

	DocumentURL *string

	IsActive bool

	CreatedAt time.Time
}
