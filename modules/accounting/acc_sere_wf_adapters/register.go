package acc_sere_wf_adapters

import (
	"sharpnr-engine/modules/accounting/accounting_repo"
	"sharpnr-engine/sere"
)

func RegisterAccountingBarWorkflows(
	s *sere.SERE,
	accountingRepo accounting_repo.AccountingRepositoryInterface,
) {
	s.SubscribeWorkflow("ORG_ONBOARD_WORKFLOW", barAccountingDefaultSetup(accountingRepo))
}

func RegisterAccountingHotelWorkflow(
	s *sere.SERE,
	accountingRepo accounting_repo.AccountingRepositoryInterface,
) {
	s.SubscribeWorkflow("ORG_ONBOARD_WORKFLOW", hotelAccountingDefaultSetup(accountingRepo))
}

func RegisterAccountingRestaurantWorkflow(
	s *sere.SERE,
	accountingRepo accounting_repo.AccountingRepositoryInterface,
) {
	s.SubscribeWorkflow("ORG_ONBOARD_WORKFLOW", restaurantAccountingDefaultSetup(accountingRepo))
}
