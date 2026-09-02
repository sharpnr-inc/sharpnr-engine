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
}

func NewOrgDepartmentService(
	repo orgDepartmentRepo.OrgDepartmentRepositoryInterface,
	txManager dbtx.TxManager,
	mailerEngine *sharp_mailer_engine.SharpMailerEngine,
	seroEngine *sero.SEROEngine,
	auditEngine *auditEngineService.AuditLogServiceInterface,
) *coreOrgDepartmentService {
	return &coreOrgDepartmentService{
		orgDepartmentRepository: repo,
	}
}
