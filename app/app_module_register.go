package app

import (
	hospitalitydomain "sharpnr-engine/app/hospitality_domain"
	auditEngineService "sharpnr-engine/audit_log_engine/service"
	org_customer_service "sharpnr-engine/core/customer/service"
	coreEmployeeServices "sharpnr-engine/core/employees/services"
	legal_entity_service "sharpnr-engine/core/legal_entity/service"
	org_property_services "sharpnr-engine/core/org_property/services"
	org_setup_wizard_services "sharpnr-engine/core/org_setup_wizard/services"

	hospitality_orchestration_service "sharpnr-engine/business_ops/hospitality/orchestration/service"
	org_signup_services "sharpnr-engine/core/org_signup/services"
	organization_services "sharpnr-engine/core/organization/services"
	iamEngineService "sharpnr-engine/iam_engine/service"
	sharpnr_identity_auth_services "sharpnr-engine/sharpnr_identity/services"
	hrmsengine "sharpnr-engine/hrms_engine/engine"

	orchestrator_service "sharpnr-engine/orchestrator/service"
	payment_engine_service "sharpnr-engine/payment_engine/services"
	accounting_service "sharpnr-engine/sharp_accounting_engine/accounting_engine/services"
	utilitiesApi_service "sharpnr-engine/utilities_apis/service"

	sse_engine "sharpnr-engine/sse_engine"
)

type AppModules struct {

	// auditLogService records every tenant event so the org has a full
	// history of what happened. Injected into every other engine.
	auditLogService auditEngineService.AuditLogServiceInterface

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

	employeesIdentityService coreEmployeeServices.EmployeesIdentityServiceInterface
	employeesHRService       coreEmployeeServices.EmployeesHRServiceInterface

	// iamEngineService exposes identity & access management operations
	// (profile, etc.) backed by the core employees module.
	iamEngineService iamEngineService.IamEngineServiceInterface

	// hrmsEngine aggregates all HRMS module services (employees, attendance,
	// payroll, recruitment, ...) behind a single entry point for the HR routes.
	hrmsEngine *hrmsengine.HRMSEngine

	accountingService    accounting_service.AccountingServiceInterface
	orgCustomerService   org_customer_service.OrgCustomerServiceInterface
	sseEngine            *sse_engine.StreamManager
	paymentEngineService payment_engine_service.PaymentServiceInterface
	utilitiesApiService  utilitiesApi_service.UtilitiesAPIServiceInterface

	// sentinelEngine is the Cedar authorization engine. It is injected into
	// services for s.sentinelEngine.IsAuthorized(...) calls and also handed to
	// the RequirePermission middleware at the route layer.
	// sentinelEngine *sentinelengine.Sentinel
}

// func (m *AppModules) Sentinel() *sentinelengine.Sentinel {
// 	return m.sentinelEngine
// }

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

func (m *AppModules) EmployeesIdentityService() coreEmployeeServices.EmployeesIdentityServiceInterface {
	return m.employeesIdentityService
}

func (m *AppModules) EmployeesHRService() coreEmployeeServices.EmployeesHRServiceInterface {
	return m.employeesHRService
}

func (m *AppModules) HRMS() *hrmsengine.HRMSEngine {
	return m.hrmsEngine
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

func (m *AppModules) UtilitiesApiService() utilitiesApi_service.UtilitiesAPIServiceInterface {
	return m.utilitiesApiService
}

func (m *AppModules) AuditLogService() auditEngineService.AuditLogServiceInterface {
	return m.auditLogService
}

func (m *AppModules) IamEngineService() iamEngineService.IamEngineServiceInterface {
	return m.iamEngineService
}
