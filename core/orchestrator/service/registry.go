package orchestrator_service

import (
	org_property_services "sharpnr-engine/core/org_property/services"
	org_setup_wizard_services "sharpnr-engine/core/org_setup_wizard/services"
	organization_services "sharpnr-engine/core/organization/services"
)

type OrchestratorServiceRegistry struct {
	Organization organization_services.OrganizationServiceInterface
	OrgProperty  org_property_services.OrganizationPropertyServicesInterfaces
	SetupWizard  org_setup_wizard_services.OrgSetupWizardServiceInterface
}
