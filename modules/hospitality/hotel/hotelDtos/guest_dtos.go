package hoteldtos

import "github.com/google/uuid"

type CreateHotelGuestsRequest struct {
	Name      string
	DOB       string
	AadhaarId string
}

type HotelGuestResponseDTO struct {
	ID        uuid.UUID `json:"id"`
	Name      string    `json:"name"`
	DOB       string    `json:"dob"`
	Gender    string    `json:"gender"`
	AadhaarId string    `json:"aadhaarId"`
	Address   string    `json:"address"`
}
