package accounting_sere_handlers

import (
	"sharpnr-engine/modules/accounting/accounting_repo"
	"sharpnr-engine/sere"
)

func Register(
	sere *sere.SERE,
	accRepo accounting_repo.AccountingRepositoryInterface,
) {
	handler := NewOrganizationCreatedEventHandler(accRepo)
	sere.SubscribeAsync("OrgCreated", handler.Handle)
}
