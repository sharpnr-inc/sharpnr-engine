package app

import (
	orchestrator_service "sharpnr-engine/core/orchestrator/service"
	org_signup_services "sharpnr-engine/core/org-signup/services"
	sharpnr_identity_auth_services "sharpnr-engine/core/sharpnr-identity/services"
	"sharpnr-engine/modules/accounting/accounting_service"
	hr_employee_services "sharpnr-engine/modules/hr/employees/services"
)

type AppModules struct {
	orgSignupService    org_signup_services.OrgSignUpServiceInterface
	sharpAuthService    sharpnr_identity_auth_services.SharpnrIdentityAuthServiceInterface
	orchestratorService orchestrator_service.OrchestratorServiceInterface

	hospitality *HospitalityDomain

	hrService         hr_employee_services.HrEmployeeServiceInterface
	accountingService accounting_service.AccountingServiceInterface
}

func (m *AppModules) SharpnrOrchestrator() orchestrator_service.OrchestratorServiceInterface {
	return m.orchestratorService
}

func (m *AppModules) Hospitality() *HospitalityDomain {
	return m.hospitality
}

// AccountingService implements ModulesProvider.
func (m *AppModules) AccountingService() accounting_service.AccountingServiceInterface {
	return m.accountingService
}

// SharpAuthService implements ModulesProvider.
func (m *AppModules) SharpAuthService() sharpnr_identity_auth_services.SharpnrIdentityAuthServiceInterface {
	return m.sharpAuthService
}

func (m *AppModules) HrService() hr_employee_services.HrEmployeeServiceInterface {
	return m.hrService
}

func (m *AppModules) OrgSignUpService() org_signup_services.OrgSignUpServiceInterface {
	return m.orgSignupService
}
