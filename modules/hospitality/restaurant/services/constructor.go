package restaurantservices

import (
	"context"
	restaurantrepo "sharpnr-engine/modules/hospitality/restaurant/repo"
)

type restaurantService struct {
	restaurantRepo restaurantrepo.RestaurantRepositoryInterface
}

// AddRestaurantGeneralDetails implements RestaurantServiceInterface.
func (r *restaurantService) AddRestaurantGeneralDetails(ctx context.Context) {
	panic("unimplemented")
}

func NewRestaurantService(restaurantRepo restaurantrepo.RestaurantRepositoryInterface) *restaurantService {
	return &restaurantService{
		restaurantRepo: restaurantRepo,
	}

}
