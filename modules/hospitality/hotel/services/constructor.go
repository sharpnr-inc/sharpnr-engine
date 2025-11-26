package hotelservices

import (
	hotelrepo "sharpnr-engine/modules/hospitality/hotel/repo"
	kitchenrepo "sharpnr-engine/modules/hospitality/kitchen/kitchenRepo"
)

type hotelService struct {
	hotelRepo   hotelrepo.HotelRepository
	kitchenRepo kitchenrepo.KitchenRepositoryInterface
}

func NewHotelService(
	hotelRepo hotelrepo.HotelRepository,
	kitchenRepo kitchenrepo.KitchenRepositoryInterface) *hotelService {
	return &hotelService{
		hotelRepo:   hotelRepo,
		kitchenRepo: kitchenRepo,
	}
}
