package org_property_services

import org_property_repo "sharpnr-engine/core/org_property/repo"

type organizationPropertyService struct {
	orgPropertyRepo org_property_repo.OrganizationPropertyRepositoryInterface
}

// CreateProperty implements OrganizationPropertyServicesInterfaces.
func (o *organizationPropertyService) CreateProperty() {
	panic("unimplemented")
}

func NewOrganizationPropertyService(
	orgPropertyRepo org_property_repo.OrganizationPropertyRepositoryInterface,
) *organizationPropertyService {
	return &organizationPropertyService{
		orgPropertyRepo: orgPropertyRepo,
	}
}
