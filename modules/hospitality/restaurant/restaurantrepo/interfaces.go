package restaurantrepo

import (
	"context"
	restaurant_models "sharpnr-engine/modules/hospitality/restaurant/models"

	"github.com/google/uuid"
)

type RestaurantRepository interface {
	CreateRestaurant(ctx context.Context, res *restaurant_models.Restaurant) error

	GetRestaurantByOrgId(ctx context.Context, orgId uuid.UUID) (*restaurant_models.Restaurant, error)

	UpdateRestaurant(ctx context.Context, restaurant *restaurant_models.Restaurant) error

	CreateOrAddFloorInRestaurant(ctx context.Context, floor *restaurant_models.RestaurantFloor) error

	CreateOrAddTablesInRestuarant(ctx context.Context,
		tables []*restaurant_models.RestaurantTable,
	) error

	GetRestaurantFloorById(
		ctx context.Context,
		floorId uuid.UUID,
	) (*restaurant_models.RestaurantFloor, error)

	/*********************
	 * Tables Operations *
	 *********************/

	// GetTableOfRestaurantByRestaurantId()
	GetTableById(ctx context.Context, tableId uuid.UUID) (*restaurant_models.RestaurantTable, error)
	// GetAllTablesInFloors()
	UpdateTable(ctx context.Context, tableId uuid.UUID, updates map[string]any) error
	// GetTableBookingSessionByID()

	/*********************
	 * Order Operations *
	*********************/

	CreateRes_TableSession(ctx context.Context, session *restaurant_models.RestaurantTableBookingSession) error

	GetBookingSessionByTableId(
		ctx context.Context,
		tableId uuid.UUID,
	) (*restaurant_models.RestaurantTableBookingSession, error)

	GetBookingSessionById(
		ctx context.Context,
		sessionId uuid.UUID,
	) (*restaurant_models.RestaurantTableBookingSession, error)

	UpdateRestaurantOrderItemKotReference(
		ctx context.Context,
		orderItemID uuid.UUID,
		ticketID uuid.UUID,
		ticketItemID uuid.UUID,
	) error

	// WithTransaction(ctx context.Context, fn func(tx RestaurantRepository) error) error
}
