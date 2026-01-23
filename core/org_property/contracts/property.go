package org_property_contracts

import "github.com/google/uuid"

type CreateOrgPropertyInput struct {
	OrgID uuid.UUID
	Name  string
}

type PropertyAddress struct {
}

type PropertyGeoLocation struct {
}
