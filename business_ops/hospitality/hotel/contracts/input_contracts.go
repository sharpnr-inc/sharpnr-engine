package hotel_contracts

import (
	hotel_constants "sharpnr-engine/business_ops/hospitality/hotel/constants"

	"github.com/google/uuid"
)

type FloorInput struct {
	HotelID       uuid.UUID
	FloorName     string
	FloorNumber   int
	NumberOfRooms int
}

type RoomInput struct {
	HotelID   uuid.UUID
	FloorID   uuid.UUID
	Name      string
	Code      string
	Status    hotel_constants.HotelRoomStatusType
	BedType   string
	MaxGuests int
}

type AddHotelGeneralInfoInput struct {
	HotelID          uuid.UUID
	HotelName        string
	HotelOpeningTime string
	HotelClosingTime string
	Is24X7           bool
}

type HotelMainGenDetailsInput struct {
	HotelID      uuid.UUID
	Name         string
	IsConfigured bool
}
