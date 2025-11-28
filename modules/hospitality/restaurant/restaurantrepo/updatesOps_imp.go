package restaurantrepo

import (
	"context"
	"fmt"
	hospitalitymodels "sharpnr-engine/modules/hospitality/hospitality_models"

	"github.com/google/uuid"
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

func (r *restaurantRepo) UpdateTable(ctx context.Context, tableId uuid.UUID, updates map[string]any) error {
	return r.db.WithContext(ctx).
		Model(&hospitalitymodels.RestaurantTable{}).
		Where("id = ?", tableId).
		Updates(updates).
		Error
}

func (d *restaurantRepo) UpdateRestaurantOrderItemKotReference(
	ctx context.Context,
	orderItemID uuid.UUID,
	ticketID uuid.UUID,
	ticketItemID uuid.UUID,
) error {

	// We only update specific fields, nothing else.
	updateData := map[string]interface{}{
		"kitchen_ticket_id":      ticketID,
		"kitchen_ticket_item_id": ticketItemID,
	}

	// Update the matching row
	if err := d.db.WithContext(ctx).
		Model(&hospitalitymodels.RestaurantTableOrder{}).
		Where("id = ?", orderItemID).
		Updates(updateData).
		Error; err != nil {

		return err
	}

	return nil
}
