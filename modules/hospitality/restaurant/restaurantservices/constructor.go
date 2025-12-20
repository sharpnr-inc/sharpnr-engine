package restaurantservices

import (
	kitchenrepo "sharpnr-engine/modules/hospitality/kitchen/kitchenRepo"
	"sharpnr-engine/modules/hospitality/restaurant/restaurantrepo"
)

type restaurantService struct {
	restaurantRepo restaurantrepo.RestaurantRepository
	kitchenRepo    kitchenrepo.KitchenRepositoryInterface
}

func NewRestaurantRepository(restaurantRepo restaurantrepo.RestaurantRepository, kitchenRepo kitchenrepo.KitchenRepositoryInterface) *restaurantService {
	return &restaurantService{
		restaurantRepo: restaurantRepo,
		kitchenRepo:    kitchenRepo,
	}

}
