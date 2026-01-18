package accounting_sere_handlers

import (
	acchospitality "sharpnr-engine/modules/accounting/accounting_sere_handlers/acc_hospitality"
	"sharpnr-engine/sere"
)

func Register(
	sere *sere.SERE,
) {

	acchospitality.RegisterHospitalityEvents(sere)
}
