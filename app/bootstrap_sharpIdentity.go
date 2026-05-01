package app

import (
	sharpnr_identity_auth_services "sharpnr-engine/core/sharpnr_identity/services"
	employee_services "sharpnr-engine/modules/hr/employees/services"
)

func bootstrapIdentity(
	employeeService employee_services.EmployeesIdentityServiceInterface,
) sharpnr_identity_auth_services.SharpnrIdentityAuthServiceInterface {

	return sharpnr_identity_auth_services.NewSharpnrIdentityAuthService(employeeService)
}
