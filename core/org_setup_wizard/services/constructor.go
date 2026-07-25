package org_setup_wizard_services

import (
	coreEmployeeServices "sharpnr-engine/core/employees/services"
	legal_entity_service "sharpnr-engine/core/legal_entity/service"
	org_property_services "sharpnr-engine/core/org_property/services"
	org_setup_wizard_repos "sharpnr-engine/core/org_setup_wizard/repos"
	organization_services "sharpnr-engine/core/organization/services"
	dbtx "sharpnr-engine/pkg/dbTx"
	"sharpnr-engine/sharp_mailer_engine"
)

type orgSetupWizardService struct {
	orgSetWizRepo       org_setup_wizard_repos.OrganizationSetupWizardRepositoryInterface
	orgLegalEntitySvc   legal_entity_service.OrgLegalEntityServiceInterface
	orgPropertySvc      org_property_services.OrganizationPropertyServicesInterfaces
	orgService          organization_services.OrganizationServiceInterface
	coreEmployeeService coreEmployeeServices.EmployeeSetupWizServiceInterface
	txManager           dbtx.TxManager
	mailerEngine        *sharp_mailer_engine.SharpMailerEngine
}

func NewOrgSetupWizardService(
	orgSetWizRepo org_setup_wizard_repos.OrganizationSetupWizardRepositoryInterface,
	orgLegalEntitySvc legal_entity_service.OrgLegalEntityServiceInterface,
	orgPropertySvc org_property_services.OrganizationPropertyServicesInterfaces,
	orgService organization_services.OrganizationServiceInterface,
	coreEmployeeService coreEmployeeServices.EmployeeSetupWizServiceInterface,
	txManager dbtx.TxManager,
	mailerEngine *sharp_mailer_engine.SharpMailerEngine,
) *orgSetupWizardService {
	return &orgSetupWizardService{
		orgSetWizRepo:       orgSetWizRepo,
		orgLegalEntitySvc:   orgLegalEntitySvc,
		orgPropertySvc:      orgPropertySvc,
		orgService:          orgService,
		coreEmployeeService: coreEmployeeService,
		txManager:           txManager,
		mailerEngine:        mailerEngine,
	}
}
