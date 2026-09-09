package hrmsEngineService

import (
	auditEngineService "sharpnr-engine/audit_log_engine/service"
	coreEmployeeServices "sharpnr-engine/core/employees/services"
	hrmsEngineRepo "sharpnr-engine/hrms_engine/repo"
	dbtx "sharpnr-engine/pkg/dbTx"
	"sharpnr-engine/sero"
	"sharpnr-engine/sharp_mailer_engine"
)

type hrmsService struct {
	employeeService coreEmployeeServices.EmployeeHRServiceInterface
	seroEngine      *sero.SEROEngine
	txManager       dbtx.TxManager
	mailerEngine    *sharp_mailer_engine.SharpMailerEngine
	hrmsRepo        hrmsEngineRepo.HrmsEngineRepoInterface
	auditLogEngine  auditEngineService.AuditLogServiceInterface
}

func NewHrmsEngineService(
	employeeService coreEmployeeServices.EmployeeHRServiceInterface,
	seroEngine *sero.SEROEngine,
	txManager dbtx.TxManager,
	mailerEngine *sharp_mailer_engine.SharpMailerEngine,
	hrmsRepo hrmsEngineRepo.HrmsEngineRepoInterface,
	auditLogEngine auditEngineService.AuditLogServiceInterface,
) HrmsEngineServiceInterface {

	return &hrmsService{
		employeeService: employeeService,
		seroEngine:      seroEngine,
		txManager:       txManager,
		mailerEngine:    mailerEngine,
		hrmsRepo:        hrmsRepo,
		auditLogEngine:  auditLogEngine,
	}
}
