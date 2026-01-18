package app

import (
	orchestrator_service "sharpnr-engine/core/orchestrator/service"
	org_property_services "sharpnr-engine/core/org_property/services"
	org_setup_wizard_services "sharpnr-engine/core/org_setup_wizard/services"
	org_signup_services "sharpnr-engine/core/org_signup/services"
	organization_services "sharpnr-engine/core/organization/services"
	sharpnr_identity_auth_services "sharpnr-engine/core/sharpnr_identity/services"
	"sharpnr-engine/modules/accounting/accounting_service"
	employee_services "sharpnr-engine/modules/hr/employees/services"
)

type ModulesProvider interface {
	OrganizationService() organization_services.OrganizationServiceInterface
	SharpnrOrchestrator() orchestrator_service.OrchestratorServiceInterface
	OrgSetupWiz() org_setup_wizard_services.OrgSetupWizardServiceInterface
	Hospitality() *HospitalityDomain
	OrgPropertyService() org_property_services.OrganizationPropertyServicesInterfaces

	OrgSignUpService() org_signup_services.OrganizationSignUpServiceInterfaces
	SharpAuthService() sharpnr_identity_auth_services.SharpnrIdentityAuthServiceInterface

	EmployeeService() employee_services.EmployeeServiceInterface
	AccountingService() accounting_service.AccountingServiceInterface
}
