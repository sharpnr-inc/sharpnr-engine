package orgPropertyResources

import (
	property_constants "sharpnr-engine/core/org_property/constants"
	sentinelAggregatorContracts "sharpnr-engine/sentinel_aggregators/contracts"
)

// Cedar action ids for the Property resource. Defined in property_constants (a
// leaf package) so the property service can use them in engine checks without
// importing this package; re-exported here for the aggregator/orchestrator side.
const (
	ActionCanCreateNewProperty   = property_constants.ActionCanCreateNewProperty
	ActionCanViewProperty        = property_constants.ActionCanViewProperty
	ActionCanEditPropertyDetails = property_constants.ActionCanEditPropertyDetails
	ActionCanAddBusinessModules  = property_constants.ActionCanAddBusinessModules
	ActionCanAddInfraModules     = property_constants.ActionCanAddInfraModules
)

// propertyActions is the action palette for the Property resource.
func propertyActions() []sentinelAggregatorContracts.ActionDescriptor {
	return []sentinelAggregatorContracts.ActionDescriptor{
		{
			Name:        ActionCanCreateNewProperty,
			Label:       "Create property",
			Description: "Create a new property in the organization",
		},
		{
			Name:             ActionCanViewProperty,
			Label:            "View property",
			Description:      "View properties and their details",
			RequiresInstance: true,
		},
		{
			Name:             ActionCanEditPropertyDetails,
			Label:            "Edit property details",
			Description:      "Edit a property's details",
			RequiresInstance: true,
		},
		{
			Name:             ActionCanAddBusinessModules,
			Label:            "Add business modules",
			Description:      "Enable business operation modules (hotel, restaurant, ...) on a property",
			RequiresInstance: true,
		},
		{
			Name:             ActionCanAddInfraModules,
			Label:            "Add infra modules",
			Description:      "Enable infrastructure modules on a property",
			RequiresInstance: true,
		},
	}
}
