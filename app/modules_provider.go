package app

import (
	orchestrator_service "sharpnr-engine/core/orchestrator/service"
	org_property_services "sharpnr-engine/core/org_property/services"
	org_signup_services "sharpnr-engine/core/org_signup/services"
	sharpnr_identity_auth_services "sharpnr-engine/core/sharpnr_identity/services"
	"sharpnr-engine/modules/accounting/accounting_service"
	hr_employee_services "sharpnr-engine/modules/hr/employees/services"
)

type ModulesProvider interface {
	SharpnrOrchestrator() orchestrator_service.OrchestratorServiceInterface
	Hospitality() *HospitalityDomain
	OrgPropertyService() org_property_services.OrganizationPropertyServicesInterfaces

	OrgSignUpService() org_signup_services.OrgSignUpServiceInterface
	SharpAuthService() sharpnr_identity_auth_services.SharpnrIdentityAuthServiceInterface

	HrService() hr_employee_services.HrEmployeeServiceInterface
	AccountingService() accounting_service.AccountingServiceInterface
}
