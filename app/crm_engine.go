package app

import (
	auditEngineService "sharpnr-engine/audit_log_engine/service"
	org_customer_service "sharpnr-engine/core/customer/service"
	dbtx "sharpnr-engine/pkg/dbTx"
	sharpCrmEngineRepo "sharpnr-engine/sharpCRM_engine/repo"
	sharpCrmEngineService "sharpnr-engine/sharpCRM_engine/service"

	"gorm.io/gorm"
)

func bootstrapSharpCRMEngine(
	db *gorm.DB,
	txManager dbtx.TxManager,
	auditLogService auditEngineService.AuditLogServiceInterface,
	coreOrgCustomerService org_customer_service.OrgCustomerServiceInterface,
) sharpCrmEngineService.SharpCrmEngineServiceInterface {
	crmRepo := sharpCrmEngineRepo.NewSharpCrmEngineRepo(db)
	return sharpCrmEngineService.NewSharpCrmEngineService(coreOrgCustomerService, crmRepo, txManager, auditLogService)
}
