package org_property_contracts

import "github.com/google/uuid"

type CreateOrgPropertyInput struct {
	OrgID       uuid.UUID
	Name        string
	Code        string
	AddressLine string
	City        string
	State       string
	Country     string
	Pincode     string
}

type PropertyGeoLocation struct {
}
