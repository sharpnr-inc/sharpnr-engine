package employee_services

import (
	dbtx "sharpnr-engine/pkg/dbTx"
	"sharpnr-engine/sero"
	employee_repo "sharpnr-engine/hrms_engine/employees/repo"
	"sharpnr-engine/sharp_mailer_engine"
)

type employeeService struct {
	employeeRepo employee_repo.EmployeeRepositoryInterface
	seroEngine   *sero.SEROEngine
	txManager    dbtx.TxManager
	mailerEngine *sharp_mailer_engine.SharpMailerEngine
}

func NewEmployeeService(
	employeeRepo employee_repo.EmployeeRepositoryInterface,
	seroEngine *sero.SEROEngine,
	txManager dbtx.TxManager,
	mailerEngine *sharp_mailer_engine.SharpMailerEngine,
) *employeeService {

	return &employeeService{
		employeeRepo: employeeRepo,
		seroEngine:   seroEngine,
		txManager:    txManager,
		mailerEngine: mailerEngine,
	}
}
