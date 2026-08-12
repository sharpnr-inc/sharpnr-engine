package sharpCRMEngineService

import org_customer_service "sharpnr-engine/core/customer/service"

type Service struct {

	// this will give us access to the organization customer records
	coreOrgCustomerService org_customer_service.OrgCustomerServiceInterface
	// then we will soon add the all Business Ops to this service to know what service they avail and purchased
	// possibly we will add the accounting service to know how much amount they spent and how much they owe
}
