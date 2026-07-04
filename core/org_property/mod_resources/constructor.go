// Package orgPropertyResources is the org-property module's mod_resources: it
// describes the Property resource (actions, attributes) to the Sentinel
// aggregator and lists an org's live properties as policy targets.
package orgPropertyResources

import (
	"context"
	"fmt"

	org_property_services "sharpnr-engine/core/org_property/services"
	sentinelAggregatorContracts "sharpnr-engine/sentinel_aggregators/contracts"

	"github.com/google/uuid"
)

type orgPropertyResources struct {
	propertyService org_property_services.OrganizationPropertyServicesInterfaces
}

// Compile-time guard: the provider must satisfy the aggregator contract.
var _ sentinelAggregatorContracts.ResourceProvider = (*orgPropertyResources)(nil)

func NewOrgPropertyResources(propertyService org_property_services.OrganizationPropertyServicesInterfaces) *orgPropertyResources {
	return &orgPropertyResources{
		propertyService: propertyService,
	}
}

// ListInstances implements [sentinelAggregatorContracts.ResourceProvider]: the
// org's properties as selectable policy targets for the frontend builder.
func (s *orgPropertyResources) ListInstances(ctx context.Context, orgID, emplID uuid.UUID, resourceType string) ([]sentinelAggregatorContracts.ResourceInstance, error) {
	if resourceType != ResourceTypeProperty {
		return nil, fmt.Errorf("org-property resources: unknown resource type %q", resourceType)
	}

	properties, err := s.propertyService.GetPropertiesOfOrg(ctx, orgID)
	if err != nil {
		return nil, err
	}

	instances := make([]sentinelAggregatorContracts.ResourceInstance, 0, len(properties))
	for _, p := range properties {
		label := p.Name
		if p.Code != "" {
			label = fmt.Sprintf("%s (%s)", p.Name, p.Code)
		}
		instances = append(instances, sentinelAggregatorContracts.ResourceInstance{
			ID:    p.ID.String(),
			Label: label,
		})
	}
	return instances, nil
}
