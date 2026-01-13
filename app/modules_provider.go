package app

import (
	org_signup_services "sharpnr-engine/core/org-signup/services"
	sharpnr_identity_auth_services "sharpnr-engine/core/sharpnr-identity/services"
	"sharpnr-engine/modules/accounting/accounting_service"
	hr_employee_services "sharpnr-engine/modules/hr/employees/services"
)

type ModulesProvider interface {
	Hospitality() *HospitalityDomain

	OrgSignUpService() org_signup_services.OrgSignUpServiceInterface
	SharpAuthService() sharpnr_identity_auth_services.SharpnrIdentityAuthServiceInterface

	HrService() hr_employee_services.HrEmployeeServiceInterface
	AccountingService() accounting_service.AccountingServiceInterface
}
