package restaurantrepo

import (
	"context"
	restaurant_constants "sharpnr-engine/modules/hospitality/restaurant/constants"
	restaurant_models "sharpnr-engine/modules/hospitality/restaurant/models"
	"time"

	"github.com/google/uuid"
)

type RestaurantRepositoryInterface interface {

	//Restaurant
	CreateRestaurant(ctx context.Context, restaurant *restaurant_models.Restaurant) error
	GetRestaurantByID(ctx context.Context, restaurantID uuid.UUID) (*restaurant_models.Restaurant, error)
	GetRestaurantByOrgID(ctx context.Context, orgId uuid.UUID) ([]*restaurant_models.Restaurant, error)
	GetRestaurantByPropertyID(ctx context.Context, propertyID uuid.UUID) ([]*restaurant_models.Restaurant, error)
	UpdateGeneralDetails(ctx context.Context, restaurantID uuid.UUID, updates map[string]interface{}) error
	GetRestaurantPreferredKitchenID(ctx context.Context, restaurantID uuid.UUID) (*restaurant_models.RestaurantPreferedKitchen, error)
	CreatePreferredKitchen(
		ctx context.Context,
		pfktcn *restaurant_models.RestaurantPreferedKitchen,
	) error

	UpdatePreferredKitchen(ctx context.Context, pfk *restaurant_models.RestaurantPreferedKitchen) error

	// Restaurant--> Floors
	CreateFloor(ctx context.Context, floors *restaurant_models.RestaurantFloor) error
	GetFloorByID(ctx context.Context, floorID uuid.UUID) (*restaurant_models.RestaurantFloor, error)
	GetAllFloorOfRestaurant(ctx context.Context, restaurantID uuid.UUID) ([]*restaurant_models.RestaurantFloor, error)

	// Restaurant --> Floors --> tables
	CreateTables(ctx context.Context, resTables *restaurant_models.RestaurantTable) error
	UpdateTableStatus(ctx context.Context, tableID uuid.UUID, status restaurant_constants.RestaurantTableStatusType) error
	GetTableById(ctx context.Context, tableId uuid.UUID) (*restaurant_models.RestaurantTable, error)
	GetTablesByStatus(ctx context.Context, restaurantID uuid.UUID, status restaurant_constants.RestaurantTableStatusType) ([]*restaurant_models.RestaurantTable, error)
	GetAllTablesOfFloor(ctx context.Context, floorID uuid.UUID) ([]*restaurant_models.RestaurantTable, error)

	// Restaurant -- Floors --> Tables --> Dining Session

	// Dining Session payments

	// Orders
	CreateOrder(ctx context.Context, order *restaurant_models.RestaurantOrders) error
	CreateOrderItem(ctx context.Context, orderItem *restaurant_models.RestaurantOrderItems) error
	CreateOrderItemVariant(ctx context.Context, orderVariant *restaurant_models.RestaurantOrderItemsVariants) error
	CreateOrderItemAddon(ctx context.Context, orderAddon *restaurant_models.RestaurantOrderItemsAddOns) error

	GetOrderItemsByOrderID(ctx context.Context, restaurantOrderID uuid.UUID) ([]*restaurant_models.RestaurantOrderItems, error)
	GetOrderItemVariantByOrderID(ctx context.Context, restaurantOrderID uuid.UUID) ([]*restaurant_models.RestaurantOrderItemsVariants, error)
	GetOrderItemAddOnByOrderID(ctx context.Context, restaurantOrderID uuid.UUID) ([]*restaurant_models.RestaurantOrderItemsAddOns, error)

	GetOrdersByDiningSessionID(ctx context.Context, sessionID uuid.UUID) ([]*restaurant_models.RestaurantOrders, error)

	UpdateOrderStatus(ctx context.Context, orderID uuid.UUID, status string) error

	GetOrderSequenceNumber(ctx context.Context, restaurantID uuid.UUID, orgID uuid.UUID, nowTime time.Time) (string, error)

	RestaurantConfigurationRepoInterface
	RestaurantDiningFlowRepoInterface
	// RestaurantPickUpOrderFlowRepoInterface
}

type RestaurantConfigurationRepoInterface interface {
	CreateRestaurantConfiguration(ctx context.Context, restaurantConfig *restaurant_models.RestaurantConfiguration) error
	GetCurrentConfig(ctx context.Context, restaurantID uuid.UUID) (*restaurant_models.RestaurantConfiguration, error)
	UpdateRestuarantConfigStatus(ctx context.Context, restaurantID uuid.UUID, update restaurant_constants.RestaurantConfigurationStatusType) error

	// Methods for active and configuration

	UpdateConfiguration(ctx context.Context, restaurantID uuid.UUID, config bool) error
	UpdateActivation(ctx context.Context, restaurantID uuid.UUID, isActive bool) error
}

type RestaurantDiningFlowRepoInterface interface {
	CreateRestaurantDiningSession(ctx context.Context, diningSession *restaurant_models.RestaurantDiningSession) error
	UpdateDiningSessionStatus(ctx context.Context, diningSessionID uuid.UUID, status string) error
	GetDiningSessionByStatus(ctx context.Context, restaurantID uuid.UUID, status restaurant_constants.RestaurantDiningSessionStatusType) ([]*restaurant_models.RestaurantDiningSession, error)
	GetDiningSessionByTableID(ctx context.Context, tableID uuid.UUID) (*restaurant_models.RestaurantDiningSession, error)
}

type RestaurantPickUpOrderFlowRepoInterface interface {
	CreatePickUpOrder(ctx context.Context, pickUpOrder *restaurant_models.RestaurantPickUpOrders) error
	GetPickUpOrderByRestaurantID(ctx context.Context, restaurantID uuid.UUID) ([]restaurant_models.RestaurantPickUpOrders, error)
	GetPickUpOrderByID(ctx context.Context, pickUpOrderID uuid.UUID) (*restaurant_models.RestaurantPickUpOrders, error)

	CreatePickUpOrderCustomer(ctx context.Context, pickUpOrderCustomer *restaurant_models.RestaurantPickUpOrderCustomer) error
	GetPickUpOrderCustomerByPickUpOrderID(ctx context.Context, pickUpOrderID uuid.UUID) (*restaurant_models.RestaurantPickUpOrderCustomer, error)
	
}
