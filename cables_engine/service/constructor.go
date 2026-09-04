package cablesEngineService

import (
	auditEngineService "sharpnr-engine/audit_log_engine/service"
	cables_repo "sharpnr-engine/cables_engine/repo"
	coreEmployeeServices "sharpnr-engine/core/employees/services"
	dbtx "sharpnr-engine/pkg/dbTx"
	"sharpnr-engine/sharp_mailer_engine"
	ws_engine "sharpnr-engine/ws_engine"
)

type cablesService struct {
	cablesRepo cables_repo.CablesRepoInterface

	// txManager wraps a write and the audit entry that records it in one
	// transaction, so a rolled-back change never leaves a log line claiming it
	// happened.
	txManager      dbtx.TxManager
	auditLogEngine auditEngineService.AuditLogServiceInterface

	// wsEngine carries the fan-out. Nil is a working configuration — every
	// publish becomes a no-op and the HTTP API still behaves — so tests and any
	// build without the socket wired do not have to special-case it.
	wsEngine *ws_engine.Engine

	// employeeIamService is read exactly once per bootstrap, to reconcile the
	// org roster into cables_members. Never on a hot path: after the sync,
	// rendering an author is a single-table read of cables' own data.
	employeeIamService coreEmployeeServices.EmployeeIamEngineServiceInterface

	// Parked: no cables flow mails anyone yet. An invite will.
	mailerEngine *sharp_mailer_engine.SharpMailerEngine
}

func NewCablesEngineService(
	cablesRepo cables_repo.CablesRepoInterface,
	txManager dbtx.TxManager,
	mailerEngine *sharp_mailer_engine.SharpMailerEngine,
	auditLogEngine auditEngineService.AuditLogServiceInterface,
	employeeIamService coreEmployeeServices.EmployeeIamEngineServiceInterface,
	wsEngine *ws_engine.Engine,
) *cablesService {
	return &cablesService{
		cablesRepo:         cablesRepo,
		txManager:          txManager,
		mailerEngine:       mailerEngine,
		auditLogEngine:     auditLogEngine,
		employeeIamService: employeeIamService,
		wsEngine:           wsEngine,
	}
}
