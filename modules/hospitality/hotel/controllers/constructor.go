package hotelcontrollers

import (
	hotelservices "sharpnr-engine/modules/hospitality/hotel/services"
)

type HotelController struct {
	hotelService hotelservices.HotelServiceInterface
}

func NewHotelController(s hotelservices.HotelServiceInterface) *HotelController {
	return &HotelController{
		hotelService: s,
	}
}
