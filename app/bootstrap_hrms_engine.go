package app

import (
	coreEmployeeServices "sharpnr-engine/core/employees/services"
	hrmsEngineService "sharpnr-engine/hrms_engine/service"
	dbtx "sharpnr-engine/pkg/dbTx"
	"sharpnr-engine/sero"
	sharp_mailer_engine "sharpnr-engine/sharp_mailer_engine"
)

// bootstrapHRMSService builds the HRMS engine's service — the single surface
// behind the /sharp-hr routes. Attendance, payroll and recruitment flows join
// it as methods rather than as separate services.
func bootstrapHRMSService(
	seroEngine *sero.SEROEngine,
	txManager dbtx.TxManager,
	mailerEngine *sharp_mailer_engine.SharpMailerEngine,
	employeesHRService coreEmployeeServices.EmployeeHRServiceInterface,
) hrmsEngineService.HrmsEngineServiceInterface {

	return hrmsEngineService.NewHrmsEngineService(
		employeesHRService,
		seroEngine,
		txManager,
		mailerEngine,
	)
}
