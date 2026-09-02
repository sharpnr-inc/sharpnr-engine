package orgDepartmentService

import (
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
) *coreOrgDepartmentService {
	return &coreOrgDepartmentService{
		orgDepartmentRepository: repo,
	}
}
