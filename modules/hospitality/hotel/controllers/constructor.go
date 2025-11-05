package hotelcontrollers

import (
	hotelservices "sharpnr-engine/modules/hospitality/hotel/services"
)

type HotelController struct {
	hotelService hotelservices.HotelService
}

func NewHotelController(s hotelservices.HotelService) *HotelController {
	return &HotelController{
		hotelService: s,
	}
}
