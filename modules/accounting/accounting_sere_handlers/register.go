package accounting_sere_handlers

import (
	"sharpnr-engine/modules/accounting/accounting_repo"
	acchospitality "sharpnr-engine/modules/accounting/accounting_sere_handlers/acc_hospitality"
	"sharpnr-engine/sere"
)

func Register(
	sere *sere.SERE,
	accRepo accounting_repo.AccountingRepositoryInterface,
) {

	acchospitality.RegisterHospitalityEvents(sere, accRepo)
}
