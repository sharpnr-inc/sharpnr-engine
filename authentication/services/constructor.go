package authenticationEngineServices

import (
	auditEngineService "sharpnr-engine/audit_log_engine/service"
	coreEmployeeServices "sharpnr-engine/core/employees/services"
	organization_services "sharpnr-engine/core/organization/services"

	sentinelservice "sharpnr-engine/sentinel/service"
)

type authenticationService struct {
	employeeService coreEmployeeServices.EmployeeAuthenticationServiceInterface
	orgService      organization_services.OrganizationServiceInterface
	sentinelService sentinelservice.SentinelServiceInterface
	auditEngine     auditEngineService.AuditLogServiceInterface
}

func NewAuthenticationService(
	employeeService coreEmployeeServices.EmployeeAuthenticationServiceInterface,
	orgService organization_services.OrganizationServiceInterface,
	sentinelService sentinelservice.SentinelServiceInterface,
	auditEngine auditEngineService.AuditLogServiceInterface,
) *authenticationService {
	return &authenticationService{
		employeeService: employeeService,
		orgService:      orgService,
		sentinelService: sentinelService,
		auditEngine:     auditEngine,
	}
}
