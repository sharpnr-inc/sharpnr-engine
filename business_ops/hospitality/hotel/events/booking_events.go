package hotelevents

import "github.com/google/uuid"

type Hotel_NewRoomBookingCreateEvent struct {
	BookingId uuid.UUID
}

type Hotel_ChargesAddedInBooking struct {
}
