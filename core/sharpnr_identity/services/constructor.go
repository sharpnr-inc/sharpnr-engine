package sharpnr_identity_auth_services

import (
	organization_services "sharpnr-engine/core/organization/services"
	employee_services "sharpnr-engine/hrms_engine/employees/services"
	sentinelservice "sharpnr-engine/sentinel/service"
)

type sharpnrIdentityAuthService struct {
	employeeService employee_services.EmployeesIdentityServiceInterface
	orgService      organization_services.OrganizationServiceInterface
	sentinelService sentinelservice.SentinelServiceInterface
}

func NewSharpnrIdentityAuthService(
	employeeService employee_services.EmployeesIdentityServiceInterface,
	orgService organization_services.OrganizationServiceInterface,
	sentinelService sentinelservice.SentinelServiceInterface,
) *sharpnrIdentityAuthService {
	return &sharpnrIdentityAuthService{
		employeeService: employeeService,
		orgService:      orgService,
		sentinelService: sentinelService,
	}
}
