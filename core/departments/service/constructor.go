package orgDepartmentService

import (
	auditEngineService "sharpnr-engine/audit_log_engine/service"
	orgDepartmentRepo "sharpnr-engine/core/departments/repos"
	dbtx "sharpnr-engine/pkg/dbTx"
	"sharpnr-engine/sero"
	"sharpnr-engine/sharp_mailer_engine"
)

type coreOrgDepartmentService struct {
	// Repository interface for department repository operations
	orgDepartmentRepository orgDepartmentRepo.OrgDepartmentRepositoryInterface

	// txManager wraps a write and the audit entry that records it in one
	// transaction, so a rolled-back change never leaves a log line claiming it
	// happened.
	txManager   dbtx.TxManager
	auditEngine auditEngineService.AuditLogServiceInterface

	// Parked: wired here but not consumed by any department flow yet. A
	// department change notifies nobody and publishes no event today; both
	// land on this struct the day one does.
	mailerEngine *sharp_mailer_engine.SharpMailerEngine
	seroEngine   *sero.SEROEngine
}

func NewOrgDepartmentService(
	repo orgDepartmentRepo.OrgDepartmentRepositoryInterface,
	txManager dbtx.TxManager,
	mailerEngine *sharp_mailer_engine.SharpMailerEngine,
	seroEngine *sero.SEROEngine,
	auditEngine auditEngineService.AuditLogServiceInterface,
) *coreOrgDepartmentService {
	return &coreOrgDepartmentService{
		orgDepartmentRepository: repo,
		txManager:               txManager,
		auditEngine:             auditEngine,
		mailerEngine:            mailerEngine,
		seroEngine:              seroEngine,
	}
}
