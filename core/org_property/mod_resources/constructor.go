package orgPropertyResources

import (
	"context"
	org_property_services "sharpnr-engine/core/org_property/services"

	"github.com/google/uuid"
)

type orgPropertyResources struct {
	propertyService org_property_services.OrganizationPropertyServicesInterfaces
}

func NewOrgPropertyResources(propertyService org_property_services.OrganizationPropertyServicesInterfaces) *orgPropertyResources {
	return &orgPropertyResources{
		propertyService: propertyService,
	}
}

// this method will return all the properties of an organization
// this is method will used send the list of properties to the frontend for the resources for sentinel
func (s *orgPropertyResources) ListAllProperty(ctx context.Context, orgID, emplID uuid.UUID) ([]string, error) {
	panic("implement me")
}
