package app

import (
	orgDepartmentRepo "sharpnr-engine/core/departments/repos"
	orgDepartmentService "sharpnr-engine/core/departments/service"
	dbtx "sharpnr-engine/pkg/dbTx"
	"sharpnr-engine/sero"
	sharp_mailer_engine "sharpnr-engine/sharp_mailer_engine"

	"gorm.io/gorm"
)

func bootstrapOrgDepartments(
	db *gorm.DB,
	txManager dbtx.TxManager,
	mailerEngine *sharp_mailer_engine.SharpMailerEngine,
	seroEngine *sero.SEROEngine,
) orgDepartmentService.OrgDepartmentServiceInterface {
	// Initialize the organization department repository
	orgDepartmentRepository := orgDepartmentRepo.NewOrgDepartmentRepository(db)
	return orgDepartmentService.NewOrgDepartmentService(orgDepartmentRepository, txManager, mailerEngine, seroEngine)
}
