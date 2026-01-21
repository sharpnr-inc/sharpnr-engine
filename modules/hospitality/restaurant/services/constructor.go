package restaurantservices

import (
	restaurantrepo "sharpnr-engine/modules/hospitality/restaurant/repo"
)

type restaurantService struct {
	restaurantRepo restaurantrepo.RestaurantRepositoryInterface
}

func NewRestaurantService(restaurantRepo restaurantrepo.RestaurantRepositoryInterface) *restaurantService {
	return &restaurantService{
		restaurantRepo: restaurantRepo,
	}

}

// Right now Restaurant Service only implements the restaurant repository
