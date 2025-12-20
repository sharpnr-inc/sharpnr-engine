package restaurantrepo

import (
	"context"
	hospitalitymodels "sharpnr-engine/modules/hospitality/hospitality_models"

	"github.com/google/uuid"
)

type RestaurantRepository interface {
	CreateRestaurant(ctx context.Context, res *hospitalitymodels.Restaurant) error

	GetRestaurantByOrgId(ctx context.Context, orgId uuid.UUID) (*hospitalitymodels.Restaurant, error)

	UpdateRestaurant(ctx context.Context, restaurant *hospitalitymodels.Restaurant) error

	CreateOrAddFloorInRestaurant(ctx context.Context, floor *hospitalitymodels.RestaurantFloor) error

	CreateOrAddTablesInRestuarant(ctx context.Context,
		tables []*hospitalitymodels.RestaurantTable,
	) error

	GetRestaurantFloorById(
		ctx context.Context,
		floorId uuid.UUID,
	) (*hospitalitymodels.RestaurantFloor, error)

	/*********************
	 * Tables Operations *
	 *********************/

	// GetTableOfRestaurantByRestaurantId()
	GetTableById(ctx context.Context, tableId uuid.UUID) (*hospitalitymodels.RestaurantTable, error)
	// GetAllTablesInFloors()
	UpdateTable(ctx context.Context, tableId uuid.UUID, updates map[string]any) error
	// GetTableBookingSessionByID()

	/*********************
	 * Order Operations *
	*********************/

	CreateRes_TableSession(ctx context.Context, session *hospitalitymodels.RestaurantTableBookingSession) error

	GetBookingSessionByTableId(
		ctx context.Context,
		tableId uuid.UUID,
	) (*hospitalitymodels.RestaurantTableBookingSession, error)

	GetBookingSessionById(
		ctx context.Context,
		sessionId uuid.UUID,
	) (*hospitalitymodels.RestaurantTableBookingSession, error)

	UpdateRestaurantOrderItemKotReference(
		ctx context.Context,
		orderItemID uuid.UUID,
		ticketID uuid.UUID,
		ticketItemID uuid.UUID,
	) error

	// WithTransaction(ctx context.Context, fn func(tx RestaurantRepository) error) error
}
