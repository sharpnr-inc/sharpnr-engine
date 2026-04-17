package org_suppliers_models

import (
	"time"

	"github.com/google/uuid"
)

type Supplier struct {
	ID              uuid.UUID
	OrganizationID  uuid.UUID
	Name            string `gorm:"not null"`
	GSTNumber       string
	PANNumber       string
	Phone           string
	Email           string
	AddressLine     string
	City            string
	State           string
	Pincode         string
	Country         string
	CreatedAt       time.Time
	UpdatedAt       time.Time
	CreatedByEngine string `gorm:"size:20"` // ACCOUNTING | ORCHESTRATOR | INVENTORY
}

func (Supplier) TableName() string {
	return "core.org_suppliers"
}

/* ---------------------------------- <***> --------------------------------- */

type SupplierLegalEntityLink struct {
	ID               uuid.UUID
	SupplierID       uuid.UUID
	LegalEntityID    uuid.UUID
	PayableAccountID uuid.UUID
	CreatedAt        time.Time
	UpdatedAt        time.Time
}

func (SupplierLegalEntityLink) TableName() string {
	return "core.org_supplier_legal_entity_links"
}

/* ---------------------------------- <***> --------------------------------- */
