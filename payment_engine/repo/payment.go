package payment_engine_repo

import (
	"context"
	"fmt"
	dbtx "sharpnr-engine/core/dbTx"
	payment_engine_models "sharpnr-engine/payment_engine/models"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

func (r *paymentRepo) CreatePayment(ctx context.Context, payment *payment_engine_models.Payment) error {
	db, ok := dbtx.ExtractTx(ctx)
	if !ok {
		db = r.db
	}

	err := db.WithContext(ctx).Create(payment).Error
	if err != nil {
		return fmt.Errorf("failed to create payment: %w", err)
	}

	return nil
}

func (r *paymentRepo) GetPaymentByID(ctx context.Context, id uuid.UUID) (*payment_engine_models.Payment, error) {
	db, ok := dbtx.ExtractTx(ctx)
	if !ok {
		db = r.db
	}

	var payment payment_engine_models.Payment
	err := db.WithContext(ctx).First(&payment, "id = ?", id).Error
	if err != nil {
		return nil, fmt.Errorf("failed to find payment %s: %w", id, err)
	}

	return &payment, nil
}

func (r *paymentRepo) GetPaymentByGatewayOrderRef(ctx context.Context, gatewayOrderRef string) (*payment_engine_models.Payment, error) {
	db, ok := dbtx.ExtractTx(ctx)
	if !ok {
		db = r.db
	}

	var payment payment_engine_models.Payment
	err := db.WithContext(ctx).Where("gateway_order_ref = ?", gatewayOrderRef).First(&payment).Error
	if err != nil {
		return nil, fmt.Errorf("failed to find payment by gateway order ref %s: %w", gatewayOrderRef, err)
	}

	return &payment, nil
}

func (r *paymentRepo) GetPaymentByGatewayTxnRef(ctx context.Context, gatewayTxnRef string) (*payment_engine_models.Payment, error) {
	db, ok := dbtx.ExtractTx(ctx)
	if !ok {
		db = r.db
	}

	var payment payment_engine_models.Payment
	err := db.WithContext(ctx).Where("gateway_txn_ref = ?", gatewayTxnRef).First(&payment).Error
	if err != nil {
		return nil, fmt.Errorf("failed to find payment by gateway txn ref %s: %w", gatewayTxnRef, err)
	}

	return &payment, nil
}

func (r *paymentRepo) GetPaymentsByOrderID(ctx context.Context, orderID uuid.UUID) ([]*payment_engine_models.Payment, error) {
	db, ok := dbtx.ExtractTx(ctx)
	if !ok {
		db = r.db
	}

	var payments []*payment_engine_models.Payment
	err := db.WithContext(ctx).Where("order_id = ?", orderID).Find(&payments).Error
	if err != nil {
		return nil, fmt.Errorf("failed to find payments for order %s: %w", orderID, err)
	}

	return payments, nil
}

func (r *paymentRepo) UpdatePaymentStatus(ctx context.Context, id uuid.UUID, status string) error {
	db, ok := dbtx.ExtractTx(ctx)
	if !ok {
		db = r.db
	}

	err := db.WithContext(ctx).Model(&payment_engine_models.Payment{}).
		Where("id = ?", id).
		Update("status", status).Error
	if err != nil {
		return fmt.Errorf("failed to update status for payment %s: %w", id, err)
	}

	return nil
}

func (r *paymentRepo) UpdatePaymentGatewayRefs(
	ctx context.Context,
	payment *payment_engine_models.Payment,
) error {

	db, ok := dbtx.ExtractTx(ctx)
	if !ok {
		db = r.db
	}

	return db.Model(&payment_engine_models.Payment{}).
		Where("id = ?", payment.ID).
		Updates(map[string]any{
			"gateway_order_ref": payment.GatewayOrderRef,
			"gateway_response":  payment.GatewayResponse,
			"status":            payment.Status,
		}).Error
}

func (r *paymentRepo) UpdatePaymentConfirmation(ctx context.Context, payment *payment_engine_models.Payment) error {
	db, ok := dbtx.ExtractTx(ctx)
	if !ok {
		db = r.db
	}

	// This updates the status and relevant timestamps when a webhook/callback confirms the payment
	err := db.WithContext(ctx).Model(payment).
		Select("status", "confirmed_at", "gateway_response_raw", "updated_at").
		Updates(payment).Error
	if err != nil {
		return fmt.Errorf("failed to confirm payment %s: %w", payment.ID, err)
	}

	return nil
}

func (r *paymentRepo) IncrementPaymentRefunded(ctx context.Context, id uuid.UUID, amountPaise int64) error {
	db, ok := dbtx.ExtractTx(ctx)
	if !ok {
		db = r.db
	}

	// Atomic increment using gorm.Expr to prevent race conditions
	err := db.WithContext(ctx).Model(&payment_engine_models.Payment{}).
		Where("id = ?", id).
		Update("refunded_paise", gorm.Expr("refunded_paise + ?", amountPaise)).Error
	if err != nil {
		return fmt.Errorf("failed to increment refunded amount for payment %s: %w", id, err)
	}

	return nil
}
