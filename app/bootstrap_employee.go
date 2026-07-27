package app

import (
	auditEngineService "sharpnr-engine/audit_log_engine/service"
	coreEmployeeRepo "sharpnr-engine/core/employees/repo"
	coreEmployeeServices "sharpnr-engine/core/employees/services"
	dbtx "sharpnr-engine/pkg/dbTx"
	jwt_manager "sharpnr-engine/pkg/jwtmanager"
	sentinelservice "sharpnr-engine/sentinel/service"
	sentinelEngine "sharpnr-engine/sentinel_engine"
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
	sentinelService sentinelservice.SentinelServiceInterface,
	sentinelAuthzEngine sentinelEngine.SentinelEngineInterface,
	auditEngine auditEngineService.AuditLogServiceInterface,
	jwtManager jwt_manager.Manager,
) (
	coreEmployeeServices.EmployeeAuthenticationServiceInterface,
	coreEmployeeServices.EmployeeHRServiceInterface,
	coreEmployeeServices.EmployeeIamEngineServiceInterface,
	coreEmployeeServices.EmployeeSignUpServiceInterface,
	coreEmployeeServices.EmployeeSetupWizServiceInterface,
	coreEmployeeServices.EmployeeOrchestratorServiceInterface,
) {
	employeeRepo := coreEmployeeRepo.NewEmployeeRepository(db)
	svc := coreEmployeeServices.NewEmployeeService(
		employeeRepo,
		seroEngine,
		txManager,
		mailerEngine,
		sentinelService,
		sentinelAuthzEngine,
		auditEngine,
		jwtManager,
	)
	return svc, svc, svc, svc, svc, svc
}
