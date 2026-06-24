// Package orgPropertySentinalActions declares the sentinel actions the
// org_property module owns. It is the module's slice of the permission catalog:
// the composition root registers Permissions into sentinel_engine/actions so the
// engine can validate them and the FE can render them in the policy builder.
//
// Action ids are stable Cedar action strings — once a stored managed policy
// references one, renaming it silently breaks that policy. They use the same
// camelCase convention as the rest of the catalog.
package orgPropertySentinalActions

import sentinel_actions "sharpnr-engine/sentinel_engine/actions"

const (
	CreateProperty  sentinel_actions.Name = "createProperty"
	ViewAllProperty sentinel_actions.Name = "viewProperty"
	EditProperty    sentinel_actions.Name = "editProperty"
	AddBusinessOps  sentinel_actions.Name = "addBusinessModulesInProperty"
)

// Permissions is this module's contribution to the catalog. Register it once at
// startup: sentinel_actions.Register(orgPropertySentinalActions.Permissions...).
var Permissions = []sentinel_actions.Permission{
	{Key: ViewAllProperty, Label: "View properties", Description: "View the organization's properties.", Category: sentinel_actions.CategoryProperties},
	{Key: CreateProperty, Label: "Create property", Description: "Add a new property to the organization.", Category: sentinel_actions.CategoryProperties},
	{Key: EditProperty, Label: "Edit property", Description: "Edit a property's details.", Category: sentinel_actions.CategoryProperties},
	{Key: AddBusinessOps, Label: "Add business modules", Description: "Add business-operation modules (hotel, bar, restaurant…) to a property.", Category: sentinel_actions.CategoryProperties},
}
