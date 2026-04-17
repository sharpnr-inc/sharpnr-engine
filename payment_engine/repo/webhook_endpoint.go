package payment_engine_repo

import (
	"context"
	"fmt"
	dbtx "sharpnr-engine/core/dbTx"
	payment_engine_models "sharpnr-engine/payment_engine/models"

	"github.com/google/uuid"
)

func (r *paymentRepo) CreateWebhookEndpoint(ctx context.Context, endpoint *payment_engine_models.WebhookEndpoint) error {
	db, ok := dbtx.ExtractTx(ctx)
	if !ok {
		db = r.db
	}

	err := db.WithContext(ctx).Create(endpoint).Error
	if err != nil {
		return fmt.Errorf("failed to create webhook endpoint: %w", err)
	}

	return nil
}

func (r *paymentRepo) GetWebhookEndpointByID(ctx context.Context, id uuid.UUID) (*payment_engine_models.WebhookEndpoint, error) {
	db, ok := dbtx.ExtractTx(ctx)
	if !ok {
		db = r.db
	}

	var endpoint payment_engine_models.WebhookEndpoint
	err := db.WithContext(ctx).First(&endpoint, "id = ?", id).Error
	if err != nil {
		return nil, fmt.Errorf("failed to find webhook endpoint %s: %w", id, err)
	}

	return &endpoint, nil
}

func (r *paymentRepo) GetActiveWebhookEndpoints(ctx context.Context, tenantID uuid.UUID, eventType string) ([]*payment_engine_models.WebhookEndpoint, error) {
	db, ok := dbtx.ExtractTx(ctx)
	if !ok {
		db = r.db
	}

	var endpoints []*payment_engine_models.WebhookEndpoint
	// This assumes subscribed_events is stored in a way that supports a LIKE query or contains the eventType
	err := db.WithContext(ctx).
		Where("tenant_id = ? AND is_active = ? AND subscribed_events LIKE ?", tenantID, true, "%"+eventType+"%").
		Find(&endpoints).Error
	if err != nil {
		return nil, fmt.Errorf("failed to list active webhook endpoints for event %s: %w", eventType, err)
	}

	return endpoints, nil
}

func (r *paymentRepo) ListWebhookEndpoints(ctx context.Context, tenantID uuid.UUID) ([]*payment_engine_models.WebhookEndpoint, error) {
	db, ok := dbtx.ExtractTx(ctx)
	if !ok {
		db = r.db
	}

	var endpoints []*payment_engine_models.WebhookEndpoint
	err := db.WithContext(ctx).Where("tenant_id = ?", tenantID).Find(&endpoints).Error
	if err != nil {
		return nil, fmt.Errorf("failed to list webhook endpoints for tenant %s: %w", tenantID, err)
	}

	return endpoints, nil
}

func (r *paymentRepo) UpdateWebhookEndpoint(ctx context.Context, id uuid.UUID, url string, subscribedEvents string) error {
	db, ok := dbtx.ExtractTx(ctx)
	if !ok {
		db = r.db
	}

	err := db.WithContext(ctx).Model(&payment_engine_models.WebhookEndpoint{}).
		Where("id = ?", id).
		Updates(map[string]interface{}{
			"url":               url,
			"subscribed_events": subscribedEvents,
		}).Error
	if err != nil {
		return fmt.Errorf("failed to update webhook endpoint %s: %w", id, err)
	}

	return nil
}

func (r *paymentRepo) ToggleWebhookEndpoint(ctx context.Context, id uuid.UUID, isActive bool) error {
	db, ok := dbtx.ExtractTx(ctx)
	if !ok {
		db = r.db
	}

	err := db.WithContext(ctx).Model(&payment_engine_models.WebhookEndpoint{}).
		Where("id = ?", id).
		Update("is_active", isActive).Error
	if err != nil {
		return fmt.Errorf("failed to toggle webhook status for %s: %w", id, err)
	}

	return nil
}

func (r *paymentRepo) DeleteWebhookEndpoint(ctx context.Context, id uuid.UUID) error {
	db, ok := dbtx.ExtractTx(ctx)
	if !ok {
		db = r.db
	}

	err := db.WithContext(ctx).Delete(&payment_engine_models.WebhookEndpoint{}, "id = ?", id).Error
	if err != nil {
		return fmt.Errorf("failed to delete webhook endpoint %s: %w", id, err)
	}

	return nil
}
