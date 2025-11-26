package kitchenservices

import (
	barrepo "sharpnr-engine/modules/hospitality/bar/repo"
	hotelrepo "sharpnr-engine/modules/hospitality/hotel/repo"
	kitchenrepo "sharpnr-engine/modules/hospitality/kitchen/kitchenRepo"
	"sharpnr-engine/modules/hospitality/restaurant/restaurantrepo"
)

type kitchenService struct {
	kitchenRepo    kitchenrepo.KitchenRepositoryInterface
	restaurantRepo restaurantrepo.RestaurantRepository
	barRepo        barrepo.BarRepository
	hotelRepo      hotelrepo.HotelRepository
}

func NewKitchenService(
	kitchenRepo kitchenrepo.KitchenRepositoryInterface,
	restaurantRepo restaurantrepo.RestaurantRepository,
	barRepo barrepo.BarRepository,
	hotelRepo hotelrepo.HotelRepository) *kitchenService {
	return &kitchenService{
		kitchenRepo:    kitchenRepo,
		restaurantRepo: restaurantRepo,
		barRepo:        barRepo,
		hotelRepo:      hotelRepo,
	}
}
