package restaurantrepo

import (
	"context"
	"fmt"
	hospitalitymodels "sharpnr-engine/modules/hospitality/hospitality_models"
)

// CreateRestaurant implements RestaurantRepository.
func (r *restaurantRepo) CreateRestaurant(ctx context.Context, res *hospitalitymodels.Restaurant) error {
	if res == nil {
		return fmt.Errorf("restaurant data is nil")
	}
	// Insert record into database
	if err := r.db.WithContext(ctx).Create(res).Error; err != nil {
		return fmt.Errorf("failed to create restaurant: %w", err)
	}

	return nil
}
