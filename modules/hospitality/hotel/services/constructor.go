package hotelservices

import (
	hotelrepo "sharpnr-engine/modules/hospitality/hotel/repo"
	menurepo "sharpnr-engine/modules/hospitality/menu/menu_repo"
)

type hotelService struct {
	hotelRepo hotelrepo.HotelRepository
	menuRepo  menurepo.MenuRepositoryInterface
}

func NewHotelService(hotelRepo hotelrepo.HotelRepository) *hotelService {
	return &hotelService{
		hotelRepo: hotelRepo,
	}
}
