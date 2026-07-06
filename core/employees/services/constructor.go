package coreEmployeeServices

import (
	coreEmployeeRepo "sharpnr-engine/core/employees/repo"
	dbtx "sharpnr-engine/pkg/dbTx"
	sentinelservice "sharpnr-engine/sentinel/service"
	"sharpnr-engine/sero"

	"sharpnr-engine/sharp_mailer_engine"
)

type employeeService struct {
	employeeRepo    coreEmployeeRepo.EmployeeRepositoryInterface
	seroEngine      *sero.SEROEngine
	txManager       dbtx.TxManager
	mailerEngine    *sharp_mailer_engine.SharpMailerEngine
	sentinelService sentinelservice.SentinelServiceInterface
}

func NewEmployeeService(
	employeeRepo coreEmployeeRepo.EmployeeRepositoryInterface,
	seroEngine *sero.SEROEngine,
	txManager dbtx.TxManager,
	mailerEngine *sharp_mailer_engine.SharpMailerEngine,
	sentinelService sentinelservice.SentinelServiceInterface,
) *employeeService {

	return &employeeService{
		employeeRepo:    employeeRepo,
		seroEngine:      seroEngine,
		txManager:       txManager,
		mailerEngine:    mailerEngine,
		sentinelService: sentinelService,
	}
}
