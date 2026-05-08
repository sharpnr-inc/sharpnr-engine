package org_property_services

import (
	dbtx "sharpnr-engine/core/dbTx"
	org_property_repo "sharpnr-engine/core/org_property/repo"
	"sharpnr-engine/sharp_mailer_engine"
)

type organizationPropertyService struct {
	orgPropertyRepo org_property_repo.OrganizationPropertyRepositoryInterface
	txManager       dbtx.TxManager
	mailerEngine    *sharp_mailer_engine.SharpMailerEngine
}

// CreateProperty implements OrganizationPropertyServicesInterfaces.
func (o *organizationPropertyService) CreateProperty() {
	panic("unimplemented")
}

func NewOrganizationPropertyService(
	orgPropertyRepo org_property_repo.OrganizationPropertyRepositoryInterface,
	txManager dbtx.TxManager,
	mailerEngine *sharp_mailer_engine.SharpMailerEngine,
) *organizationPropertyService {
	return &organizationPropertyService{
		orgPropertyRepo: orgPropertyRepo,
		txManager:       txManager,
		mailerEngine:    mailerEngine,
	}
}
