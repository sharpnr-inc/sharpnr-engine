package hotelservices

import (
	hotelrepo "sharpnr-engine/modules/hospitality/hotel/repo"
	kitchenrepo "sharpnr-engine/modules/hospitality/kitchen/kitchenRepo"
)

type hotelService struct {
	hotelRepo   hotelrepo.HotelRepositoryInterface
	houseKeRepo hotelrepo.HotelHouseKeepingRepositoryInterface
	kitchenRepo kitchenrepo.KitchenRepositoryInterface
}

func NewHotelService(
	hotelRepo hotelrepo.HotelRepositoryInterface,
	houseKeRepo hotelrepo.HotelHouseKeepingRepositoryInterface,
	kitchenRepo kitchenrepo.KitchenRepositoryInterface) *hotelService {
	return &hotelService{
		hotelRepo:   hotelRepo,
		houseKeRepo: houseKeRepo,
		kitchenRepo: kitchenRepo,
	}
}
