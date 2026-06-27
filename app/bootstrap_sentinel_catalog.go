package app

import (
	orgPropertyresources "sharpnr-engine/core/org_property/resources"
	org_property_services "sharpnr-engine/core/org_property/services"
	sentinel_resources "sharpnr-engine/sentinel_engine/resources"
)

// registerSentinelCatalog wires every module's resource into the sentinel policy
// builder: the static Descriptor (actions + attributes, which also seeds the
// engine's action-validity set) and the dynamic instance Provider. Call once at
// startup, after the owning module's service exists. Add a block here as each new
// module gains a resource.
func registerSentinelCatalog(
	orgPropertyService org_property_services.OrganizationPropertyServicesInterfaces,
) {
	// Property
	sentinel_resources.RegisterDescriptor(orgPropertyresources.Descriptor)
	sentinel_resources.RegisterProvider(orgPropertyresources.NewOrgPropertyResources(orgPropertyService))
}
