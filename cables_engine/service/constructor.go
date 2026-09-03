package cablesEngineService

import (
	auditEngineService "sharpnr-engine/audit_log_engine/service"
	cables_repo "sharpnr-engine/cables_engine/repo"
	coreEmployeeServices "sharpnr-engine/core/employees/services"
	"sharpnr-engine/sharp_mailer_engine"
)

type cablesService struct {
	cablesRepo         cables_repo.CablesRepoInterface
	mailerEngine       *sharp_mailer_engine.SharpMailerEngine
	auditLogEngine     auditEngineService.AuditLogServiceInterface
	employeeIamService coreEmployeeServices.EmployeeIamEngineServiceInterface
}

func NewCablesEngineService(
	cablesRepo cables_repo.CablesRepoInterface,
	mailerEngine *sharp_mailer_engine.SharpMailerEngine,
	auditLogEngine auditEngineService.AuditLogServiceInterface,
	employeeIamService coreEmployeeServices.EmployeeIamEngineServiceInterface,
) *cablesService {
	return &cablesService{
		cablesRepo:         cablesRepo,
		mailerEngine:       mailerEngine,
		auditLogEngine:     auditLogEngine,
		employeeIamService: employeeIamService,
	}
}
