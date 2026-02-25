package accounting_controllers

import accounting_service "sharpnr-engine/service_engines/accounting_engine/services"

type accountingController struct {
	accSvc accounting_service.AccountingServiceInterface
}

func NewAccountingController(accSvc accounting_service.AccountingServiceInterface) *accountingController {
	return &accountingController{
		accSvc: accSvc,
	}
}
