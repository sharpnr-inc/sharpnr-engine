package hotel_events

import (
	"sharpnr-engine/modules/hospitality/hotel/hotel_events_names"
	"sharpnr-engine/sere"
)

func RegisterHotelEvents(s *sere.SERE) {

	handleCharges := NewHospitality_ChargesCreatedEventHandler()
	s.SubscribeAsync(string(hotel_events_names.Hotel_ChargeIsCreatedInBooking), handleCharges.HandleCharges)
}
