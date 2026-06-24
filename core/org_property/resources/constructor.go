// Package orgPropertyresources implements the sentinel resource Provider for
// properties. It is the dynamic half of the policy builder: the action catalog
// is static, but the properties a policy can be scoped to are fetched live for
// the org. Registered at startup so sentinel_resources.List(ctx, "Property", org)
// dispatches here.
package orgPropertyresources

import (
	"context"

	org_property_services "sharpnr-engine/core/org_property/services"
	sentinel_resources "sharpnr-engine/sentinel_engine/resources"

	"github.com/google/uuid"
)

// ResourceType is the Cedar entity type properties are exposed as. Enforcement
// call sites and the generator must use this same string.
const ResourceType = "Property"

type orgPropertyResources struct {
	svc org_property_services.OrganizationPropertyServicesInterfaces
}

// NewOrgPropertyResources builds the provider over the property service (not the
// repo) so it respects the module's service boundary.
func NewOrgPropertyResources(svc org_property_services.OrganizationPropertyServicesInterfaces) *orgPropertyResources {
	return &orgPropertyResources{svc: svc}
}

// Type implements sentinel_resources.Provider.
func (p *orgPropertyResources) Type() string { return ResourceType }

// List returns the org's properties as policy-builder refs (id + name). Tenant
// scoping is the service's responsibility (it queries by orgID).
func (p *orgPropertyResources) List(ctx context.Context, orgID uuid.UUID) ([]sentinel_resources.Ref, error) {
	props, err := p.svc.GetPropertiesOfOrg(ctx, orgID)
	if err != nil {
		return nil, err
	}
	refs := make([]sentinel_resources.Ref, 0, len(props))
	for _, pr := range props {
		refs = append(refs, sentinel_resources.Ref{ID: pr.ID, Type: ResourceType, Label: pr.Name})
	}
	return refs, nil
}

// compile-time guard: the provider must satisfy the interface.
var _ sentinel_resources.Provider = (*orgPropertyResources)(nil)
