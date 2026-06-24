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
