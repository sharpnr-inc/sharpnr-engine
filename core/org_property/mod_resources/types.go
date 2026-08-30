// Package orgPropertyResources is the org-property module's mod_resources: it
// describes the Property resource to the Sentinel aggregator, lists an org's
// live properties as policy targets, and answers for the attributes it declares.
package orgPropertyResources

import property_constants "sharpnr-engine/core/org_property/constants"

// Cedar ids, re-exported from the leaf constants package.
//
// They live there rather than here so the property service can use them in its
// own guard calls without importing this package — which imports the service,
// and would cycle. Both sides therefore read the same constant, so enforcement
// and the policy-builder palette cannot drift apart.
const (
	ResourceTypeProperty = property_constants.ResourceTypeProperty

	ActionCanCreateNewProperty   = property_constants.ActionCanCreateNewProperty
	ActionCanViewProperty        = property_constants.ActionCanViewProperty
	ActionCanEditPropertyDetails = property_constants.ActionCanEditPropertyDetails
	ActionCanAddBusinessModules  = property_constants.ActionCanAddBusinessModules
	ActionCanAddInfraModules     = property_constants.ActionCanAddInfraModules
)

// Attribute keys a policy condition may test on a Property. Each must be
// answered by propertyFacts — the aggregator refuses to boot otherwise.
const (
	AttrCode    = "code"
	AttrName    = "name"
	AttrCity    = "city"
	AttrState   = "state"
	AttrCountry = "country"
)
