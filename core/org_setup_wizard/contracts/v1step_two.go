package org_setup_wizard_contracts

import "github.com/google/uuid"

type PropertySetupInput struct {
	StepID       uuid.UUID
	OrgID        uuid.UUID
	Name         string
	PropertyCode string
	AddressLine  string
	City         string
	State        string
	Country      string
	Pincode      string
}
