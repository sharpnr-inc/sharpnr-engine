package hotelevents

import (
	"time"

	"github.com/google/uuid"
)

type HotelOrderCreated struct {
	OrganizationID uuid.UUID
	HotelID        uuid.UUID
	BookingID      uuid.UUID
	RoomID         uuid.UUID

	OrderID uuid.UUID // logical order ID (not KOT yet)

	Items []struct {
		MenuItemID uuid.UUID
		VariantID  *uuid.UUID
		Quantity   int
		Remarks    string
	}

	OrderedBy string // RECEPTION, GUEST
	PlacedAt  time.Time
}
