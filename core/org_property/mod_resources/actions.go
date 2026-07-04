package orgPropertyResources

import sentinelAggregatorContracts "sharpnr-engine/sentinel_aggregators/contracts"

// Cedar action ids for the Property resource. Services embed these in
// authorization checks and the policy builder writes them into policy text
// (`action == Action::"canViewProperty"`), so they must stay stable once
// released.
const (
	ActionCanCreateNewProperty   = "canCreateNewProperty"
	ActionCanViewProperty        = "canViewProperty"
	ActionCanEditPropertyDetails = "canEditPropertyDetails"
	ActionCanAddBusinessModules  = "canAddBusinessModules"
	ActionCanAddInfraModules     = "canAddInfraModules"
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
