package app

import (
	auditEngineService "sharpnr-engine/audit_log_engine/service"
	cablesEngineRepo "sharpnr-engine/cables_engine/repo"
	cablesEngineService "sharpnr-engine/cables_engine/service"
	coreEmployeeServices "sharpnr-engine/core/employees/services"
	dbtx "sharpnr-engine/pkg/dbTx"
	sharp_mailer_engine "sharpnr-engine/sharp_mailer_engine"
	ws_engine "sharpnr-engine/ws_engine"

	"gorm.io/gorm"
)

func bootstrapCablesEngine(db *gorm.DB,
	txManager dbtx.TxManager,
	mailerEngine *sharp_mailer_engine.SharpMailerEngine,
	auditEngine auditEngineService.AuditLogServiceInterface,
	employeeIamService coreEmployeeServices.EmployeeIamEngineServiceInterface,
	wsEngine *ws_engine.Engine,
) cablesEngineService.CablesEngineServiceInterface {
	cablesEngineRepo := cablesEngineRepo.NewCablesRepo(db)

	service := cablesEngineService.NewCablesEngineService(
		cablesEngineRepo,
		txManager,
		mailerEngine,
		auditEngine,
		employeeIamService,
		wsEngine,
	)

	// Claim the socket namespaces cables owns. The engine denies any namespace
	// nobody registered, so without this every "conv:" and "member:" subscribe
	// is refused — the topics genuinely do not exist until this runs.
	cablesEngineService.RegisterRealtimeTopics(wsEngine, service)

	return service
}
