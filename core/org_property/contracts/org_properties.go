package org_property_contracts

import "github.com/google/uuid"

type OrganizationProperties struct {
	Properties []PropertiesReader
}

type PropertiesReader struct {
	ID     uuid.UUID
	Name   string
	City   string
	State  string
	Domain string
	Code   string

	// We can add Country later
}
type PropertyDetailsReader struct {
	Domain      string
	Name        string
	Code        string
	AddressLine string
	Location    *any
	City        string
	State       string
	Country     string
	Pincode     string
}
