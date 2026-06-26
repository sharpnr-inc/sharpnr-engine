package orgPropertyresources

import (
	orgPropertySentinalActions "sharpnr-engine/core/org_property/actions"
	sentinel_resources "sharpnr-engine/sentinel_engine/resources"
)

// Descriptor is the Property resource's policy-builder metadata: its actions
// (owned by the org_property/actions package) and the ABAC attributes a policy
// may test in a `when { resource.<attr> ... }` clause.
//
// Attributes are advertised here, but a `when` clause only evaluates if the
// enforcement call site stamps the same key onto the resource entity
// (AuthorizeInput.Resource.Attrs). Keep the two in sync as policies start using
// an attribute.
var Descriptor = sentinel_resources.Descriptor{
	Type:        ResourceType,
	Label:       "Property",
	Description: "An organization's properties (hotels, outlets, sites).",
	Actions:     orgPropertySentinalActions.Actions,
	Attributes: []sentinel_resources.Attribute{
		{Name: "city", Label: "City", Type: "String", Description: "The property's city."},
		{Name: "state", Label: "State", Type: "String", Description: "The property's state/region."},
		{Name: "country", Label: "Country", Type: "String", Description: "The property's country."},
		{Name: "code", Label: "Code", Type: "String", Description: "The property's short code."},
	},
}
