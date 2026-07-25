package authenticationEngineServices

import (
	coreEmployeeServices "sharpnr-engine/core/employees/services"
	organization_services "sharpnr-engine/core/organization/services"

	sentinelservice "sharpnr-engine/sentinel/service"
)

type authenticationService struct {
	employeeService coreEmployeeServices.EmployeeAuthenticationServiceInterface
	orgService      organization_services.OrganizationServiceInterface
	sentinelService sentinelservice.SentinelServiceInterface
}

func NewAuthenticationService(
	employeeService coreEmployeeServices.EmployeeAuthenticationServiceInterface,
	orgService organization_services.OrganizationServiceInterface,
	sentinelService sentinelservice.SentinelServiceInterface,
) *authenticationService {
	return &authenticationService{
		employeeService: employeeService,
		orgService:      orgService,
		sentinelService: sentinelService,
	}
}
