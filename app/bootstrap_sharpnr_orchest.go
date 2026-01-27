package app

import (
	hospitalitydomain "sharpnr-engine/app/hospitality_domain"
	orchestrator_service "sharpnr-engine/core/orchestrator/service"
	org_legal_entity_service "sharpnr-engine/core/org_legal_entity/service"
	org_property_services "sharpnr-engine/core/org_property/services"
	org_setup_wizard_services "sharpnr-engine/core/org_setup_wizard/services"
	organization_services "sharpnr-engine/core/organization/services"
	accounting_service "sharpnr-engine/sharp_services/accounting_engine/services"
)

func bootstrapSharpnrOrchestrator(
	organizationService organization_services.OrganizationServiceInterface,
	orgPropertyService org_property_services.OrganizationPropertyServicesInterfaces,
	orgSetupWizService org_setup_wizard_services.OrgSetupWizardServiceInterface,
	hospitalityDomainSvc *hospitalitydomain.HospitalityDomain,
	accountingEngineService accounting_service.AccountingServiceInterface,
	legalEntityService org_legal_entity_service.OrgLegalEntityServiceInterface,
) orchestrator_service.OrchestratorServiceInterface {
	return orchestrator_service.NewSharpnrOrchestratorService(
		organizationService,
		orgPropertyService,
		orgSetupWizService,
		hospitalityDomainSvc,
		accountingEngineService,
		legalEntityService,
	)
}
