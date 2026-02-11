package hotel_contracts

import "github.com/google/uuid"

type CreateHotelFloorsInput struct {
	OrgID      uuid.UUID
	PropertyID uuid.UUID
	HotelID    uuid.UUID
	Floors     []FloorInput
}

type FloorInput struct {
	FloorName   string
	FloorNumber int
}

type AddHotelGeneralInfoInput struct {
	HotelID          string
	HotelName        string
	HotelOpeningTime string
}
