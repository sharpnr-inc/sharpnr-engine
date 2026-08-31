package org_property_contracts

import (
	"github.com/google/uuid"
)

type PropertiesReader struct {
	ID          uuid.UUID
	Name        string
	AddressLine string
	City        string
	State       string
	Code        string
	Country     string
}

type PropertyDetailsReader struct {
	Name        string
	Code        string
	AddressLine string
	Location    *any
	City        string
	State       string
	Country     string
	Pincode     string
}

// PropertyFacts is what Sentinel needs to know about a property to authorize
// against it: who owns it, and the attribute values a policy condition may test.
//
// Separate from PropertiesReader because it answers a different question and is
// read on the authorization path — adding a display field to the reader must not
// silently widen what every authorize call loads.
type PropertyFacts struct {
	ID             uuid.UUID
	OrganizationID uuid.UUID

	Code    string
	Name    string
	City    string
	State   string
	Country string
}
