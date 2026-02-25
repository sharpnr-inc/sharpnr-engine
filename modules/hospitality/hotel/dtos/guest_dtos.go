package hoteldtos

import "github.com/google/uuid"

type CreateHotelGuestsRequest struct {
	Name      string `json:"name" gorm:"size:100;not null"`
	Mobile    string `json:"mobile" gorm:"size:15;not null;index"`
	Email     string `json:"email" gorm:"size:100"`
	DOB       string `json:"dob" gorm:"size:20"`
	AadhaarId string `json:"aadhaarId" gorm:"size:20;index"`
	Address   string `json:"address" gorm:"type:text"`
}

type HotelGuestResponseDTO struct {
	ID        uuid.UUID `json:"id"`
	Name      string    `json:"name"`
	Mobile    string    `json:"mobile"`
	DOB       string    `json:"dob"`
	Gender    string    `json:"gender"`
	AadhaarId string    `json:"aadhaarId"`
	Address   string    `json:"address"`
}
