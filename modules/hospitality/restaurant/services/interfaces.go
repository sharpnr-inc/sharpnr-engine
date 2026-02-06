package restaurantservices

import (
	"context"
	restaurant_contracts "sharpnr-engine/modules/hospitality/restaurant/contracts"

	"github.com/google/uuid"
)

/****************************************
 * Service Interface for the Restuarant *
 *                                      *
****************************************/

type RestaurantServiceInterface interface {
	SetupDefaultRestaurant(ctx context.Context, orgID uuid.UUID, propertyID uuid.UUID) (restuarantID uuid.UUID, err error)

	GetAllRestaurantOfProperty(ctx context.Context, PropertyID uuid.UUID) (restaurants []*restaurant_contracts.RestaurantDetails, err error)

	AddFloors(ctx context.Context, AddFloorInput *restaurant_contracts.CreateFloorInput) (floorID uuid.UUID, err error)
	AddTablesInFloor(ctx context.Context, AddTablesInFloorInput *restaurant_contracts.CreateTablesInFloor) (tableID uuid.UUID, err error)

	AddRestaurantGeneralDetails(ctx context.Context)
}
