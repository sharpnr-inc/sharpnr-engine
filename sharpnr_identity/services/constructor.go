package sharpnr_identity_auth_services

import (
	coreEmployeeServices "sharpnr-engine/core/employees/services"
	organization_services "sharpnr-engine/core/organization/services"

	sentinelservice "sharpnr-engine/sentinel/service"
)

type sharpnrIdentityAuthService struct {
	employeeService coreEmployeeServices.EmployeesIdentityServiceInterface
	orgService      organization_services.OrganizationServiceInterface
	sentinelService sentinelservice.SentinelServiceInterface
}

func NewSharpnrIdentityAuthService(
	employeeService coreEmployeeServices.EmployeesIdentityServiceInterface,
	orgService organization_services.OrganizationServiceInterface,
	sentinelService sentinelservice.SentinelServiceInterface,
) *sharpnrIdentityAuthService {
	return &sharpnrIdentityAuthService{
		employeeService: employeeService,
		orgService:      orgService,
		sentinelService: sentinelService,
	}
}
