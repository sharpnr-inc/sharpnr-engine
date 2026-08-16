package hotelcontrollers

import (
	hotelservices "sharpnr-engine/property_business_ops/hospitality/hotel/services"
)

type hotelController struct {
	hotelService hotelservices.HotelServiceInterface
}

func NewHotelController(s hotelservices.HotelServiceInterface) *hotelController {
	return &hotelController{
		hotelService: s,
	}
}
