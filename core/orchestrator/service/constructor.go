package orchestrator_service

import (
	org_property_services "sharpnr-engine/core/org_property/services"
	org_setup_wizard_services "sharpnr-engine/core/org_setup_wizard/services"
	organization_services "sharpnr-engine/core/organization/services"
)

type orchestratorService struct {
	organizationService organization_services.OrganizationServiceInterface
	orgPropertyService  org_property_services.OrganizationPropertyServicesInterfaces
	orgSetupWizService  org_setup_wizard_services.OrgSetupWizardServiceInterface
}

func NewSharpnrOrchestratorService(
	organizationService organization_services.OrganizationServiceInterface,
	orgPropertyService org_property_services.OrganizationPropertyServicesInterfaces,
	orgSetupWizService org_setup_wizard_services.OrgSetupWizardServiceInterface,
) *orchestratorService {
	return &orchestratorService{
		organizationService: organizationService,
		orgPropertyService:  orgPropertyService,
		orgSetupWizService:  orgSetupWizService,
	}
}
