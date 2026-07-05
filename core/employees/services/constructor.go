package coreEmployeeServices

import (
	coreEmployeeRepo "sharpnr-engine/core/employees/repo"
	dbtx "sharpnr-engine/pkg/dbTx"
	"sharpnr-engine/sero"

	"sharpnr-engine/sharp_mailer_engine"
)

type employeeService struct {
	employeeRepo coreEmployeeRepo.EmployeeRepositoryInterface
	seroEngine   *sero.SEROEngine
	txManager    dbtx.TxManager
	mailerEngine *sharp_mailer_engine.SharpMailerEngine
}

func NewEmployeeService(
	employeeRepo coreEmployeeRepo.EmployeeRepositoryInterface,
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
