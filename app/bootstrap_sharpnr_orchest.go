package app

import (
	hospitalitydomain "sharpnr-engine/app/hospitality_domain"
	dbtx "sharpnr-engine/core/dbTx"
	org_legal_entity_service "sharpnr-engine/core/org_legal_entity/service"
	org_property_services "sharpnr-engine/core/org_property/services"
	org_setup_wizard_services "sharpnr-engine/core/org_setup_wizard/services"
	organization_services "sharpnr-engine/core/organization/services"
	orchestrator_service "sharpnr-engine/orchestrator/service"
	accounting_service "sharpnr-engine/sharp_services/accounting_engine/services"
)

func bootstrapSharpnrOrchestrator(
	organizationService organization_services.OrganizationServiceInterface,
	orgPropertyService org_property_services.OrganizationPropertyServicesInterfaces,
	orgSetupWizService org_setup_wizard_services.OrgSetupWizardServiceInterface,
	hospitalityDomainSvc *hospitalitydomain.HospitalityDomain,
	accountingEngineService accounting_service.AccountingServiceInterface,
	legalEntityService org_legal_entity_service.OrgLegalEntityServiceInterface,
	txManager dbtx.TxManager,
) orchestrator_service.OrchestratorServiceInterface {
	return orchestrator_service.NewSharpnrOrchestratorService(
		organizationService,
		orgPropertyService,
		orgSetupWizService,
		hospitalityDomainSvc,
		accountingEngineService,
		legalEntityService,
		txManager,
	)
}
