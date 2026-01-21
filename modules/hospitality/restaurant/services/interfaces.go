package restaurantservices

import (
	"context"
)

/****************************************
 * Service Interface for the Restuarant *
 *                                      *
****************************************/

type RestaurantServiceInterface interface {
	CreateFloor(ctx context.Context)
}
