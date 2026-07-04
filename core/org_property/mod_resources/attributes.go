package orgPropertyResources

import sentinelAggregatorContracts "sharpnr-engine/sentinel_aggregators/contracts"

// propertyAttributes lists the resource attributes a policy condition can test
// on a Property (`resource.city == "Mumbai"`). They must match the attribute
// names the engine callers pass as ResourceAttributes.
func propertyAttributes() []sentinelAggregatorContracts.AttributeDescriptor {
	return []sentinelAggregatorContracts.AttributeDescriptor{
		{Name: "code", Type: "string", Description: "Property code"},
		{Name: "name", Type: "string", Description: "Property name"},
		{Name: "city", Type: "string", Description: "City the property is in"},
		{Name: "state", Type: "string", Description: "State the property is in"},
		{Name: "country", Type: "string", Description: "Country the property is in"},
	}
}
