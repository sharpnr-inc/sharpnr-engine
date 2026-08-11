package restaurantrepo

import (
	"context"
	"errors"
	"fmt"
	restaurant_constants "sharpnr-engine/property_business_ops/hospitality/restaurant/constants"
	restaurant_models "sharpnr-engine/property_business_ops/hospitality/restaurant/models"
	dbtx "sharpnr-engine/pkg/dbTx"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

func (r *restaurantRepo) CreateRestaurantDiningSession(ctx context.Context, diningSession *restaurant_models.RestaurantDiningSession) error {
	db, ok := dbtx.ExtractTx(ctx)
	if !ok {
		db = r.db
	}

	if diningSession == nil {
		return fmt.Errorf("dining session data is nil")
	}

	// Inserts the session record (e.g., table assigned, start time)
	return db.WithContext(ctx).Create(diningSession).Error
}

/* ----------------------------------- <*> ---------------------------------- */

func (r *restaurantRepo) GetDiningSessionByID(ctx context.Context, sessionID uuid.UUID) (*restaurant_models.RestaurantDiningSession, error) {
	db, ok := dbtx.ExtractTx(ctx)
	if !ok {
		db = r.db
	}

	var session restaurant_models.RestaurantDiningSession

	// Use WithContext to ensure trace propagation and timeouts are honored
	err := db.WithContext(ctx).
		Where("id = ?", sessionID).
		First(&session).
		Error

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			// Return nil to indicate the record does not exist
			return nil, nil
		}
		return nil, fmt.Errorf("failed to fetch dining session by id: %w", err)
	}

	return &session, nil
}

/* ----------------------------------- <*> ---------------------------------- */

func (r *restaurantRepo) UpdateDiningSessionStatus(ctx context.Context, diningSessionID uuid.UUID, status string) error {
	db, ok := dbtx.ExtractTx(ctx)
	if !ok {
		db = r.db
	}

	result := db.WithContext(ctx).
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

/* ----------------------------------- <*> ---------------------------------- */

func (r *restaurantRepo) GetDiningSessionByStatus(
	ctx context.Context,
	restaurantID uuid.UUID,
	status restaurant_constants.RestaurantDiningSessionStatusType) ([]*restaurant_models.RestaurantDiningSession, error) {
	db, ok := dbtx.ExtractTx(ctx)
	if !ok {
		db = r.db
	}

	var sessions []*restaurant_models.RestaurantDiningSession

	// We filter by restaurant_id to ensure multi-tenancy security,
	// and by status (e.g., 'OPEN', 'COMPLETED', 'CANCELLED').
	err := db.WithContext(ctx).
		Where("restaurant_id = ? AND status = ?", restaurantID, status).
		Order("created_at DESC"). // Newest sessions first
		Find(&sessions).Error

	if err != nil {
		return nil, err
	}

	return sessions, nil
}

/* ----------------------------------- <*> ---------------------------------- */

func (r *restaurantRepo) GetDiningSessionByTableID(ctx context.Context, tableID uuid.UUID) (*restaurant_models.RestaurantDiningSession, error) {
	db, ok := dbtx.ExtractTx(ctx)
	if !ok {
		db = r.db
	}

	var session restaurant_models.RestaurantDiningSession

	// We filter by TableID and ensure the status is 'OPEN' (Running)
	// We use First to get the most relevant record and check for errors
	err := db.WithContext(ctx).
		Where("table_id = ? AND status = ?", tableID, restaurant_constants.DineInSessionOpened).
		Order("started_at DESC").
		First(&session).
		Error

	if err != nil {
		return nil, err
	}

	return &session, nil
}

/* ----------------------------------- <*> ---------------------------------- */

func (r *restaurantRepo) GetAllOrdersOfSession(ctx context.Context, dineSessionID uuid.UUID) error {
	return nil
}
