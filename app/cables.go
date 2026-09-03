package app

import (
	auditEngineService "sharpnr-engine/audit_log_engine/service"
	cablesEngineRepo "sharpnr-engine/cables_engine/repo"
	cablesEngineService "sharpnr-engine/cables_engine/service"
	coreEmployeeServices "sharpnr-engine/core/employees/services"
	dbtx "sharpnr-engine/pkg/dbTx"
	sharp_mailer_engine "sharpnr-engine/sharp_mailer_engine"

	"gorm.io/gorm"
)

func bootstrapCablesEngine(db *gorm.DB,
	txManager dbtx.TxManager,
	mailerEngine *sharp_mailer_engine.SharpMailerEngine,
	auditEngine auditEngineService.AuditLogServiceInterface,
	employeeIamService coreEmployeeServices.EmployeeIamEngineServiceInterface,
) cablesEngineService.CablesEngineServiceInterface {
	cablesEngineRepo := cablesEngineRepo.NewCablesRepo(db)
	return cablesEngineService.NewCablesEngineService(
		cablesEngineRepo,
		txManager,
		mailerEngine,
		auditEngine,
		employeeIamService,
	)

}
