package kitchenservices

import (
	barrepo "sharpnr-engine/modules/hospitality/bar/repo"
	hotelrepo "sharpnr-engine/modules/hospitality/hotel/repo"
	kitchenrepo "sharpnr-engine/modules/hospitality/kitchen/kitchenRepo"
	menurepo "sharpnr-engine/modules/hospitality/menu/menu_repo"
	"sharpnr-engine/modules/hospitality/restaurant/restaurantrepo"
)

type kitchenService struct {
	kitchenRepo    kitchenrepo.KitchenRepositoryInterface
	restaurantRepo restaurantrepo.RestaurantRepository
	barRepo        barrepo.BarRepository
	hotelRepo      hotelrepo.HotelRepositoryInterface
	menuRepo       menurepo.MenuRepositoryInterface
}

func NewKitchenService(
	kitchenRepo kitchenrepo.KitchenRepositoryInterface,
	restaurantRepo restaurantrepo.RestaurantRepository,
	barRepo barrepo.BarRepository,
	hotelRepo hotelrepo.HotelRepositoryInterface,
	menuRepo menurepo.MenuRepositoryInterface) *kitchenService {
	return &kitchenService{
		kitchenRepo:    kitchenRepo,
		restaurantRepo: restaurantRepo,
		barRepo:        barRepo,
		hotelRepo:      hotelRepo,
		menuRepo:       menuRepo,
	}
}
