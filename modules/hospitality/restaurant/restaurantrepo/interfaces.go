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
}
