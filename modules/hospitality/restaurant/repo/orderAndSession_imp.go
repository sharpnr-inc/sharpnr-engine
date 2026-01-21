package restaurantrepo

import (
	"context"
	"errors"
	"fmt"
	restaurant_models "sharpnr-engine/modules/hospitality/restaurant/models"
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

func (d *restaurantRepo) CreateRes_TableSession(ctx context.Context, session *restaurant_models.RestaurantDiningSession) error {
	// If you're passing zero UUIDs or missing required fields,
	// GORM won't magically fix your life. So validate as needed.
	if session == nil {
		return fmt.Errorf("session is nil")
	}

	// Set defaults if needed
	if session.ID == uuid.Nil {
		session.ID = uuid.New()
	}
	session.CreatedAt = time.Now()

	result := d.db.WithContext(ctx).Create(session)
	if result.Error != nil {
		return result.Error
	}

	return nil
}

func (d *restaurantRepo) GetBookingSessionByTableId(
	ctx context.Context,
	tableId uuid.UUID,
) (*restaurant_models.RestaurantDiningSession, error) {

	var session restaurant_models.RestaurantDiningSession

	err := d.db.WithContext(ctx).
		Preload("RestaurantTable").
		Preload("CustomerAndGuest").
		Preload("KOTs").
		Where("restaurant_table_id = ? AND to_time IS NULL", tableId).
		First(&session).Error

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil // no active session yet
		}
		return nil, err
	}

	return &session, nil
}

func (d *restaurantRepo) GetBookingSessionById(
	ctx context.Context,
	sessionId uuid.UUID,
) (*restaurant_models.RestaurantDiningSession, error) {

	var session restaurant_models.RestaurantDiningSession

	err := d.db.WithContext(ctx).
		Preload("CustomerAndGuest").
		Preload("KOTs").
		Where("id = ? AND to_time IS NULL", sessionId).
		First(&session).Error

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil // no active session yet
		}
		return nil, err
	}

	return &session, nil
}
