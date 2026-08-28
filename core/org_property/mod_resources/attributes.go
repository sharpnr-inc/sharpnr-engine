package orgPropertyResources

import (
	sentinelAggregatorContracts "sharpnr-engine/sentinel_aggregators/contracts"
	sentinelAggregatorTypes "sharpnr-engine/sentinel_aggregators/types"
)

// propertyAttributes lists the resource attributes a policy condition can test
// on a Property (`resource.city == "Mumbai"`). They must match the attribute
// names the engine callers pass as ResourceAttributes.
func propertyAttributes() []sentinelAggregatorContracts.AttributeDescriptor {
	return []sentinelAggregatorContracts.AttributeDescriptor{
		{Name: "code", Type: sentinelAggregatorTypes.AttrString, Description: "Property code"},
		{Name: "name", Type: sentinelAggregatorTypes.AttrString, Description: "Property name"},
		{Name: "city", Type: sentinelAggregatorTypes.AttrString, Description: "City the property is in"},
		{Name: "state", Type: sentinelAggregatorTypes.AttrString, Description: "State the property is in"},
		{Name: "country", Type: sentinelAggregatorTypes.AttrString, Description: "Country the property is in"},
	}
}
