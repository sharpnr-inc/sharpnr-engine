package payment_engine_repo

import (
	"context"
	"fmt"
	payment_engine_models "sharpnr-engine/payment_engine/models"
	dbtx "sharpnr-engine/pkg/dbTx"
	"time"

	"github.com/google/uuid"
)

func (r *paymentRepo) CreateAllocation(ctx context.Context, allocation *payment_engine_models.PaymentAllocation) error {
	db, ok := dbtx.ExtractTx(ctx)
	if !ok {
		db = r.db
	}

	err := db.WithContext(ctx).Create(allocation).Error
	if err != nil {
		return fmt.Errorf("failed to create allocation: %w", err)
	}

	return nil
}

func (r *paymentRepo) GetAllocationByID(ctx context.Context, id uuid.UUID) (*payment_engine_models.PaymentAllocation, error) {
	db, ok := dbtx.ExtractTx(ctx)
	if !ok {
		db = r.db
	}

	var allocation payment_engine_models.PaymentAllocation
	err := db.WithContext(ctx).First(&allocation, "id = ?", id).Error
	if err != nil {
		return nil, fmt.Errorf("failed to find allocation %s: %w", id, err)
	}

	return &allocation, nil
}

func (r *paymentRepo) GetAllocationsByPaymentID(ctx context.Context, paymentID uuid.UUID) ([]*payment_engine_models.PaymentAllocation, error) {
	db, ok := dbtx.ExtractTx(ctx)
	if !ok {
		db = r.db
	}

	var allocations []*payment_engine_models.PaymentAllocation
	err := db.WithContext(ctx).Where("payment_id = ?", paymentID).Find(&allocations).Error
	if err != nil {
		return nil, fmt.Errorf("failed to list allocations for payment %s: %w", paymentID, err)
	}

	return allocations, nil
}

func (r *paymentRepo) GetAllocationsByTarget(ctx context.Context, targetType string, targetID string) ([]*payment_engine_models.PaymentAllocation, error) {
	db, ok := dbtx.ExtractTx(ctx)
	if !ok {
		db = r.db
	}

	var allocations []*payment_engine_models.PaymentAllocation
	err := db.WithContext(ctx).Where("target_type = ? AND target_id = ?", targetType, targetID).Find(&allocations).Error
	if err != nil {
		return nil, fmt.Errorf("failed to find allocations for target %s/%s: %w", targetType, targetID, err)
	}

	return allocations, nil
}

func (r *paymentRepo) GetTotalAllocatedByPayment(ctx context.Context, paymentID uuid.UUID) (int64, error) {
	db, ok := dbtx.ExtractTx(ctx)
	if !ok {
		db = r.db
	}

	var total int64
	err := db.WithContext(ctx).Model(&payment_engine_models.PaymentAllocation{}).
		Where("payment_id = ? AND is_reversed = ?", paymentID, false).
		Select("COALESCE(SUM(allocated_paise), 0)").
		Scan(&total).Error

	if err != nil {
		return 0, fmt.Errorf("failed to calculate total allocation for payment %s: %w", paymentID, err)
	}

	return total, nil
}

func (r *paymentRepo) GetTotalAllocatedByTarget(ctx context.Context, targetType string, targetID string) (int64, error) {
	db, ok := dbtx.ExtractTx(ctx)
	if !ok {
		db = r.db
	}

	var total int64
	err := db.WithContext(ctx).Model(&payment_engine_models.PaymentAllocation{}).
		Where("target_type = ? AND target_id = ? AND is_reversed = ?", targetType, targetID, false).
		Select("COALESCE(SUM(allocated_paise), 0)").
		Scan(&total).Error

	if err != nil {
		return 0, fmt.Errorf("failed to calculate total allocation for target %s/%s: %w", targetType, targetID, err)
	}

	return total, nil
}

func (r *paymentRepo) ReverseAllocation(ctx context.Context, id uuid.UUID, by string, note string) error {
	db, ok := dbtx.ExtractTx(ctx)
	if !ok {
		db = r.db
	}

	result := db.WithContext(ctx).Model(&payment_engine_models.PaymentAllocation{}).
		Where("id = ?", id).
		Updates(map[string]interface{}{
			"is_reversed":   true,
			"reversed_at":   time.Now(),
			"reversed_by":   by,
			"reversal_note": note,
		})

	if result.Error != nil {
		return fmt.Errorf("failed to reverse allocation %s: %w", id, result.Error)
	}

	if result.RowsAffected == 0 {
		return fmt.Errorf("no allocation found with id %s to reverse", id)
	}

	return nil
}
