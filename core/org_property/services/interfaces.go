package org_property_services

import (
	"context"
	org_property_contracts "sharpnr-engine/core/org_property/contracts"
)

type OrganizationPropertyServicesInterfaces interface {
	CreatePropertyDefault(ctx context.Context, input org_property_contracts.CreateOrgPropertyInput) error
}
