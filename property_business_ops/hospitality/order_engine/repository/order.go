package hospitality_order_engine_repository

import (
	"context"
	"errors"
	"fmt"
	hospitality_order_engine_models "sharpnr-engine/property_business_ops/hospitality/order_engine/models"

	"github.com/google/uuid"
)

func (r *orderEngineRepo) CreateOrder(ctx context.Context, order *hospitality_order_engine_models.HospitalityOrder) error {
	// .Create inserts the order record.
	// If 'order' contains a slice of Items, GORM inserts them and
	// links them via foreign keys automatically.
	err := r.db.WithContext(ctx).Create(order).Error

	if err != nil {
		return fmt.Errorf("failed to create hospitality order: %w", err)
	}

	return nil
}

func (r *orderEngineRepo) GetOrderByID(ctx context.Context, orderID uuid.UUID) (*hospitality_order_engine_models.HospitalityOrder, error) {
	var order hospitality_order_engine_models.HospitalityOrder

	// .Preload("OrderItems") ensures that the slice of items is populated.
	// Replace "OrderItems" with the exact field name in your HospitalityOrder struct.
	err := r.db.WithContext(ctx).
		Preload("OrderItems").
		Where("id = ?", orderID).
		First(&order).Error

	if err != nil {
		return nil, err
	}

	return &order, nil
}

func (r *orderEngineRepo) GetOrdersByOwnersID(ctx context.Context, ownersID uuid.UUID) ([]*hospitality_order_engine_models.HospitalityOrder, error) {
	var orders []*hospitality_order_engine_models.HospitalityOrder

	// We filter by owners_id and usually want to see the most recent orders first.
	// We also Preload "OrderItems" so the UI has all the data it needs immediately.
	err := r.db.WithContext(ctx).
		Preload("OrderItems").
		Where("owners_id = ?", ownersID).
		Order("created_at DESC").
		Find(&orders).Error

	if err != nil {
		return nil, err
	}

	return orders, nil
}

func (r *orderEngineRepo) GetOrdersByPropertyID(ctx context.Context, propertyID uuid.UUID) ([]*hospitality_order_engine_models.HospitalityOrder, error) {
	var orders []*hospitality_order_engine_models.HospitalityOrder

	// We filter by property_id and usually want to see the most recent orders first.
	// We also Preload "OrderItems" so the UI has all the data it needs immediately.
	err := r.db.WithContext(ctx).
		Preload("OrderItems").
		Where("property_id = ?", propertyID).
		Order("created_at DESC").
		Find(&orders).Error

	if err != nil {
		return nil, err
	}

	return orders, nil
}

func (r *orderEngineRepo) UpdateOrderStatus(ctx context.Context, orderID uuid.UUID, status string) error {
	// We target the HospitalityOrder model and filter by its unique ID.
	// .Update() only modifies the "status" column and automatically refreshes "updated_at".
	result := r.db.WithContext(ctx).
		Model(&hospitality_order_engine_models.HospitalityOrder{}).
		Where("id = ?", orderID).
		Update("status", status)

	if result.Error != nil {
		return result.Error
	}

	// It's critical to check if a row was actually changed.
	// If RowsAffected is 0, the orderID provided doesn't exist.
	if result.RowsAffected == 0 {
		return errors.New("order not found")
	}

	return nil
}
