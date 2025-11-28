package restaurantrepo

import (
	"context"
	"errors"
	"fmt"
	hospitalitymodels "sharpnr-engine/modules/hospitality/hospitality_models"
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

func (d *restaurantRepo) CreateRes_TableSession(ctx context.Context, session *hospitalitymodels.RestaurantTableBookingSession) error {
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
	session.UpdatedAt = time.Now()

	result := d.db.WithContext(ctx).Create(session)
	if result.Error != nil {
		return result.Error
	}

	return nil
}

func (d *restaurantRepo) CreateOrderInTableSession(ctx context.Context, tableOrder *hospitalitymodels.RestaurantTableOrder) error {
	// generate UUID if empty
	if tableOrder.ID == uuid.Nil {
		tableOrder.ID = uuid.New()
	}

	// insert table order
	if err := d.db.WithContext(ctx).Create(tableOrder).Error; err != nil {
		return fmt.Errorf("repo: failed to create table order: %w", err)
	}

	return nil
}

func (d *restaurantRepo) GetBookingSessionByTableId(
	ctx context.Context,
	tableId uuid.UUID,
) (*hospitalitymodels.RestaurantTableBookingSession, error) {

	var session hospitalitymodels.RestaurantTableBookingSession

	err := d.db.WithContext(ctx).
		Preload("RestaurantTable").
		Preload("CustomerAndGuest").
		Preload("Orders").
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
) (*hospitalitymodels.RestaurantTableBookingSession, error) {

	var session hospitalitymodels.RestaurantTableBookingSession

	err := d.db.WithContext(ctx).
		Preload("RestaurantTable").
		Preload("CustomerAndGuest").
		Preload("Orders").
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

func (d *restaurantRepo) CreateRestaurantTableOrder(
	ctx context.Context,
	resOrder *hospitalitymodels.RestaurantTableOrder) (*hospitalitymodels.RestaurantTableOrder, error) {
	// Create the order row
	if err := d.db.WithContext(ctx).Create(resOrder).Error; err != nil {
		return nil, err
	}

	return resOrder, nil
}
