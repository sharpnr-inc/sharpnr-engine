package app

import (
	employee_services "sharpnr-engine/hrms_engine/employees/services"
	dbtx "sharpnr-engine/pkg/dbTx"
	sentinel_repo "sharpnr-engine/sentinel_engine/repo"
	sentinelservice "sharpnr-engine/sentinel_engine/service"

	"gorm.io/gorm"
)

// bootstrapSentinel wires the authorization (Sentinel) service: repo over the
// DB, service over the repo + transaction manager.
func bootstrapSentinel(
	db *gorm.DB,
	txManager dbtx.TxManager,
	employeeService employee_services.EmployeesIdentityServiceInterface,
) sentinelservice.SentinelService {
	repo := sentinel_repo.NewSentinelRepo(db)
	return sentinelservice.NewSentinelService(repo, txManager, employeeService)
}
