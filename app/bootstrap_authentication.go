package app

import (
	authenticationEngineServices "sharpnr-engine/authentication/services"
	coreEmployeeServices "sharpnr-engine/core/employees/services"
	organization_services "sharpnr-engine/core/organization/services"

	sentinelservice "sharpnr-engine/sentinel/service"
)

func bootstrapAuthentication(
	employeeService coreEmployeeServices.EmployeeAuthenticationServiceInterface,
	orgService organization_services.OrganizationServiceInterface,
	sentinelService sentinelservice.SentinelServiceInterface,
) authenticationEngineServices.AuthenticationServiceInterface {
	return authenticationEngineServices.NewAuthenticationService(
		employeeService,
		orgService,
		sentinelService,
	)
}
