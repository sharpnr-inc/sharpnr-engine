package accounting_controllers

import accounting_service "sharpnr-engine/modules/accounting/accounting_service"

type accountingController struct {
	accSvc accounting_service.AccountingServiceInterface
}

func NewAccountingController(accSvc accounting_service.AccountingServiceInterface) *accountingController {
	return &accountingController{
		accSvc: accSvc,
	}
}
