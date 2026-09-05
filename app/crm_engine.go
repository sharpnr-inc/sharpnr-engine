package app

import (
	org_customer_service "sharpnr-engine/core/customer/service"
	sharpCrmEngineRepo "sharpnr-engine/sharpCRM_engine/repo"
	sharpCrmEngineService "sharpnr-engine/sharpCRM_engine/service"

	"gorm.io/gorm"
)

func bootstrapSharpCRMEngine(
	db *gorm.DB,
	coreOrgCustomerService org_customer_service.OrgCustomerServiceInterface,
) sharpCrmEngineService.SharpCrmEngineServiceInterface {
	crmRepo := sharpCrmEngineRepo.NewSharpCrmEngineRepo(db)
	return sharpCrmEngineService.NewSharpCrmEngineService(coreOrgCustomerService, crmRepo)
}
