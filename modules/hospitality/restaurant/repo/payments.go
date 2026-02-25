package restaurantrepo

import (
	"context"
	"fmt"
	dbtx "sharpnr-engine/core/dbTx"
	restaurant_models "sharpnr-engine/modules/hospitality/restaurant/models"

	"github.com/google/uuid"
)

func (r *restaurantRepo) CreateSessionPayment(ctx context.Context, sessionPayment *restaurant_models.RestaurantSessionPayment) error {
	db, ok := dbtx.ExtractTx(ctx)
	if !ok {
		db = r.db
	}

	if sessionPayment == nil {
		return fmt.Errorf("session payment data is nil")
	}

	// .Create inserts the payment record.
	// Ensure sessionPayment.DiningSessionID is set before calling this.
	return db.WithContext(ctx).Create(sessionPayment).Error
}

func (r *restaurantRepo) GetAllPaymentsOfDiningSession(ctx context.Context, sessionID uuid.UUID) ([]*restaurant_models.RestaurantSessionPayment, error) {
	db, ok := dbtx.ExtractTx(ctx)
	if !ok {
		db = r.db
	}

	var payments []*restaurant_models.RestaurantSessionPayment

	// We filter by session_id to get all payment attempts (success or failure).
	// Sorting by created_at helps track the payment sequence.
	err := db.WithContext(ctx).
		Where("session_id = ?", sessionID).
		Order("created_at ASC").
		Find(&payments).Error

	if err != nil {
		return nil, err
	}

	return payments, nil
}
