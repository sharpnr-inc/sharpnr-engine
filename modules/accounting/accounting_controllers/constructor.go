package accounting_controllers

import accountingservice "sharpnr-engine/modules/accounting/accounting_service"

type accountingController struct {
	accSvc accountingservice.AccountingServiceInterface
}

func NewAccountingController(accSvc accountingservice.AccountingServiceInterface) *accountingController {
	return &accountingController{
		accSvc: accSvc,
	}
}
