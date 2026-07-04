package property_constants

// Cedar ids of the org-property module's protectable resource. They live in
// this leaf package (not mod_resources) so the property service can reference
// them in engine checks without importing mod_resources — which imports the
// service and would cycle. mod_resources builds its aggregator descriptors from
// these same constants, so enforcement and the policy-builder palette can't
// drift apart. They are embedded in stored policy text, so they must stay
// stable once released.
const (
	// ResourceTypeProperty is the Cedar entity type of an organization property.
	ResourceTypeProperty = "Property"

	ActionCanCreateNewProperty   = "canCreateNewProperty"
	ActionCanViewProperty        = "canViewProperty"
	ActionCanEditPropertyDetails = "canEditPropertyDetails"
	ActionCanAddBusinessModules  = "canAddBusinessModules"
	ActionCanAddInfraModules     = "canAddInfraModules"
)
