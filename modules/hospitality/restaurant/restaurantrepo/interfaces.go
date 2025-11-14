package restaurantrepo

import (
	"context"
	hospitalitymodels "sharpnr-engine/modules/hospitality/hospitality_models"

	"github.com/google/uuid"
)

type RestaurantRepository interface {
	CreateRestaurant(ctx context.Context, res *hospitalitymodels.Restaurant) error

	GetRestaurant(ctx context.Context, orgId uuid.UUID) (*hospitalitymodels.Restaurant, error)
}
