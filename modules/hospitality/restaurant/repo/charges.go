package restaurantrepo

import (
	"context"
	restaurant_models "sharpnr-engine/modules/hospitality/restaurant/models"

	"github.com/google/uuid"
)

func (r *restaurantRepo) CreateCharges(ctx context.Context, charge *restaurant_models.RestaurantCharge) error {
	// .Create inserts the charge (e.g., a "Ribeye Steak" or "Service Charge").
	// GORM will populate ID and Timestamps.
	return r.db.WithContext(ctx).Create(charge).Error
}

func (r *restaurantRepo) GetAllChargesOfSession(ctx context.Context, sessionID uuid.UUID) ([]*restaurant_models.RestaurantCharge, error) {
	var charges []*restaurant_models.RestaurantCharge

	// We filter by dining_session_id.
	// Charges are usually ordered by time created so they appear on the receipt
	// in the order they were served/ordered.
	err := r.db.WithContext(ctx).
		Where("session_id = ?", sessionID).
		Order("created_at ASC").
		Find(&charges).Error

	if err != nil {
		return nil, err
	}

	return charges, nil
}
