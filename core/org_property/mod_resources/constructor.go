package orgPropertyResources

import (
	org_property_services "sharpnr-engine/core/org_property/services"
	agg "sharpnr-engine/sentinel_aggregators/contracts"
)

type orgPropertyResources struct {
	propertyService org_property_services.OrganizationPropertyServicesInterfaces
}

// Compile-time guard: the provider must satisfy the aggregator contract.
var _ agg.ResourceProvider = (*orgPropertyResources)(nil)

func NewOrgPropertyResources(propertyService org_property_services.OrganizationPropertyServicesInterfaces) *orgPropertyResources {
	return &orgPropertyResources{propertyService: propertyService}
}
