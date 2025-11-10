package hotelcontrollers

import (
	hotelservices "sharpnr-engine/modules/hospitality/hotel/services"
)

/*
TODO:->
1 --> Need to add booking Service
2 --> Needs to wire up the Kitchen
3 --> Need to wire up the orders
4 --> payments
*/

type HotelController struct {
	hotelService hotelservices.HotelService
	// hotelBookingService hotelservices.HotelBookingServiceInterface
}

func NewHotelController(s hotelservices.HotelService) *HotelController {
	return &HotelController{
		hotelService: s,
		// hotelBookingService: hotelBookingService,
	}
}
