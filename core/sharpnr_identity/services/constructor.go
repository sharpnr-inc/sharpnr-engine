package sharpnr_identity_auth_services

import (
	organization_services "sharpnr-engine/core/organization/services"
	employee_services "sharpnr-engine/hrms_engine/employees/services"
)

type sharpnrIdentityAuthService struct {
	employeeService employee_services.EmployeesIdentityServiceInterface
	orgService      organization_services.OrganizationServiceInterface
}

func NewSharpnrIdentityAuthService(
	employeeService employee_services.EmployeesIdentityServiceInterface,
	orgService organization_services.OrganizationServiceInterface,
) *sharpnrIdentityAuthService {
	return &sharpnrIdentityAuthService{
		employeeService: employeeService,
		orgService:      orgService,
	}
}
