package restaurantservices

import (
	"context"
	restaurantrepo "sharpnr-engine/modules/hospitality/restaurant/repo"
	"sharpnr-engine/sero"
)

type restaurantService struct {
	restaurantRepo restaurantrepo.RestaurantRepositoryInterface
	seroEngine     *sero.SEROEngine
}

// CreateOrder implements RestaurantServiceInterface.
func (r *restaurantService) CreateOrder(ctx context.Context) {
	panic("unimplemented")
}

// AddRestaurantGeneralDetails implements RestaurantServiceInterface.
func (r *restaurantService) AddRestaurantGeneralDetails(ctx context.Context) {
	panic("unimplemented")
}

func NewRestaurantService(
	restaurantRepo restaurantrepo.RestaurantRepositoryInterface,
	seroEngine *sero.SEROEngine,
) *restaurantService {
	return &restaurantService{
		restaurantRepo: restaurantRepo,
		seroEngine:     seroEngine,
	}

}
