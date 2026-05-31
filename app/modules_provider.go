package app

import (
	hospitalitydomain "sharpnr-engine/app/hospitality_domain"
	hospitality_orchestration_service "sharpnr-engine/business_ops/hospitality/orchestration/service"
	org_customer_service "sharpnr-engine/core/customer/service"
	legal_entity_service "sharpnr-engine/core/legal_entity/service"
	org_property_services "sharpnr-engine/core/org_property/services"
	org_setup_wizard_services "sharpnr-engine/core/org_setup_wizard/services"
	org_signup_services "sharpnr-engine/core/org_signup/services"
	organization_services "sharpnr-engine/core/organization/services"
	sharpnr_identity_auth_services "sharpnr-engine/core/sharpnr_identity/services"
	orchestrator_service "sharpnr-engine/orchestrator/service"
	payment_engine_service "sharpnr-engine/payment_engine/services"
	accounting_service "sharpnr-engine/service_engines/accounting_engine/services"
	employee_services "sharpnr-engine/sharp_hrms_engine/employees/services"

	sse_engine "sharpnr-engine/sse"
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

	EmployeesIdentityService() employee_services.EmployeesIdentityServiceInterface
	EmployeesHRService() employee_services.EmployeesHRServiceInterface
	AccountingService() accounting_service.AccountingServiceInterface
	SseEngine() *sse_engine.StreamManager
	LegalEntityService() legal_entity_service.OrgLegalEntityServiceInterface
	CustomerEngineService() org_customer_service.OrgCustomerServiceInterface
	PaymentEngineService() payment_engine_service.PaymentServiceInterface
}
