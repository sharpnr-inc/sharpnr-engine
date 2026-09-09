package app

import (
	auditEngineService "sharpnr-engine/audit_log_engine/service"
	coreEmployeeServices "sharpnr-engine/core/employees/services"
	hrmsEngineRepo "sharpnr-engine/hrms_engine/repo"
	hrmsEngineService "sharpnr-engine/hrms_engine/service"
	dbtx "sharpnr-engine/pkg/dbTx"
	"sharpnr-engine/sero"
	sharp_mailer_engine "sharpnr-engine/sharp_mailer_engine"

	"gorm.io/gorm"
)

// bootstrapHRMSService builds the HRMS engine's service — the single surface
// behind the /sharp-hr routes. Attendance, payroll and recruitment flows join
// it as methods rather than as separate services.
func bootstrapHRMSService(
	db *gorm.DB,
	seroEngine *sero.SEROEngine,
	txManager dbtx.TxManager,
	mailerEngine *sharp_mailer_engine.SharpMailerEngine,
	employeesHRService coreEmployeeServices.EmployeeHRServiceInterface,
	auditLog auditEngineService.AuditLogServiceInterface,
) hrmsEngineService.HrmsEngineServiceInterface {

	hrmsRepo := hrmsEngineRepo.NewHrmsEngineRepo(db)

	return hrmsEngineService.NewHrmsEngineService(
		employeesHRService,
		seroEngine,
		txManager,
		mailerEngine,
		hrmsRepo,
		auditLog,
	)
}
