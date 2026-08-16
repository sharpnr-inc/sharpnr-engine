package hospitality_order_engine_repository

import (
	"context"
	"fmt"
	hospitality_order_engine_models "sharpnr-engine/property_business_ops/hospitality/order_engine/models"
)

func (r *orderEngineRepo) CreateOrderItemVariant(ctx context.Context, itemVariant *hospitality_order_engine_models.HospitalityOrderItemVariant) error {
	// .Create inserts the variant record.
	// Ensure itemVariant.HospitalityOrderItemID is set to link it to the specific order item.
	err := r.db.WithContext(ctx).Create(itemVariant).Error

	if err != nil {
		return fmt.Errorf("failed to create order item variant: %w", err)
	}

	return nil
}
