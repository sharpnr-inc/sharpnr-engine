package hotelevents

import (
	"sharpnr-engine/modules/hospitality/hospitalityconstants"

	"github.com/google/uuid"
)

type ChargesCreatedInBooking struct {
	ChargesId   uuid.UUID
	BookingID   uuid.UUID
	ChargesType hospitalityconstants.HotelChargeType
}
