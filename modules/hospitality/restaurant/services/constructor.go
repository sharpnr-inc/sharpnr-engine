package restaurantservices

import (
	"context"
	restaurant_contracts "sharpnr-engine/modules/hospitality/restaurant/contracts"
	restaurantrepo "sharpnr-engine/modules/hospitality/restaurant/repo"

	"github.com/google/uuid"
)

type restaurantService struct {
	restaurantRepo restaurantrepo.RestaurantRepositoryInterface
}

// AddFloors implements RestaurantServiceInterface.
func (r *restaurantService) AddFloors(ctx context.Context, AddFloorInput *restaurant_contracts.CreateFloorInput) (floorID uuid.UUID, err error) {
	panic("unimplemented")
}

// AddRestaurantGeneralDetails implements RestaurantServiceInterface.
func (r *restaurantService) AddRestaurantGeneralDetails(ctx context.Context) {
	panic("unimplemented")
}

// AddTablesInFloor implements RestaurantServiceInterface.
func (r *restaurantService) AddTablesInFloor(ctx context.Context, AddTablesInFloorInput *restaurant_contracts.CreateTablesInFloor) (tableID uuid.UUID, err error) {
	panic("unimplemented")
}

func NewRestaurantService(restaurantRepo restaurantrepo.RestaurantRepositoryInterface) *restaurantService {
	return &restaurantService{
		restaurantRepo: restaurantRepo,
	}

}
