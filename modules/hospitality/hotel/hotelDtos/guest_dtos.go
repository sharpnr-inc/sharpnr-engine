package hoteldtos

import "github.com/google/uuid"

type CreateHotelGuestsRequest struct {
	Name      string
	DOB       string
	AadhaarId string
}

type HotelGuest_dto struct {
	ID        uuid.UUID
	Name      string
	DOB       string
	Gender    string
	AadhaarId string
	Address   string
}
