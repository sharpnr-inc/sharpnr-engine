package org_property_services

import (
	"context"
	property_constants "sharpnr-engine/core/org_property/constants"
	org_property_contracts "sharpnr-engine/core/org_property/contracts"

	"github.com/google/uuid"
)

type OrganizationPropertyServicesInterfaces interface {
	CreatePropertyDefault(ctx context.Context, input org_property_contracts.CreateOrgPropertyInput) error
	GetPropertiesOfOrg(ctx context.Context, orgID uuid.UUID) (*org_property_contracts.OrganizationProperties, error)
	GetPropertyDetailsByID(ctx context.Context, orgID uuid.UUID, propertyID uuid.UUID) (*org_property_contracts.PropertyDetailsReader, error)
	AddDomain(ctx context.Context, orgID uuid.UUID, propertyID uuid.UUID, Domain property_constants.PropertyDomainTypes) error
}
