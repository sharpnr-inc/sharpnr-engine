package orgPropertyResources

import (
	property_constants "sharpnr-engine/core/org_property/constants"
	sentinelAggregatorContracts "sharpnr-engine/sentinel_aggregators/contracts"
)

// ResourceTypeProperty is the Cedar entity type of an organization property.
// Engine callers build Resource{Type: ResourceTypeProperty, ID: <property id>}.
// Defined in property_constants; re-exported here for the aggregator side.
const ResourceTypeProperty = property_constants.ResourceTypeProperty

// Resources implements [sentinelAggregatorContracts.ResourceProvider].
func (s *orgPropertyResources) Resources() []sentinelAggregatorContracts.ResourceDescriptor {
	return []sentinelAggregatorContracts.ResourceDescriptor{
		{
			Type:        ResourceTypeProperty,
			Label:       "Property",
			Description: "A physical property (hotel, restaurant site, ...) of the organization",
			Actions:     propertyActions(),
			Attributes:  propertyAttributes(),
		},
	}
}
