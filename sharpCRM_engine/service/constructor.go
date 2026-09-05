package sharpCrmEngineService

import (
	org_customer_service "sharpnr-engine/core/customer/service"
	sharpCrmEngineRepo "sharpnr-engine/sharpCRM_engine/repo"
)

type Service struct {

	// this will give us access to the organization customer records
	coreOrgCustomerService org_customer_service.OrgCustomerServiceInterface
	// then we will soon add the all Business Ops to this service to know what service they avail and purchased
	// possibly we will add the accounting service to know how much amount they spent and how much they owe

	crmRepo sharpCrmEngineRepo.SharpCrmEngineRepoInterace
	// mailerEngine
	// txManager
	// cables service (not now)
}

func NewSharpCrmEngineService(
	coreOrgCustomerService org_customer_service.OrgCustomerServiceInterface,
	crmRepo sharpCrmEngineRepo.SharpCrmEngineRepoInterace,
) *Service {
	return &Service{
		coreOrgCustomerService: coreOrgCustomerService,
		crmRepo:                crmRepo,
	}
}
