package app

import (
	coreEmployeeRepo "sharpnr-engine/core/employees/repo"
	coreEmployeeServices "sharpnr-engine/core/employees/services"
	dbtx "sharpnr-engine/pkg/dbTx"
	"sharpnr-engine/sero"
	sharp_mailer_engine "sharpnr-engine/sharp_mailer_engine"

	"gorm.io/gorm"
)

func bootstrapEmployeeIdentity(
	db *gorm.DB,
	seroEngine *sero.SEROEngine,
	txManager dbtx.TxManager,
	mailerEngine *sharp_mailer_engine.SharpMailerEngine,
) coreEmployeeServices.EmployeeAuthenticationServiceInterface {
	employeeRepo := coreEmployeeRepo.NewEmployeeRepository(db)
	return coreEmployeeServices.NewEmployeeService(
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
) coreEmployeeServices.EmployeesHRServiceInterface {
	employeeRepo := coreEmployeeRepo.NewEmployeeRepository(db)
	return coreEmployeeServices.NewEmployeeService(
		employeeRepo,
		seroEngine,
		txManager,
		mailerEngine,
	)
}

func bootstrapEmployeeIamEngineService(
	db *gorm.DB,
	seroEngine *sero.SEROEngine,
	txManager dbtx.TxManager,
	mailerEngine *sharp_mailer_engine.SharpMailerEngine,
) coreEmployeeServices.EmployeeIamEngineServiceInterface {
	employeeRepo := coreEmployeeRepo.NewEmployeeRepository(db)
	return coreEmployeeServices.NewEmployeeService(
		employeeRepo,
		seroEngine,
		txManager,
		mailerEngine,
	)
}
