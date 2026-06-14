package org_property_contracts

import (
	property_constants "sharpnr-engine/core/org_property/constants"

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
