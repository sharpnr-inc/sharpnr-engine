package hotelservices

import (
	hotelrepo "sharpnr-engine/modules/hospitality/hotel/repo"
)

type hotelService struct {
	hotelRepo hotelrepo.HotelRepository
}

func NewHotelService(hotelRepo hotelrepo.HotelRepository) *hotelService {
	return &hotelService{
		hotelRepo: hotelRepo,
	}
}
