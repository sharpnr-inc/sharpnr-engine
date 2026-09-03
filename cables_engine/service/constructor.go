package cablesEngineService

import (
	auditEngineService "sharpnr-engine/audit_log_engine/service"
	cables_repo "sharpnr-engine/cables_engine/repo"
	coreEmployeeServices "sharpnr-engine/core/employees/services"
	dbtx "sharpnr-engine/pkg/dbTx"
	"sharpnr-engine/sharp_mailer_engine"
)

type cablesService struct {
	cablesRepo cables_repo.CablesRepoInterface

	// txManager wraps a write and the audit entry that records it in one
	// transaction, so a rolled-back change never leaves a log line claiming it
	// happened.
	txManager      dbtx.TxManager
	auditLogEngine auditEngineService.AuditLogServiceInterface

	// Parked: wired here but not consumed by any cables flow yet. Nothing in
	// the account or channel flows mails anyone or needs to resolve an employee
	// profile — an invite does, and that is the flow that will reach for these.
	mailerEngine       *sharp_mailer_engine.SharpMailerEngine
	employeeIamService coreEmployeeServices.EmployeeIamEngineServiceInterface
}

func NewCablesEngineService(
	cablesRepo cables_repo.CablesRepoInterface,
	txManager dbtx.TxManager,
	mailerEngine *sharp_mailer_engine.SharpMailerEngine,
	auditLogEngine auditEngineService.AuditLogServiceInterface,
	employeeIamService coreEmployeeServices.EmployeeIamEngineServiceInterface,
) *cablesService {
	return &cablesService{
		cablesRepo:         cablesRepo,
		txManager:          txManager,
		mailerEngine:       mailerEngine,
		auditLogEngine:     auditLogEngine,
		employeeIamService: employeeIamService,
	}
}
