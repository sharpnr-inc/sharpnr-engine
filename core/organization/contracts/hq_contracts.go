package organization_contracts

import "github.com/google/uuid"

type OrganizationHQ struct {
	OrgID        uuid.UUID
	Address_line string
	Country      string
	City         string
	State        string
	Pincode      string
	Location     OrganizationHQGeoLocation
}

type OrganizationHQGeoLocation struct {
	Longitude *float64
	Latitude  *float64
	Accuracy  *int
}
