package app

import (
	hospitalitydomain "sharpnr-engine/app/hospitality_domain"
	orchestrator_service "sharpnr-engine/core/orchestrator/service"
	org_property_services "sharpnr-engine/core/org_property/services"
	org_setup_wizard_services "sharpnr-engine/core/org_setup_wizard/services"
	organization_services "sharpnr-engine/core/organization/services"
)

func bootstrapSharpnrOrchestrator(
	organizationService organization_services.OrganizationServiceInterface,
	orgPropertyService org_property_services.OrganizationPropertyServicesInterfaces,
	orgSetupWizService org_setup_wizard_services.OrgSetupWizardServiceInterface,
	hospitalityDomainSvc *hospitalitydomain.HospitalityDomain,
) orchestrator_service.OrchestratorServiceInterface {
	return orchestrator_service.NewSharpnrOrchestratorService(
		organizationService,
		orgPropertyService,
		orgSetupWizService,
		hospitalityDomainSvc,
	)
}
