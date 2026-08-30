package orgPropertyResources

import (
	agg "sharpnr-engine/sentinel_aggregators/contracts"
	aggTypes "sharpnr-engine/sentinel_aggregators/types"
)

// Resources implements [agg.ResourceProvider].
//
// Property is ORGANIZATION-LEVEL: it is the container other modules are housed
// in, and it sits directly under the organization. So there is no HousedIn here
// — the engine stamps the Organization edge itself from the request's org.
// Hotels, restaurants and kitchens declare HousedIn(ResourceTypeProperty, ...),
// which is what makes scoping a role to a property reach everything inside it.
func (s *orgPropertyResources) Resources() []agg.ResourceRegistration {
	return []agg.ResourceRegistration{
		agg.Standard(ResourceTypeProperty, "Properties").
			Category("Operations").
			Describe("Buildings and sites the organization manages").

			// Standard() generates canViewProperty, canCreateProperty,
			// canEditProperty and canDeleteProperty. Two of those overlap with
			// ids this module already shipped, so the shipped spellings win —
			// they are embedded in stored policy text and renaming one silently
			// breaks every policy that mentions it.
			WithActions(
				agg.Action(ActionCanCreateNewProperty, "Add a new property").
					Describe("Register a new building or site"),
				agg.Action(ActionCanEditPropertyDetails, "Edit property details").
					Describe("Change name, address and other details").
					OnInstance(),
				agg.Action(ActionCanAddBusinessModules, "Set up business services").
					Describe("Turn on hotel, restaurant and similar services").
					OnInstance(),
				agg.Action(ActionCanAddInfraModules, "Set up facilities").
					Describe("Turn on infrastructure and facility management").
					OnInstance(),
			).

			// Declaring an attribute obliges this module to answer for it. That
			// is the whole point of the contract: the previous version advertised
			// these to the policy builder with nothing anywhere to stamp them, so
			// `resource.city == "Mumbai"` compiled, stored, displayed as active —
			// and matched nothing, because Cedar skips a policy whose attribute
			// is missing.
			WithAttributes(
				agg.Attr(AttrCity, aggTypes.AttrString, "City").
					Describe("City the property is in").
					ResolveValues(s.attributeValues(AttrCity)),
				agg.Attr(AttrState, aggTypes.AttrString, "State").
					Describe("State the property is in").
					ResolveValues(s.attributeValues(AttrState)),
				agg.Attr(AttrCountry, aggTypes.AttrString, "Country").
					Describe("Country the property is in").
					ResolveValues(s.attributeValues(AttrCountry)),
				agg.Attr(AttrCode, aggTypes.AttrString, "Property code").
					Describe("The organization's own code for the property").
					ResolveValues(s.attributeValues(AttrCode)),
				// Name is deliberately free-form: it is the one attribute where
				// a picker of every property name would just duplicate the
				// instance picker, which is the better control for "this
				// specific property".
				agg.Attr(AttrName, aggTypes.AttrString, "Property name").
					Describe("Prefer scoping to specific properties instead"),
			).
			ResolveFacts(s.propertyFacts).
			ListInstances(s.listProperties).
			Build(),
	}
}
