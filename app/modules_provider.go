package app

import (
	hospitalitydomain "sharpnr-engine/app/hospitality_domain"
	org_customer_service "sharpnr-engine/core/customer/service"
	legal_entity_service "sharpnr-engine/core/legal_entity/service"
	org_property_services "sharpnr-engine/core/org_property/services"
	org_setup_wizard_services "sharpnr-engine/core/org_setup_wizard/services"
	org_signup_services "sharpnr-engine/core/org_signup/services"
	organization_services "sharpnr-engine/core/organization/services"
	sharpnr_identity_auth_services "sharpnr-engine/core/sharpnr_identity/services"
	employee_services "sharpnr-engine/modules/hr/employees/services"
	orchestrator_service "sharpnr-engine/orchestrator/service"
	accounting_service "sharpnr-engine/service_engines/accounting_engine/services"
	tax_engine_service "sharpnr-engine/service_engines/tax_engine/service"
	sse_engine "sharpnr-engine/sse"
)

type ModulesProvider interface {
	OrganizationService() organization_services.OrganizationServiceInterface
	SharpnrOrchestrator() orchestrator_service.OrchestratorServiceInterface
	OrgSetupWiz() org_setup_wizard_services.OrgSetupWizardServiceInterface
	Hospitality() *hospitalitydomain.HospitalityDomain
	OrgPropertyService() org_property_services.OrganizationPropertyServicesInterfaces

	OrgSignUpService() org_signup_services.OrganizationSignUpServiceInterfaces
	SharpAuthService() sharpnr_identity_auth_services.SharpnrIdentityAuthServiceInterface

	EmployeeService() employee_services.EmployeeServiceInterface
	AccountingService() accounting_service.AccountingServiceInterface
	SseEngine() *sse_engine.StreamManager
	LegalEntityService() legal_entity_service.OrgLegalEntityServiceInterface
	TaxEngineService() tax_engine_service.TaxEngineServiceInterface
	CustomerEngineService() org_customer_service.OrgCustomerServiceInterface
}
