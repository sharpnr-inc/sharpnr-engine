package app

import (
	hospitalitydomain "sharpnr-engine/app/hospitality_domain"
	hospitality_orchestration_service "sharpnr-engine/business_ops/hospitality/orchestration/service"
	org_customer_service "sharpnr-engine/core/customer/service"
	coreEmployeeServices "sharpnr-engine/core/employees/services"
	legal_entity_service "sharpnr-engine/core/legal_entity/service"
	org_property_services "sharpnr-engine/core/org_property/services"
	org_setup_wizard_services "sharpnr-engine/core/org_setup_wizard/services"
	org_signup_services "sharpnr-engine/core/org_signup/services"
	organization_services "sharpnr-engine/core/organization/services"
	sharpnr_identity_auth_services "sharpnr-engine/core/sharpnr_identity/services"
	hrmsengine "sharpnr-engine/hrms_engine/engine"

	orchestrator_service "sharpnr-engine/orchestrator/service"
	payment_engine_service "sharpnr-engine/payment_engine/services"
	accounting_service "sharpnr-engine/sharp_accounting_engine/accounting_engine/services"
	utilitiesApi_service "sharpnr-engine/utilities_apis/service"

	sse_engine "sharpnr-engine/sse_engine"
)

type ModulesProvider interface {
	OrganizationService() organization_services.OrganizationServiceInterface
	SharpnrOrchestrator() orchestrator_service.OrchestratorServiceInterface
	OrgSetupWiz() org_setup_wizard_services.OrgSetupWizardV1_ServiceInterface

	Hospitality() *hospitalitydomain.HospitalityDomain
	HospitalityOrchestration() hospitality_orchestration_service.OrchestrationServiceInterface

	OrgPropertyService() org_property_services.OrganizationPropertyServicesInterfaces

	OrgSignUpService() org_signup_services.OrganizationSignUpServiceInterfaces
	SharpAuthService() sharpnr_identity_auth_services.SharpnrIdentityAuthServiceInterface

	EmployeesIdentityService() coreEmployeeServices.EmployeesIdentityServiceInterface
	EmployeesHRService() coreEmployeeServices.EmployeesHRServiceInterface
	HRMS() *hrmsengine.HRMSEngine
	AccountingService() accounting_service.AccountingServiceInterface
	SseEngine() *sse_engine.StreamManager
	LegalEntityService() legal_entity_service.OrgLegalEntityServiceInterface
	CustomerEngineService() org_customer_service.OrgCustomerServiceInterface
	PaymentEngineService() payment_engine_service.PaymentServiceInterface
	UtilitiesApiService() utilitiesApi_service.UtilitiesAPIServiceInterface
}
