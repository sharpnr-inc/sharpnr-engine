package org_property_services

import "context"

type OrganizationPropertyServicesInterfaces interface {
	CreatePropertyDefault(ctx context.Context) error
}
