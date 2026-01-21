package accounting_sere_handlers

import (
	"sharpnr-engine/sere"
	acchospitality "sharpnr-engine/sharp_services/accounting_engine/sere_handlers/acc_hospitality"
)

func Register(
	sere *sere.SERE,
) {

	acchospitality.RegisterHospitalityEvents(sere)
}
