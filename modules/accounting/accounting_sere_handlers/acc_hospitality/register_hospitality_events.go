package acchospitality

import (
	hotel_events "sharpnr-engine/modules/accounting/accounting_sere_handlers/acc_hospitality/hotel"
	"sharpnr-engine/sere"
)

func RegisterHospitalityEvents(s *sere.SERE) {
	hotel_events.RegisterHotelEvents(s)
}
