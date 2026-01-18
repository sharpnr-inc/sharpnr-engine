package app

import (
	orchestrator_service "sharpnr-engine/core/orchestrator/service"
	org_legal_entity_service "sharpnr-engine/core/org_legal_entity/service"
	org_property_services "sharpnr-engine/core/org_property/services"
	org_setup_wizard_services "sharpnr-engine/core/org_setup_wizard/services"
	org_signup_services "sharpnr-engine/core/org_signup/services"
	organization_services "sharpnr-engine/core/organization/services"
	sharpnr_identity_auth_services "sharpnr-engine/core/sharpnr_identity/services"
	"sharpnr-engine/modules/accounting/accounting_service"
	employee_services "sharpnr-engine/modules/hr/employees/services"
)

type AppModules struct {

	// This is the Organization service which includes both repo
	// Redis and Main Pgsql DB
	organizationService organization_services.OrganizationServiceInterface

	orgLegalEntityService org_legal_entity_service.OrgLegalEntityServiceInterface
	//Signup service
	orgSignupService    org_signup_services.OrganizationSignUpServiceInterfaces
	sharpAuthService    sharpnr_identity_auth_services.SharpnrIdentityAuthServiceInterface
	orgPropertyService  org_property_services.OrganizationPropertyServicesInterfaces
	orchestratorService orchestrator_service.OrchestratorServiceInterface
	orgSetupWizard      org_setup_wizard_services.OrgSetupWizardServiceInterface

	hospitality *HospitalityDomain

	employeeService   employee_services.EmployeeServiceInterface
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

func (m *AppModules) EmployeeService() employee_services.EmployeeServiceInterface {
	return m.employeeService
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

func (m *AppModules) OrgSetupWiz() org_setup_wizard_services.OrgSetupWizardServiceInterface {
	return m.orgSetupWizard
}
