package restaurantrepo

import (
	"context"
	"fmt"
	restaurant_models "sharpnr-engine/modules/hospitality/restaurant/models"

	"github.com/google/uuid"
)

func (r *restaurantRepo) CreateRestaurantDiningSession(ctx context.Context, diningSession *restaurant_models.RestaurantDiningSession) error {
	// Inserts the session record (e.g., table assigned, start time)
	return r.db.WithContext(ctx).Create(diningSession).Error
}

func (r *restaurantRepo) UpdateDiningSessionStatus(ctx context.Context, diningSessionID uuid.UUID, status string) error {
	result := r.db.WithContext(ctx).
		Model(&restaurant_models.RestaurantDiningSession{}).
		Where("id = ?", diningSessionID).
		Update("status", status)

	if result.Error != nil {
		return result.Error
	}

	if result.RowsAffected == 0 {
		return fmt.Errorf("dining session not found: %s", diningSessionID)
	}

	return nil
}
