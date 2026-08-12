package hrmsEngineService

import (
	coreEmployeeServices "sharpnr-engine/core/employees/services"
	dbtx "sharpnr-engine/pkg/dbTx"
	"sharpnr-engine/sero"
	"sharpnr-engine/sharp_mailer_engine"
)

type hrmsService struct {
	employeeService coreEmployeeServices.EmployeeHRServiceInterface
	seroEngine      *sero.SEROEngine
	txManager       dbtx.TxManager
	mailerEngine    *sharp_mailer_engine.SharpMailerEngine
}

func NewHrmsEngineService(
	employeeService coreEmployeeServices.EmployeeHRServiceInterface,
	seroEngine *sero.SEROEngine,
	txManager dbtx.TxManager,
	mailerEngine *sharp_mailer_engine.SharpMailerEngine,
) HrmsEngineServiceInterface {

	return &hrmsService{
		employeeService: employeeService,
		seroEngine:      seroEngine,
		txManager:       txManager,
		mailerEngine:    mailerEngine,
	}
}
