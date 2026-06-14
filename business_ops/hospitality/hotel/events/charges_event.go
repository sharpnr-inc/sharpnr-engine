package hotelevents

import (
	"time"

	"github.com/google/uuid"
)

type ChargesCreatedInBooking struct {
	ChargesId uuid.UUID
	BookingID uuid.UUID

	HotelID uuid.UUID
	OrgID   uuid.UUID

	ChargesType string
	Quantity    float64
	UnitAmount  float64

	TotalAmount float64

	ReferenceID *uuid.UUID
	OccurredAt  time.Time
}
