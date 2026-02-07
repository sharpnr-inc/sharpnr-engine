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
	AddFloors(ctx context.Context, AddFloorInput *restaurant_contracts.CreateFloorInput) (floorID uuid.UUID, err error)
	AddTablesInFloor(ctx context.Context, AddTablesInFloorInput *restaurant_contracts.CreateTablesInFloorInput) (tableID uuid.UUID, err error)

	AddRestaurantGeneralDetails(ctx context.Context)
	GetAllRestaurantOfProperty(ctx context.Context, PropertyID uuid.UUID) (restaurants []*restaurant_contracts.RestaurantDetails, err error)

	UpdateRestaurantGeneralDetails(
		ctx context.Context,
		restaurantID uuid.UUID,
		input *restaurant_contracts.UpdateGeneralDetailsInput,
	) error
	SetupRestaurantLayout(ctx context.Context, propertyID uuid.UUID, restaurantID uuid.UUID, setupInput []restaurant_contracts.SetupRestaurantLayoutInput) error
}
