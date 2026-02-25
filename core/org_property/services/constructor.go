package org_property_services

import (
	dbtx "sharpnr-engine/core/dbTx"
	org_property_repo "sharpnr-engine/core/org_property/repo"
)

type organizationPropertyService struct {
	orgPropertyRepo org_property_repo.OrganizationPropertyRepositoryInterface
	txManager       dbtx.TxManager
}

// CreateProperty implements OrganizationPropertyServicesInterfaces.
func (o *organizationPropertyService) CreateProperty() {
	panic("unimplemented")
}

func NewOrganizationPropertyService(
	orgPropertyRepo org_property_repo.OrganizationPropertyRepositoryInterface,
	txManager dbtx.TxManager,
) *organizationPropertyService {
	return &organizationPropertyService{
		orgPropertyRepo: orgPropertyRepo,
		txManager:       txManager,
	}
}
