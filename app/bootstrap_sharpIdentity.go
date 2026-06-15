package app

import (
	organization_services "sharpnr-engine/core/organization/services"
	sharpnr_identity_auth_services "sharpnr-engine/core/sharpnr_identity/services"
	employee_services "sharpnr-engine/sharp_hrms_engine/employees/services"
)

func bootstrapIdentity(
	employeeService employee_services.EmployeesIdentityServiceInterface,
	orgService organization_services.OrganizationServiceInterface,
) sharpnr_identity_auth_services.SharpnrIdentityAuthServiceInterface {
	return sharpnr_identity_auth_services.NewSharpnrIdentityAuthService(employeeService, orgService)
}
