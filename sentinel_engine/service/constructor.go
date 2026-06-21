package sentinelservice

import (
	employee_services "sharpnr-engine/hrms_engine/employees/services"
	dbtx "sharpnr-engine/pkg/dbTx"
	sentinel_repo "sharpnr-engine/sentinel_engine/repo"
)

type sentinelService struct {
	repo            sentinel_repo.SentinelRepositoryInterface
	txManager       dbtx.TxManager
	employeeService employee_services.EmployeesIdentityServiceInterface
}

func NewSentinelService(
	repo sentinel_repo.SentinelRepositoryInterface,
	txManager dbtx.TxManager,
	employeeService employee_services.EmployeesIdentityServiceInterface,
) *sentinelService {
	return &sentinelService{
		repo:            repo,
		txManager:       txManager,
		employeeService: employeeService,
	}
}
