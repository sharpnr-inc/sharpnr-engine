package accounting_controllers

import accounting_service "sharpnr-engine/service_engines/accounting_engine/services"

type AccountingController struct {
	accSvc accounting_service.AccountingServiceInterface
}

func NewAccountingController(accSvc accounting_service.AccountingServiceInterface) *AccountingController {
	return &AccountingController{
		accSvc: accSvc,
	}
}
