package hotelevents

import (
	"sharpnr-engine/modules/hospitality/hospitalityconstants"
	"time"

	"github.com/google/uuid"
)

type ChargesCreatedInBooking struct {
	ChargesId uuid.UUID
	BookingID uuid.UUID

	HotelID uuid.UUID
	OrgID   uuid.UUID

	ChargesType hospitalityconstants.HotelChargeType
	Quantity    float64
	UnitAmount  float64

	TotalAmount float64

	ReferenceID *uuid.UUID
	OccurredAt  time.Time
}
