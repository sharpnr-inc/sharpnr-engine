package hotelcontrollers

import (
	hotelservices "sharpnr-engine/modules/hospitality/hotel/services"
)

type hotelController struct {
	hotelService hotelservices.HotelServiceInterface
}

func NewHotelController(s hotelservices.HotelServiceInterface) *hotelController {
	return &hotelController{
		hotelService: s,
	}
}
