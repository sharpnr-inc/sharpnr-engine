package hotel_events

import (
	"sharpnr-engine/modules/accounting/accounting_repo"
	"sharpnr-engine/modules/hospitality/hotel/hotel_events_names"
	"sharpnr-engine/sere"
)

func RegisterHotelEvents(s *sere.SERE, accRepo accounting_repo.AccountingRepositoryInterface) {

	handleCharges := NewHospitality_ChargesCreatedEventHandler(accRepo)
	s.SubscribeAsync(string(hotel_events_names.Hotel_ChargeIsCreatedInBooking), handleCharges.HandleCharges)
}
