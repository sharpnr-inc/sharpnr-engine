package hospitalitydtos

type FloorResponse struct {
	FloorID     string `json:"floor_id"`
	FloorNumber int    `json:"floor_number"`
	FloorName   string `json:"floor_name"`
}

type GetHotel_RESPONSE struct {
	HotelId   string          `json:"hotel_id"`
	HotelName string          `json:"hotel_name"`
	Floors    []FloorResponse `json:"floors"`
}
