package app

import (
	coreEmployeeServices "sharpnr-engine/core/employees/services"
	hrmsEmployeeServices "sharpnr-engine/hrms_engine/employees/services"
	hrmsengine "sharpnr-engine/hrms_engine/engine"
	dbtx "sharpnr-engine/pkg/dbTx"
	"sharpnr-engine/sero"
	sharp_mailer_engine "sharpnr-engine/sharp_mailer_engine"

	"gorm.io/gorm"
)

// bootstrapHRMSEngine wires up every HRMS module (employees today; attendance,
// payroll, recruitment later — each with its own repo/service built here) and
// returns the single HRMSEngine aggregate consumed by the HR routes.
func bootstrapHRMSEngine(
	db *gorm.DB,
	seroEngine *sero.SEROEngine,
	txManager dbtx.TxManager,
	mailerEngine *sharp_mailer_engine.SharpMailerEngine,
	employeesHRService coreEmployeeServices.EmployeeHRServiceInterface,
) *hrmsengine.HRMSEngine {

	hrmsEmployeeService := hrmsEmployeeServices.NewHRMSEmployeeService(
		employeesHRService,
		seroEngine,
		txManager,
		mailerEngine,
	)

	return hrmsengine.NewHRMSEngine(
		hrmsEmployeeService,
	)
}
