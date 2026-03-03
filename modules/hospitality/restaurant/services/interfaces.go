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
	GetGeneralDetails(
		ctx context.Context,
		orgID uuid.UUID,
		emplID uuid.UUID,
		restaurantID uuid.UUID,
	) (restaurant_contracts.RestaurantGeneralInfo_resp_dto, error)
	UpdateGeneralDetails(ctx context.Context, orgID, emplID, restaurantID uuid.UUID, req restaurant_contracts.UpdateBodyRequestDto) error
	AddFloors(ctx context.Context, AddFloorInput *restaurant_contracts.CreateFloorInput) (floorID uuid.UUID, err error)
	AddTablesInFloor(ctx context.Context, AddTablesInFloorInput *restaurant_contracts.CreateTablesInFloorInput) (tableID uuid.UUID, err error)

	AddRestaurantGeneralDetails(ctx context.Context)
	GetAllRestaurantOfProperty(ctx context.Context, PropertyID uuid.UUID) (restaurants []*restaurant_contracts.RestaurantDetails, err error)
	GetRestaurantPreferredKitchenID(
		ctx context.Context,
		restaurantID uuid.UUID,
	) (restaurant_contracts.RestaurantPreferredKitchen_resp_dto, error)

	UpdateRestaurantGeneralDetails(
		ctx context.Context,
		restaurantID uuid.UUID,
		input *restaurant_contracts.UpdateGeneralDetailsInput,
	) error
	SetupRestaurantLayout(ctx context.Context, orgID uuid.UUID, emplID uuid.UUID, restaurantID uuid.UUID, setupInput restaurant_contracts.AddLayoutRequestDto) error

	GetRestaurantLayout(
		ctx context.Context,
		orgID uuid.UUID, // PropertyID
		restaurantID uuid.UUID, // RestaurantID
	) (layout restaurant_contracts.RestaurantLayout_Response_dto, err error)

	/* ---------------------------------- Order --------------------------------- */

	CreateOrder(ctx context.Context)
	UpdateOrderStatus(ctx context.Context, orderID uuid.UUID, status string) error

	/* -------------------------------------------------------------------------- */
	/*                               DineIn Services                              */
	/* -------------------------------------------------------------------------- */

	StartDineSession(ctx context.Context,
		restaurantID uuid.UUID,
		orgID uuid.UUID,
		employeeID uuid.UUID,
		tableID uuid.UUID,
	) (uuid.UUID, error)

	GetDineSessionOrders(ctx context.Context, dineInSessionID uuid.UUID) (*restaurant_contracts.DiningSessionOrders_resp_dto, error)
	AddOrdersInSession(
		ctx context.Context,
		dineInSessionID uuid.UUID,
		req *restaurant_contracts.DiningSessionOrder_req_dto,
		orgID uuid.UUID,
		restaurantID uuid.UUID,
	) error

	GetRestaurantsOfOrg(
		ctx context.Context,
		orgID uuid.UUID,
	) ([]restaurant_contracts.RestaurantDetails, error)
}
