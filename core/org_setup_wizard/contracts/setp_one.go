package org_setup_wizard_contracts

import "github.com/google/uuid"

type StepOneOrganizationHQ struct {
	OrgID        uuid.UUID
	Address_line string
	Country      string
	City         string
	State        string
	Pincode      string
	Location     StepOneOrganizationHQGeoLocation
}

type StepOneOrganizationHQGeoLocation struct {
	Longitude *float64
	Latitude  *float64
	Accuracy  *int
}
