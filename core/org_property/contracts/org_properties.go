package org_property_contracts

import (
	property_constants "sharpnr-engine/core/org_property/constants"

	"github.com/google/uuid"
)

type OrganizationProperties struct {
	Properties []PropertiesReader
}

type PropertiesReader struct {
	ID     uuid.UUID
	Name   string
	City   string
	State  string
	Domain property_constants.PropertyDomainTypes
	Code   string

	// We can add Country later
}
type PropertyDetailsReader struct {
	Domain      property_constants.PropertyDomainTypes
	Name        string
	Code        string
	AddressLine string
	Location    *any
	City        string
	State       string
	Country     string
	Pincode     string
}
