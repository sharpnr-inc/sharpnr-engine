package hospitalitydtos

import "github.com/google/uuid"

type CreateFloorInHotel_REQ struct {
	HotelId   string `json:"hotel_id"`
	FloorName string `json:"floor_name"`
}

type CreateFloorInHotel_RESPONSE struct {
	HotelId   string    `json:"hotel_id"`
	FloorName string    `json:"floor_name"`
	FloorID   uuid.UUID `json:"floor_id"`
}
