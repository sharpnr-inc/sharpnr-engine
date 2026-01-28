package restaurantservices

import (
	"context"

	"github.com/google/uuid"
)

/****************************************
 * Service Interface for the Restuarant *
 *                                      *
****************************************/

type RestaurantServiceInterface interface {
	SetupDefaultRestaurant(ctx context.Context, orgID uuid.UUID, propertyID uuid.UUID) (restuarantID uuid.UUID, err error)
}
