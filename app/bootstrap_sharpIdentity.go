package app

import (
	coreEmployeeServices "sharpnr-engine/core/employees/services"
	organization_services "sharpnr-engine/core/organization/services"
	sharpnr_identity_auth_services "sharpnr-engine/core/sharpnr_identity/services"

	sentinelservice "sharpnr-engine/sentinel/service"
)

func bootstrapIdentity(
	employeeService coreEmployeeServices.EmployeesIdentityServiceInterface,
	orgService organization_services.OrganizationServiceInterface,
	sentinelService sentinelservice.SentinelServiceInterface,
) sharpnr_identity_auth_services.SharpnrIdentityAuthServiceInterface {
	return sharpnr_identity_auth_services.NewSharpnrIdentityAuthService(
		employeeService,
		orgService,
		sentinelService,
	)
}
