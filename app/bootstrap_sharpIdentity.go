package app

import (
	organization_services "sharpnr-engine/core/organization/services"
	sharpnr_identity_auth_services "sharpnr-engine/core/sharpnr_identity/services"
	employee_services "sharpnr-engine/hrms_engine/employees/services"
	sentinelservice "sharpnr-engine/sentinel/service"
)

func bootstrapIdentity(
	employeeService employee_services.EmployeesIdentityServiceInterface,
	orgService organization_services.OrganizationServiceInterface,
	sentinelService sentinelservice.SentinelServiceInterface,
) sharpnr_identity_auth_services.SharpnrIdentityAuthServiceInterface {
	return sharpnr_identity_auth_services.NewSharpnrIdentityAuthService(
		employeeService,
		orgService,
		sentinelService,
	)
}
