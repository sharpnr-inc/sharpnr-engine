package restaurantrepo

import (
	"context"
	hospitalitymodels "sharpnr-engine/modules/hospitality/hospitality_models"
)

func (r *restaurantRepo) UpdateRestaurant(ctx context.Context, restaurant *hospitalitymodels.Restaurant) error {
	tx := r.db.WithContext(ctx).Model(&hospitalitymodels.Restaurant{}).
		Where("id = ?", restaurant.ID).
		Updates(restaurant)

	return tx.Error
}
