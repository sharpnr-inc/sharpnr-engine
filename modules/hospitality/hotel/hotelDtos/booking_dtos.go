package hoteldtos

import "github.com/google/uuid"

type BookingGuest_dto struct {
	Name   string
	DOB    string
	Gender string
}

type RoomBookingRequest_dto struct {
	RoomId         uuid.UUID
	PrimaryGuestID uuid.UUID
	CheckinDate    string
	CheckoutDate   string
	BookingGuests  []BookingGuest_dto
	BookingDate    string
}
