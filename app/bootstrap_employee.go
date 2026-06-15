package app

import (
	dbtx "sharpnr-engine/pkg/dbTx"
	"sharpnr-engine/sero"
	employee_repo "sharpnr-engine/sharp_hrms_engine/employees/repo"
	employee_services "sharpnr-engine/sharp_hrms_engine/employees/services"
	sharp_mailer_engine "sharpnr-engine/sharp_mailer_engine"

	"gorm.io/gorm"
)

func bootstrapEmployeeIdentity(
	db *gorm.DB,
	seroEngine *sero.SEROEngine,
	txManager dbtx.TxManager,
	mailerEngine *sharp_mailer_engine.SharpMailerEngine,
) employee_services.EmployeesIdentityServiceInterface {
	employeeRepo := employee_repo.NewEmployeeRepository(db)
	return employee_services.NewEmployeeService(
		employeeRepo,
		seroEngine,
		txManager,
		mailerEngine,
	)
}

func bootstrapEmployeeHR(
	db *gorm.DB,
	seroEngine *sero.SEROEngine,
	txManager dbtx.TxManager,
	mailerEngine *sharp_mailer_engine.SharpMailerEngine,
) employee_services.EmployeesHRServiceInterface {
	employeeRepo := employee_repo.NewEmployeeRepository(db)
	return employee_services.NewEmployeeService(
		employeeRepo,
		seroEngine,
		txManager,
		mailerEngine,
	)
}
