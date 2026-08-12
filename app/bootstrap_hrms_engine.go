package app

import (
	coreEmployeeServices "sharpnr-engine/core/employees/services"
	hrmsEmployeeServices "sharpnr-engine/hrms_engine/employees/services"
	dbtx "sharpnr-engine/pkg/dbTx"
	"sharpnr-engine/sero"
	sharp_mailer_engine "sharpnr-engine/sharp_mailer_engine"
)

// bootstrapHRMSEmployeeService builds the HRMS employees module service.
// Attendance, payroll and recruitment get their own bootstrap function here as
// they land, each exposed on ModulesProvider in its own right.
func bootstrapHRMSEmployeeService(
	seroEngine *sero.SEROEngine,
	txManager dbtx.TxManager,
	mailerEngine *sharp_mailer_engine.SharpMailerEngine,
	employeesHRService coreEmployeeServices.EmployeeHRServiceInterface,
) hrmsEmployeeServices.HRMSEmployeeServiceInterface {

	return hrmsEmployeeServices.NewHRMSEmployeeService(
		employeesHRService,
		seroEngine,
		txManager,
		mailerEngine,
	)
}
