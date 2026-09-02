package app

import (
	orgDepartmentResources "sharpnr-engine/core/departments/mod_resources"
	orgDepartmentService "sharpnr-engine/core/departments/service"
	orgPropertyResources "sharpnr-engine/core/org_property/mod_resources"
	org_property_services "sharpnr-engine/core/org_property/services"
	sentinelOwnResources "sharpnr-engine/sentinel/mod_resources"
	sentinelservice "sharpnr-engine/sentinel/service"
	sentinelAggregatorService "sharpnr-engine/sentinel_aggregators/service"
)

// bootstrapSentinelAggregator registers every module's mod_resources with the
// resources aggregator — the catalog the frontend policy builder reads. New
// modules plug in here by adding their provider.
func bootstrapSentinelAggregator(
	orgPropertyService org_property_services.OrganizationPropertyServicesInterfaces,
	sentinelService sentinelservice.SentinelServiceInterface,
	orgDepartmentService orgDepartmentService.OrgDepartmentServiceInterface,
) sentinelAggregatorService.SentinelAggregatorServiceInterface {
	return sentinelAggregatorService.NewSentinelAggregatorService(
		orgPropertyResources.NewOrgPropertyResources(orgPropertyService),
		sentinelOwnResources.NewSentinelOwnResources(sentinelService),
		orgDepartmentResources.NewOrgDepartmentResources(orgDepartmentService),
	)
}
