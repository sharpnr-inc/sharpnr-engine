package hotelservices

import (
	hotelrepo "sharpnr-engine/modules/hospitality/hotel/repo"
)

type hotelService struct {
	hotelRepo hotelrepo.HotelRepository
	// menuRepo  menurepo.MenuRepositoryInterface
}

func NewHotelService(hotelRepo hotelrepo.HotelRepository) *hotelService {
	return &hotelService{
		hotelRepo: hotelRepo,
		// menuRepo: m,
	}
}
