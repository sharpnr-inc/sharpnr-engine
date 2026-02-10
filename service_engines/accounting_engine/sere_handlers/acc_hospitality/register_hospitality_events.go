package acchospitality

import (
	"sharpnr-engine/sere"
	hotel_events "sharpnr-engine/service_engines/accounting_engine/sere_handlers/acc_hospitality/hotel"
)

func RegisterHospitalityEvents(s *sere.SERE) {
	hotel_events.RegisterHotelEvents(s)
}
