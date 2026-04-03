package restaurantservices

import (
	"context"
	restaurant_constants "sharpnr-engine/modules/hospitality/restaurant/constants"
	restaurant_contracts "sharpnr-engine/modules/hospitality/restaurant/contracts"
	inventory_engine_models "sharpnr-engine/service_engines/inventory_engine/models"

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

	AddRestaurantGeneralDetails(ctx context.Context)
	GetAllRestaurantOfProperty(ctx context.Context, PropertyID uuid.UUID) (restaurants []*restaurant_contracts.RestaurantDetails, err error)
	GetRestaurantPreferredKitchenID(
		ctx context.Context,
		orgID, employeeID,
		restaurantID uuid.UUID,
	) (restaurant_contracts.RestaurantPreferredKitchen_resp_dto, error)

	AddPreferredKitchen(
		ctx context.Context,
		orgID, employeeID,
		restaurantID, kitchenID uuid.UUID,
	) error

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

	GetRestaurantsOfOrg(
		ctx context.Context,
		orgID uuid.UUID,
	) ([]restaurant_contracts.RestaurantDetails, error)

	RestaurantConfigsServiceInterface
	RestaurantTableAndFloorConfigsServices
	RestaurantDiningSessionService
	RestaurantInventoryServiceInterface
}

type RestaurantConfigsServiceInterface interface {
	AddConfigurationStatus(ctx context.Context, orgID uuid.UUID, restaurantID uuid.UUID, status restaurant_constants.RestaurantConfigurationStatusType) error
	UpdateConfigurationStatus(ctx context.Context, orgID uuid.UUID, restaurantID uuid.UUID, status restaurant_constants.RestaurantConfigurationStatusType) error
	RestaurantSwitch(ctx context.Context, orgID, emplID, restaurantID uuid.UUID, restaurantSwitch bool) error
	// GetConfigurationStatus(ctx context.Context, orgID uuid.UUID, restaurantID uuid.UUID) error
}

type RestaurantTableAndFloorConfigsServices interface {
	AddFloors(ctx context.Context, AddFloorInput *restaurant_contracts.CreateFloorInput) (floorID uuid.UUID, err error)
	AddTablesInFloor(ctx context.Context, AddTablesInFloorInput *restaurant_contracts.CreateTablesInFloorInput) (tableID uuid.UUID, err error)
	SetupRestaurantLayout(ctx context.Context, orgID uuid.UUID, emplID uuid.UUID, restaurantID uuid.UUID, setupInput restaurant_contracts.AddLayoutRequestDto) error
}

type RestaurantDiningSessionService interface {
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

	CloseDineInSession(
		ctx context.Context,
		orgID, emplID, restaurantID, dineInSessionID uuid.UUID,
		request restaurant_contracts.CloseSessionRequest,
	) (*restaurant_contracts.CloseSessionResponse, error)
}

type RestaurantPickUpOrderServiceInterface interface {
}

type RestaurantInventoryServiceInterface interface {
	GetAllInventoryAccounts(ctx context.Context, orgID, emplID uuid.UUID) ([]inventory_engine_models.InventoryAccount, error)
	AddNewInventoryInRestaurant(ctx context.Context, orgID, emplID, restaurantID uuid.UUID, request *restaurant_contracts.AddInventoryRequestDto) error
	GetRestaurantAllInventories(ctx context.Context, orgID, emplID, restaurantID uuid.UUID, inventoryAccountID *uuid.UUID) ([]restaurant_contracts.RestuarantInventoriesResponseDto, error)
	// GetRestaurantInventoriesByInventoryAccount(ctx context.Context, orgID, emplID, restaurantID, inventoryAccountID uuid.UUID) ([]restaurant_contracts.RestuarantInventoriesResponseDto, error)
}
