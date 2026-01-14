package restaurantrepo

import (
	"context"
	"fmt"
	restaurant_models "sharpnr-engine/modules/hospitality/restaurant/models"
)

// CreateRestaurant implements RestaurantRepository.
func (r *restaurantRepo) CreateRestaurant(ctx context.Context, res *restaurant_models.Restaurant) error {
	if res == nil {
		return fmt.Errorf("restaurant data is nil")
	}
	// Insert record into database
	if err := r.db.WithContext(ctx).Create(res).Error; err != nil {
		return fmt.Errorf("failed to create restaurant: %w", err)
	}

	return nil
}

// -----------------------------------------------------------------------------------------------------------------------------------//
func (r *restaurantRepo) CreateOrAddTablesInRestuarant(ctx context.Context,
	tables []*restaurant_models.RestaurantTable,
) error {

	return r.db.WithContext(ctx).Create(&tables).Error
}

//-----------------------------------------------------------------------------------------------------------------------------------//

func (r *restaurantRepo) CreateOrAddFloorInRestaurant(
	ctx context.Context,
	floor *restaurant_models.RestaurantFloor,
) error {

	return r.db.WithContext(ctx).Create(floor).Error
}
