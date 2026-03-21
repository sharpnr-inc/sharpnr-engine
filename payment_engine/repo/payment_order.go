package payment_engine_repo

import (
	"context"
	"fmt"
	dbtx "sharpnr-engine/core/dbTx"
	payment_engine_models "sharpnr-engine/payment_engine/models"
	"time"

	"github.com/google/uuid"
)

func (r *paymentRepo) CreateOrder(ctx context.Context, order *payment_engine_models.PaymentOrder) error {
	db, ok := dbtx.ExtractTx(ctx)
	if !ok {
		db = r.db
	}

	err := db.WithContext(ctx).Create(order).Error
	if err != nil {
		return fmt.Errorf("failed to create order: %w", err)
	}

	return nil
}

func (r *paymentRepo) GetOrderByID(ctx context.Context, id uuid.UUID) (*payment_engine_models.PaymentOrder, error) {
	db, ok := dbtx.ExtractTx(ctx)
	if !ok {
		db = r.db
	}

	var order payment_engine_models.PaymentOrder
	err := db.WithContext(ctx).First(&order, "id = ?", id).Error
	if err != nil {
		return nil, fmt.Errorf("failed to find order %s: %w", id, err)
	}

	return &order, nil
}

func (r *paymentRepo) GetOrderByNumber(ctx context.Context, tenantID uuid.UUID, orderNumber string) (*payment_engine_models.PaymentOrder, error) {
	db, ok := dbtx.ExtractTx(ctx)
	if !ok {
		db = r.db
	}

	var order payment_engine_models.PaymentOrder
	err := db.WithContext(ctx).Where("tenant_id = ? AND order_number = ?", tenantID, orderNumber).First(&order).Error
	if err != nil {
		return nil, fmt.Errorf("failed to find order number %s: %w", orderNumber, err)
	}

	return &order, nil
}

func (r *paymentRepo) ListOrdersBySource(ctx context.Context, tenantID uuid.UUID, sourceType string, sourceRefID string) ([]*payment_engine_models.PaymentOrder, error) {
	db, ok := dbtx.ExtractTx(ctx)
	if !ok {
		db = r.db
	}

	var orders []*payment_engine_models.PaymentOrder
	err := db.WithContext(ctx).
		Where("tenant_id = ? AND source_type = ? AND source_ref_id = ?", tenantID, sourceType, sourceRefID).
		Find(&orders).Error
	if err != nil {
		return nil, fmt.Errorf("failed to list orders for source %s: %w", sourceRefID, err)
	}

	return orders, nil
}

func (r *paymentRepo) UpdateOrderStatus(ctx context.Context, id uuid.UUID, status string) error {
	db, ok := dbtx.ExtractTx(ctx)
	if !ok {
		db = r.db
	}

	err := db.WithContext(ctx).Model(&payment_engine_models.PaymentOrder{}).
		Where("id = ?", id).
		Update("status", status).Error
	if err != nil {
		return fmt.Errorf("failed to update order status for %s: %w", id, err)
	}

	return nil
}

func (r *paymentRepo) UpdateOrderPaidAmount(ctx context.Context, id uuid.UUID, paidPaise int64, newStatus string) error {
	db, ok := dbtx.ExtractTx(ctx)
	if !ok {
		db = r.db
	}

	err := db.WithContext(ctx).Model(&payment_engine_models.PaymentOrder{}).
		Where("id = ?", id).
		Updates(map[string]interface{}{
			"paid_paise": paidPaise,
			"status":     newStatus,
		}).Error
	if err != nil {
		return fmt.Errorf("failed to update paid amount for order %s: %w", id, err)
	}

	return nil
}

func (r *paymentRepo) GetExpiredPendingOrders(ctx context.Context, before time.Time) ([]*payment_engine_models.PaymentOrder, error) {
	db, ok := dbtx.ExtractTx(ctx)
	if !ok {
		db = r.db
	}

	var orders []*payment_engine_models.PaymentOrder
	// Assuming 'PENDING' is your status constant and 'expires_at' or 'created_at' is the threshold
	err := db.WithContext(ctx).
		Where("status = ? AND created_at < ?", "PENDING", before).
		Find(&orders).Error
	if err != nil {
		return nil, fmt.Errorf("failed to fetch expired orders: %w", err)
	}

	return orders, nil
}
