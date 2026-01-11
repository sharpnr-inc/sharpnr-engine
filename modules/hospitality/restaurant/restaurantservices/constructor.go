package restaurantservices

import (
	"sharpnr-engine/modules/hospitality/restaurant/restaurantrepo"
)

type restaurantService struct {
	restaurantRepo restaurantrepo.RestaurantRepository
}

func NewRestaurantService(restaurantRepo restaurantrepo.RestaurantRepository) *restaurantService {
	return &restaurantService{
		restaurantRepo: restaurantRepo,
	}

}
