package acchospitality

import (
	"sharpnr-engine/modules/accounting/accounting_repo"
	hotel_events "sharpnr-engine/modules/accounting/accounting_sere_handlers/acc_hospitality/hotel"
	"sharpnr-engine/sere"
)

func RegisterHospitalityEvents(s *sere.SERE, accRepo accounting_repo.AccountingRepositoryInterface) {
	hotel_events.RegisterHotelEvents(s, accRepo)
}
