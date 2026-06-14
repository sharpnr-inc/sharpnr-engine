package app

import (
	hospitalitydomain "sharpnr-engine/app/hospitality_domain"
	org_customer_service "sharpnr-engine/core/customer/service"
	legal_entity_service "sharpnr-engine/core/legal_entity/service"
	org_property_services "sharpnr-engine/core/org_property/services"
	org_setup_wizard_services "sharpnr-engine/core/org_setup_wizard/services"

	hospitality_orchestration_service "sharpnr-engine/business_ops/hospitality/orchestration/service"
	org_signup_services "sharpnr-engine/core/org_signup/services"
	organization_services "sharpnr-engine/core/organization/services"
	sharpnr_identity_auth_services "sharpnr-engine/core/sharpnr_identity/services"
	orchestrator_service "sharpnr-engine/orchestrator/service"
	payment_engine_service "sharpnr-engine/payment_engine/services"
	accounting_service "sharpnr-engine/sharp_accounting_engine/accounting_engine/services"
	employee_services "sharpnr-engine/sharp_hrms_engine/employees/services"

	sse_engine "sharpnr-engine/sse_engine"
)

type AppModules struct {

	// This is the Organization service which includes both repo
	// Redis and Main Pgsql DB
	organizationService organization_services.OrganizationServiceInterface

	orgLegalEntityService legal_entity_service.OrgLegalEntityServiceInterface
	//Signup service
	orgSignupService    org_signup_services.OrganizationSignUpServiceInterfaces
	sharpAuthService    sharpnr_identity_auth_services.SharpnrIdentityAuthServiceInterface
	orgPropertyService  org_property_services.OrganizationPropertyServicesInterfaces
	orchestratorService orchestrator_service.OrchestratorServiceInterface
	orgSetupWizard      org_setup_wizard_services.OrgSetupWizardV1_ServiceInterface

	hospitality              *hospitalitydomain.HospitalityDomain
	hospitalityOrchestration hospitality_orchestration_service.OrchestrationServiceInterface

	employeesIdentityService employee_services.EmployeesIdentityServiceInterface
	employeesHRService       employee_services.EmployeesHRServiceInterface

	accountingService    accounting_service.AccountingServiceInterface
	orgCustomerService   org_customer_service.OrgCustomerServiceInterface
	sseEngine            *sse_engine.StreamManager
	paymentEngineService payment_engine_service.PaymentServiceInterface
}

func (m *AppModules) SharpnrOrchestrator() orchestrator_service.OrchestratorServiceInterface {
	return m.orchestratorService
}

func (m *AppModules) Hospitality() *hospitalitydomain.HospitalityDomain {
	return m.hospitality
}

func (m *AppModules) HospitalityOrchestration() hospitality_orchestration_service.OrchestrationServiceInterface {
	return m.hospitalityOrchestration
}

// AccountingService implements ModulesProvider.
func (m *AppModules) AccountingService() accounting_service.AccountingServiceInterface {
	return m.accountingService
}

// SharpAuthService implements ModulesProvider.
func (m *AppModules) SharpAuthService() sharpnr_identity_auth_services.SharpnrIdentityAuthServiceInterface {
	return m.sharpAuthService
}

func (m *AppModules) EmployeesIdentityService() employee_services.EmployeesIdentityServiceInterface {
	return m.employeesIdentityService
}

func (m *AppModules) EmployeesHRService() employee_services.EmployeesHRServiceInterface {
	return m.employeesHRService
}

func (m *AppModules) OrgSignUpService() org_signup_services.OrganizationSignUpServiceInterfaces {
	return m.orgSignupService
}

func (m *AppModules) OrgPropertyService() org_property_services.OrganizationPropertyServicesInterfaces {
	return m.orgPropertyService
}

func (m *AppModules) OrganizationService() organization_services.OrganizationServiceInterface {
	return m.organizationService
}

func (m *AppModules) OrgSetupWiz() org_setup_wizard_services.OrgSetupWizardV1_ServiceInterface {
	return m.orgSetupWizard
}

func (m *AppModules) SseEngine() *sse_engine.StreamManager {
	return m.sseEngine
}

func (m *AppModules) LegalEntityService() legal_entity_service.OrgLegalEntityServiceInterface {
	return m.orgLegalEntityService
}

func (m *AppModules) CustomerEngineService() org_customer_service.OrgCustomerServiceInterface {
	return m.orgCustomerService
}

func (m *AppModules) PaymentEngineService() payment_engine_service.PaymentServiceInterface {
	return m.paymentEngineService
}
