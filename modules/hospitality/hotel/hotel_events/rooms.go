package hotelevents

import "github.com/google/uuid"

type RoomBookingCreated struct {
	HotelRoomBookingId uuid.UUID
	HotelId            uuid.UUID
	RoomId             uuid.UUID
	PrimaryGuestId     uuid.UUID
}
