package hospitalitydtos

type CreateFloorInHotel_REQ struct {
	HotelId   string `json:"hotel_id"`
	FloorName string `json:"floor_name"`
}

type CreateFloorInHotel_RESPONSE struct {
	HotelId   string `json:"hotel_id"`
	FloorName string `json:"floor_name"`
	FloorID   string `json:"floor_id"`
}
