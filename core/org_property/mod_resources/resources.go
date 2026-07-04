package orgPropertyResources

import sentinelAggregatorContracts "sharpnr-engine/sentinel_aggregators/contracts"

// ResourceTypeProperty is the Cedar entity type of an organization property.
// Engine callers build Resource{Type: ResourceTypeProperty, ID: <property id>}.
const ResourceTypeProperty = "Property"

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
