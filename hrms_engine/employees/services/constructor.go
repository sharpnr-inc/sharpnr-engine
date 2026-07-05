package hrmsEmployeeServices

import (
	employee_services "sharpnr-engine/core/employees/services"
	dbtx "sharpnr-engine/pkg/dbTx"
	"sharpnr-engine/sero"

	"sharpnr-engine/sharp_mailer_engine"
)

type hrmsemployeeService struct {
	employeeService employee_services.EmployeesHRServiceInterface
	seroEngine      *sero.SEROEngine
	txManager       dbtx.TxManager
	mailerEngine    *sharp_mailer_engine.SharpMailerEngine
}

func NewHRMSEmployeeService(
	employeeService employee_services.EmployeesHRServiceInterface,
	seroEngine *sero.SEROEngine,
	txManager dbtx.TxManager,
	mailerEngine *sharp_mailer_engine.SharpMailerEngine,
) *hrmsemployeeService {

	return &hrmsemployeeService{
		employeeService: employeeService,
		seroEngine:      seroEngine,
		txManager:       txManager,
		mailerEngine:    mailerEngine,
	}
}
