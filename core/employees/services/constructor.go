package coreEmployeeServices

import (
	coreEmployeeRepo "sharpnr-engine/core/employees/repo"
	dbtx "sharpnr-engine/pkg/dbTx"
	sentinelservice "sharpnr-engine/sentinel/service"
	sentinelEngine "sharpnr-engine/sentinel_engine"
	"sharpnr-engine/sero"

	"sharpnr-engine/sharp_mailer_engine"
)

type employeeService struct {
	employeeRepo    coreEmployeeRepo.EmployeeRepositoryInterface
	seroEngine      *sero.SEROEngine
	txManager       dbtx.TxManager
	mailerEngine    *sharp_mailer_engine.SharpMailerEngine
	sentinelService sentinelservice.SentinelServiceInterface
	sentinelEngine  sentinelEngine.SentinelEngineInterface
}

func NewEmployeeService(
	employeeRepo coreEmployeeRepo.EmployeeRepositoryInterface,
	seroEngine *sero.SEROEngine,
	txManager dbtx.TxManager,
	mailerEngine *sharp_mailer_engine.SharpMailerEngine,
	sentinelService sentinelservice.SentinelServiceInterface,
	sentinelAuthzEngine sentinelEngine.SentinelEngineInterface,
) *employeeService {

	return &employeeService{
		employeeRepo:    employeeRepo,
		seroEngine:      seroEngine,
		txManager:       txManager,
		mailerEngine:    mailerEngine,
		sentinelService: sentinelService,
		sentinelEngine:  sentinelAuthzEngine,
	}
}
