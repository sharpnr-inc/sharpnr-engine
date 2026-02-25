package restaurantrepo

import (
	"context"
	restaurant_constants "sharpnr-engine/modules/hospitality/restaurant/constants"
	restaurant_models "sharpnr-engine/modules/hospitality/restaurant/models"

	"github.com/google/uuid"
)

type RestaurantRepositoryInterface interface {

	//Restaurant
	CreateRestaurant(ctx context.Context, restaurant *restaurant_models.Restaurant) error
	GetRestaurantByID(ctx context.Context, restaurantID uuid.UUID) (*restaurant_models.Restaurant, error)
	GetRestaurantByOrgID(ctx context.Context, orgId uuid.UUID) ([]*restaurant_models.Restaurant, error)
	GetRestaurantByPropertyID(ctx context.Context, propertyID uuid.UUID) ([]*restaurant_models.Restaurant, error)
	UpdateGeneralDetails(ctx context.Context, restaurantID uuid.UUID, updates map[string]interface{}) error

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
	CreateRestaurantDiningSession(ctx context.Context, diningSession *restaurant_models.RestaurantDiningSession) error
	UpdateDiningSessionStatus(ctx context.Context, diningSessionID uuid.UUID, status string) error
	GetDiningSessionByStatus(ctx context.Context, restaurantID uuid.UUID, status restaurant_constants.RestaurantDiningSessionStatusType) ([]*restaurant_models.RestaurantDiningSession, error)
	GetDiningSessionByTableID(ctx context.Context, tableID uuid.UUID) (*restaurant_models.RestaurantDiningSession, error)

	// Dining Session payments
	CreateSessionPayment(ctx context.Context, sessionPayment *restaurant_models.RestaurantSessionPayment) error
	GetAllPaymentsOfDiningSession(ctx context.Context, sessionID uuid.UUID) ([]*restaurant_models.RestaurantSessionPayment, error)

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

	// Charges
	CreateCharges(ctx context.Context, charges *restaurant_models.RestaurantCharge) error
	GetAllChargesOfSession(ctx context.Context, sessionID uuid.UUID) ([]*restaurant_models.RestaurantCharge, error)
}
