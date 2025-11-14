package restaurantservices

import (
	"context"
	hospitalitydtos "sharpnr-engine/modules/hospitality/hospitality_dtos"

	"github.com/google/uuid"
)

type RestaurantService interface {
	InitializeRestaurant()
	GetRestaurantDetails(ctx context.Context, orgId uuid.UUID) (*hospitalitydtos.Restaurant_GetRestaurant_ResponseDTO, error)
}
