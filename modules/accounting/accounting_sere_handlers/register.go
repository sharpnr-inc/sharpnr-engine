package accounting_sere_handlers

import (
	core_events_names "sharpnr-engine/core/core_modules_events/coreEventsNames"
	"sharpnr-engine/modules/accounting/accounting_repo"
	acchospitality "sharpnr-engine/modules/accounting/accounting_sere_handlers/acc_hospitality"
	"sharpnr-engine/modules/accounting/accounting_sere_handlers/acc_organization"
	"sharpnr-engine/sere"
)

func Register(
	sere *sere.SERE,
	accRepo accounting_repo.AccountingRepositoryInterface,
) {

	handler := acc_organization.NewOrganizationCreatedEventHandler(accRepo)
	sere.SubscribeAsync(string(core_events_names.OrgCreated), handler.Handle)

	acchospitality.RegisterHospitalityEvents(sere, accRepo)
}
