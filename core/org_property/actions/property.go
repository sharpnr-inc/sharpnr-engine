// Package orgPropertySentinalActions declares the sentinel actions the
// org_property module owns. These are the actions that apply to the "Property"
// resource — the org_property/resources Descriptor folds them in, the composition
// root registers the descriptor, and the FE renders them in the policy builder.
//
// Action ids are stable Cedar action strings — once a stored policy references
// one, renaming it silently breaks that policy. They use camelCase.
package orgPropertySentinalActions

import (
	sentinel_actions "sharpnr-engine/sentinel/actions"
	sentinel_resources "sharpnr-engine/sentinel/resources"
)

const (
	CreateProperty  sentinel_actions.Name = "createProperty"
	ViewAllProperty sentinel_actions.Name = "viewProperty"
	EditProperty    sentinel_actions.Name = "editProperty"
	AddBusinessOps  sentinel_actions.Name = "addBusinessModulesInProperty"
	AddInfraModules sentinel_actions.Name = "addInfraModulesInProperty"
)

// Actions is this module's contribution to the Property resource descriptor.
var Actions = []sentinel_resources.Action{
	{Key: ViewAllProperty, Label: "View properties", Description: "View the organization's properties."},
	{Key: CreateProperty, Label: "Create property", Description: "Add a new property to the organization."},
	{Key: EditProperty, Label: "Edit property", Description: "Edit a property's details."},
	{Key: AddBusinessOps, Label: "Add business modules", Description: "Add business-operation modules (hotel, bar, restaurant…) to a property."},
	{Key: AddInfraModules, Label: "Add infra modules", Description: "Add infrastructure modules to a property."},
}
