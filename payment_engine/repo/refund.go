package payment_engine_repo

import (
	"context"
	"fmt"
	payment_engine_models "sharpnr-engine/payment_engine/models"
	dbtx "sharpnr-engine/pkg/dbTx"

	"github.com/google/uuid"
)

func (r *paymentRepo) CreateRefund(ctx context.Context, refund *payment_engine_models.Refund) error {
	db, ok := dbtx.ExtractTx(ctx)
	if !ok {
		db = r.db
	}

	err := db.WithContext(ctx).Create(refund).Error
	if err != nil {
		return fmt.Errorf("failed to create refund record: %w", err)
	}

	return nil
}

func (r *paymentRepo) GetRefundByID(ctx context.Context, id uuid.UUID) (*payment_engine_models.Refund, error) {
	db, ok := dbtx.ExtractTx(ctx)
	if !ok {
		db = r.db
	}

	var refund payment_engine_models.Refund
	err := db.WithContext(ctx).First(&refund, "id = ?", id).Error
	if err != nil {
		return nil, fmt.Errorf("failed to find refund %s: %w", id, err)
	}

	return &refund, nil
}

func (r *paymentRepo) GetRefundsByPaymentID(ctx context.Context, paymentID uuid.UUID) ([]*payment_engine_models.Refund, error) {
	db, ok := dbtx.ExtractTx(ctx)
	if !ok {
		db = r.db
	}

	var refunds []*payment_engine_models.Refund
	err := db.WithContext(ctx).Where("payment_id = ?", paymentID).Find(&refunds).Error
	if err != nil {
		return nil, fmt.Errorf("failed to list refunds for payment %s: %w", paymentID, err)
	}

	return refunds, nil
}

func (r *paymentRepo) UpdateRefundStatus(ctx context.Context, id uuid.UUID, status string) error {
	db, ok := dbtx.ExtractTx(ctx)
	if !ok {
		db = r.db
	}

	err := db.WithContext(ctx).Model(&payment_engine_models.Refund{}).
		Where("id = ?", id).
		Update("status", status).Error
	if err != nil {
		return fmt.Errorf("failed to update refund status for %s: %w", id, err)
	}

	return nil
}

func (r *paymentRepo) UpdateRefundGatewayRefs(ctx context.Context, refund *payment_engine_models.Refund) error {
	db, ok := dbtx.ExtractTx(ctx)
	if !ok {
		db = r.db
	}

	// Specifically update gateway-related fields and the update timestamp
	err := db.WithContext(ctx).Model(refund).
		Select("gateway_refund_ref", "gateway_response_raw", "status", "updated_at").
		Updates(refund).Error
	if err != nil {
		return fmt.Errorf("failed to update gateway refs for refund %s: %w", refund.ID, err)
	}

	return nil
}
