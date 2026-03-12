package org_setup_wizard_services

import (
	legal_entity_service "sharpnr-engine/core/legal_entity/service"
	org_property_services "sharpnr-engine/core/org_property/services"
	org_setup_wizard_repos "sharpnr-engine/core/org_setup_wizard/repos"
	organization_services "sharpnr-engine/core/organization/services"
)

type orgSetupWizardService struct {
	orgSetWizRepo     org_setup_wizard_repos.OrganizationSetupWizardRepositoryInterface
	orgLegalEntitySvc legal_entity_service.OrgLegalEntityServiceInterface
	orgPropertySvc    org_property_services.OrganizationPropertyServicesInterfaces
	orgService        organization_services.OrganizationServiceInterface
}

func NewOrgSetupWizardService(
	orgSetWizRepo org_setup_wizard_repos.OrganizationSetupWizardRepositoryInterface,
	orgLegalEntitySvc legal_entity_service.OrgLegalEntityServiceInterface,
	orgPropertySvc org_property_services.OrganizationPropertyServicesInterfaces,
	orgService organization_services.OrganizationServiceInterface,
) *orgSetupWizardService {
	return &orgSetupWizardService{
		orgSetWizRepo:     orgSetWizRepo,
		orgLegalEntitySvc: orgLegalEntitySvc,
		orgPropertySvc:    orgPropertySvc,
		orgService:        orgService,
	}
}
