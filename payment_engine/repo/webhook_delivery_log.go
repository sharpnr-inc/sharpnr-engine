package payment_engine_repo

import (
	"context"
	"fmt"
	payment_engine_models "sharpnr-engine/payment_engine/models"
	dbtx "sharpnr-engine/pkg/dbTx"
	"time"

	"github.com/google/uuid"
)

func (r *paymentRepo) CreateWebhookDeliveryLog(ctx context.Context, log *payment_engine_models.WebhookDeliveryLog) error {
	db, ok := dbtx.ExtractTx(ctx)
	if !ok {
		db = r.db
	}

	err := db.WithContext(ctx).Create(log).Error
	if err != nil {
		return fmt.Errorf("failed to create webhook delivery log: %w", err)
	}

	return nil
}

func (r *paymentRepo) GetWebhookDeliveryLogByID(ctx context.Context, id uuid.UUID) (*payment_engine_models.WebhookDeliveryLog, error) {
	db, ok := dbtx.ExtractTx(ctx)
	if !ok {
		db = r.db
	}

	var log payment_engine_models.WebhookDeliveryLog
	err := db.WithContext(ctx).First(&log, "id = ?", id).Error
	if err != nil {
		return nil, fmt.Errorf("failed to find webhook delivery log %s: %w", id, err)
	}

	return &log, nil
}

func (r *paymentRepo) GetPendingWebhookRetries(ctx context.Context, before time.Time) ([]*payment_engine_models.WebhookDeliveryLog, error) {
	db, ok := dbtx.ExtractTx(ctx)
	if !ok {
		db = r.db
	}

	var logs []*payment_engine_models.WebhookDeliveryLog
	// Assuming status 'FAILED' or 'PENDING' with a retry timestamp
	err := db.WithContext(ctx).
		Where("status IN (?) AND next_retry_at < ? AND retry_count < max_retries",
			[]string{"FAILED", "PENDING"}, before).
		Find(&logs).Error
	if err != nil {
		return nil, fmt.Errorf("failed to fetch pending webhook retries: %w", err)
	}

	return logs, nil
}

func (r *paymentRepo) GetDeliveryLogsByEndpoint(ctx context.Context, endpointID uuid.UUID) ([]*payment_engine_models.WebhookDeliveryLog, error) {
	db, ok := dbtx.ExtractTx(ctx)
	if !ok {
		db = r.db
	}

	var logs []*payment_engine_models.WebhookDeliveryLog
	err := db.WithContext(ctx).
		Where("webhook_endpoint_id = ?", endpointID).
		Order("created_at DESC").
		Find(&logs).Error
	if err != nil {
		return nil, fmt.Errorf("failed to list logs for endpoint %s: %w", endpointID, err)
	}

	return logs, nil
}

func (r *paymentRepo) UpdateWebhookDeliveryLog(ctx context.Context, log *payment_engine_models.WebhookDeliveryLog) error {
	db, ok := dbtx.ExtractTx(ctx)
	if !ok {
		db = r.db
	}

	// Update the full log object, including status, response, and retry metadata
	err := db.WithContext(ctx).Model(log).Updates(log).Error
	if err != nil {
		return fmt.Errorf("failed to update delivery log %s: %w", log.ID, err)
	}

	return nil
}
