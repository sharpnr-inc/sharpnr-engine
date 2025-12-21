package accounting_sere_handlers

import (
	core_events_names "sharpnr-engine/core/core_modules_events/coreEventsNames"
	"sharpnr-engine/modules/accounting/accounting_repo"
	"sharpnr-engine/sere"
)

func Register(
	sere *sere.SERE,
	accRepo accounting_repo.AccountingRepositoryInterface,
) {
	handler := NewOrganizationCreatedEventHandler(accRepo)
	sere.SubscribeAsync(string(core_events_names.OrgCreated), handler.Handle)
}
