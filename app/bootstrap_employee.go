package app

import (
	coreEmployeeRepo "sharpnr-engine/core/employees/repo"
	coreEmployeeServices "sharpnr-engine/core/employees/services"
	dbtx "sharpnr-engine/pkg/dbTx"
	"sharpnr-engine/sero"
	sharp_mailer_engine "sharpnr-engine/sharp_mailer_engine"

	"gorm.io/gorm"
)

// bootstrapEmployeeServices wires a single core employee service instance and
// exposes it to each consuming module through its own segregated interface,
// so every consumer only sees the employee Ops it is allowed to call.
func bootstrapEmployeeServices(
	db *gorm.DB,
	seroEngine *sero.SEROEngine,
	txManager dbtx.TxManager,
	mailerEngine *sharp_mailer_engine.SharpMailerEngine,
) (
	coreEmployeeServices.EmployeeAuthenticationServiceInterface,
	coreEmployeeServices.EmployeeHRServiceInterface,
	coreEmployeeServices.EmployeeIamEngineServiceInterface,
	coreEmployeeServices.EmployeeSignUpServiceInterface,
	coreEmployeeServices.EmployeeSetupWizServiceInterface,
) {
	employeeRepo := coreEmployeeRepo.NewEmployeeRepository(db)
	svc := coreEmployeeServices.NewEmployeeService(
		employeeRepo,
		seroEngine,
		txManager,
		mailerEngine,
	)
	return svc, svc, svc, svc, svc
}
