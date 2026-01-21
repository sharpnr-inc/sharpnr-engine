package restaurantservices

import (
	restaurantrepo "sharpnr-engine/modules/hospitality/restaurant/repo"
)

type restaurantService struct {
	restaurantRepo restaurantrepo.RestaurantRepository
}

func NewRestaurantService(restaurantRepo restaurantrepo.RestaurantRepository) *restaurantService {
	return &restaurantService{
		restaurantRepo: restaurantRepo,
	}

}
