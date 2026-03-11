package app

import (
	legal_entity_service "sharpnr-engine/core/org_legal_entity/service"
	org_property_services "sharpnr-engine/core/org_property/services"
	org_setup_wizard_repos "sharpnr-engine/core/org_setup_wizard/repos"
	org_setup_wizard_services "sharpnr-engine/core/org_setup_wizard/services"
	organization_services "sharpnr-engine/core/organization/services"

	"gorm.io/gorm"
)

func bootstrapOrgSetupWizard(
	db *gorm.DB,
	orgLegalEntitySvc legal_entity_service.OrgLegalEntityServiceInterface,
	orgPropertySvc org_property_services.OrganizationPropertyServicesInterfaces,
	orgService organization_services.OrganizationServiceInterface,
) org_setup_wizard_services.OrgSetupWizardServiceInterface {
	repo := org_setup_wizard_repos.NewOrgSetupWizardRepository(db)
	return org_setup_wizard_services.NewOrgSetupWizardService(
		repo,
		orgLegalEntitySvc,
		orgPropertySvc,
		orgService,
	)
}
