package restaurantrepo

import (
	"context"
	"fmt"
	hospitalitymodels "sharpnr-engine/modules/hospitality/hospitality_models"
)

func (r *restaurantRepo) UpdateRestaurant(ctx context.Context, restaurant *hospitalitymodels.Restaurant) error {
	// Only update fields that should be editable
	result := r.db.WithContext(ctx).
		Model(&hospitalitymodels.Restaurant{}).
		Where("id = ?", restaurant.ID).
		Select(
			"legal_name",
			"display_name",
			"fssai_number",
			"gst_number",
			"cin_number",
			"address",
			"city",
			"state",
			"country",
			"pincode",
			"contact_email",
			"contact_number",
			"manager_name",
			"updated_at", // GORM auto-updates this
		).
		Updates(restaurant)

	if result.Error != nil {
		return result.Error
	}

	// Check if restaurant exists
	if result.RowsAffected == 0 {
		return fmt.Errorf("restaurant not found")
	}

	return nil
}
