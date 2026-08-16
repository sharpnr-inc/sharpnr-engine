package hospitality_order_engine_repository

import (
	"context"
	"fmt"
	hospitality_order_engine_models "sharpnr-engine/property_business_ops/hospitality/order_engine/models"

	"github.com/google/uuid"
)

func (r *orderEngineRepo) CreateOrderItem(ctx context.Context, orderItem *hospitality_order_engine_models.HospitalityOrderItem) error {
	// .Create inserts the specific line item into the database.
	// Ensure orderItem.HospitalityOrderID is set to link it to the parent order.
	err := r.db.WithContext(ctx).Create(orderItem).Error

	if err != nil {
		return fmt.Errorf("failed to create order item: %w", err)
	}

	return nil
}

func (r *orderEngineRepo) GetOrderItemByOrderID(ctx context.Context, orderID uuid.UUID) ([]*hospitality_order_engine_models.HospitalityOrderItem, error) {
	var orderItems []*hospitality_order_engine_models.HospitalityOrderItem

	// We filter by hospitality_order_id to fetch all items linked to the specific order.
	// Ordering by created_at ensures items appear in the order they were added.
	err := r.db.WithContext(ctx).
		Where("order_id = ?", orderID).
		Order("created_at ASC").
		Find(&orderItems).Error

	if err != nil {
		return nil, err
	}

	return orderItems, nil
}
