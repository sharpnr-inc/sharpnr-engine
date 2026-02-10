package restaurantrepo

import (
	"context"
	restaurant_models "sharpnr-engine/modules/hospitality/restaurant/models"

	"github.com/google/uuid"
)

func (r *restaurantRepo) CreateSessionPayment(ctx context.Context, sessionPayment *restaurant_models.RestaurantSessionPayment) error {
	// .Create inserts the payment record.
	// Ensure sessionPayment.DiningSessionID is set before calling this.
	return r.db.WithContext(ctx).Create(sessionPayment).Error
}

func (r *restaurantRepo) GetAllPaymentsOfDiningSession(ctx context.Context, sessionID uuid.UUID) ([]*restaurant_models.RestaurantSessionPayment, error) {
	var payments []*restaurant_models.RestaurantSessionPayment

	// We filter by dining_session_id to get all payment attempts (success or failure).
	// Sorting by created_at helps track the payment sequence.
	err := r.db.WithContext(ctx).
		Where("session_id = ?", sessionID).
		Order("created_at ASC").
		Find(&payments).Error

	if err != nil {
		return nil, err
	}

	return payments, nil
}
