package hrmsEmployeeServices

import (
	coreEmployeeServices "sharpnr-engine/core/employees/services"
	dbtx "sharpnr-engine/pkg/dbTx"
	"sharpnr-engine/sero"

	"sharpnr-engine/sharp_mailer_engine"
)

type hrmsemployeeService struct {
	employeeService coreEmployeeServices.EmployeeHRServiceInterface
	seroEngine      *sero.SEROEngine
	txManager       dbtx.TxManager
	mailerEngine    *sharp_mailer_engine.SharpMailerEngine
}

func NewHRMSEmployeeService(
	employeeService coreEmployeeServices.EmployeeHRServiceInterface,
	seroEngine *sero.SEROEngine,
	txManager dbtx.TxManager,
	mailerEngine *sharp_mailer_engine.SharpMailerEngine,
) HRMSEmployeeServiceInterface {

	return &hrmsemployeeService{
		employeeService: employeeService,
		seroEngine:      seroEngine,
		txManager:       txManager,
		mailerEngine:    mailerEngine,
	}
}
