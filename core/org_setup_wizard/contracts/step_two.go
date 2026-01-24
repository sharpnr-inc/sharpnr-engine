package org_setup_wizard_contracts

import "github.com/google/uuid"

type PropertySetupInput struct {
	OrgID       uuid.UUID
	Name        string
	AddressLine string
	City        string
	State       string
	Country     string
	Pincode     string
}
