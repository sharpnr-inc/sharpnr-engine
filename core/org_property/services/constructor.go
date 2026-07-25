package org_property_services

import (
	org_property_repo "sharpnr-engine/core/org_property/repo"
	dbtx "sharpnr-engine/pkg/dbTx"
	sentinelEngine "sharpnr-engine/sentinel_engine"
	"sharpnr-engine/sharp_mailer_engine"
)

type organizationPropertyService struct {
	orgPropertyRepo org_property_repo.OrganizationPropertyRepositoryInterface
	txManager       dbtx.TxManager
	mailerEngine    *sharp_mailer_engine.SharpMailerEngine
	sentinelEngine  sentinelEngine.SentinelEngineInterface
}

// CreateProperty implements OrganizationPropertyServicesInterfaces.
func (o *organizationPropertyService) CreateProperty() {
	panic("unimplemented")
}

func NewOrganizationPropertyService(
	orgPropertyRepo org_property_repo.OrganizationPropertyRepositoryInterface,
	txManager dbtx.TxManager,
	mailerEngine *sharp_mailer_engine.SharpMailerEngine,
	sentinelAuthzEngine sentinelEngine.SentinelEngineInterface,
) *organizationPropertyService {
	return &organizationPropertyService{
		orgPropertyRepo: orgPropertyRepo,
		txManager:       txManager,
		mailerEngine:    mailerEngine,
		sentinelEngine:  sentinelAuthzEngine,
	}
}
