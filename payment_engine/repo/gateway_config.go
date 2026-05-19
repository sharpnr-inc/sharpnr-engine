package payment_engine_repo

import (
	"context"
	"fmt"
	payment_engine_models "sharpnr-engine/payment_engine/models"
	dbtx "sharpnr-engine/pkg/dbTx"

	"github.com/google/uuid"
)

func (r *paymentRepo) CreateGatewayConfig(ctx context.Context, config *payment_engine_models.PaymentGatewayConfig) error {
	db, ok := dbtx.ExtractTx(ctx)
	if !ok {
		db = r.db
	}

	err := db.WithContext(ctx).Create(config).Error
	if err != nil {
		return fmt.Errorf("failed to create gateway config: %w", err)
	}

	return nil
}

func (r *paymentRepo) GetGatewayConfigByID(ctx context.Context, id uuid.UUID) (*payment_engine_models.PaymentGatewayConfig, error) {
	db, ok := dbtx.ExtractTx(ctx)
	if !ok {
		db = r.db
	}

	var config payment_engine_models.PaymentGatewayConfig
	err := db.WithContext(ctx).First(&config, "id = ?", id).Error
	if err != nil {
		return nil, fmt.Errorf("failed to find gateway config %s: %w", id, err)
	}

	return &config, nil
}

func (r *paymentRepo) GetActiveGatewayConfig(ctx context.Context, paymentAccountID uuid.UUID, providerCode string) (*payment_engine_models.PaymentGatewayConfig, error) {
	db, ok := dbtx.ExtractTx(ctx)
	if !ok {
		db = r.db
	}

	var config payment_engine_models.PaymentGatewayConfig
	err := db.WithContext(ctx).
		Where("payment_account_id = ? AND provider_code = ? AND is_active = ?", paymentAccountID, providerCode, true).
		First(&config).Error
	if err != nil {
		return nil, fmt.Errorf("failed to find active config for account %s and provider %s: %w", paymentAccountID, providerCode, err)
	}

	return &config, nil
}

func (r *paymentRepo) ListGatewayConfigsByAccount(ctx context.Context, paymentAccountID uuid.UUID) ([]*payment_engine_models.PaymentGatewayConfig, error) {
	db, ok := dbtx.ExtractTx(ctx)
	if !ok {
		db = r.db
	}

	var configs []*payment_engine_models.PaymentGatewayConfig
	err := db.WithContext(ctx).Where("payment_account_id = ?", paymentAccountID).Find(&configs).Error
	if err != nil {
		return nil, fmt.Errorf("failed to list configs for account %s: %w", paymentAccountID, err)
	}

	return configs, nil
}

func (r *paymentRepo) UpdateGatewayConfig(ctx context.Context, id uuid.UUID, configEncrypted string) error {
	db, ok := dbtx.ExtractTx(ctx)
	if !ok {
		db = r.db
	}

	result := db.WithContext(ctx).Model(&payment_engine_models.PaymentGatewayConfig{}).
		Where("id = ?", id).
		Updates(map[string]interface{}{
			"config_encrypted": configEncrypted,
		})

	if result.Error != nil {
		return fmt.Errorf("failed to update gateway config %s: %w", id, result.Error)
	}

	return nil
}

func (r *paymentRepo) ToggleGatewayConfig(ctx context.Context, id uuid.UUID, isActive bool) error {
	db, ok := dbtx.ExtractTx(ctx)
	if !ok {
		db = r.db
	}

	err := db.WithContext(ctx).Model(&payment_engine_models.PaymentGatewayConfig{}).
		Where("id = ?", id).
		Update("is_active", isActive).Error
	if err != nil {
		return fmt.Errorf("failed to toggle config status for %s: %w", id, err)
	}

	return nil
}

func (r *paymentRepo) DeleteGatewayConfig(ctx context.Context, id uuid.UUID) error {
	db, ok := dbtx.ExtractTx(ctx)
	if !ok {
		db = r.db
	}

	// Using Unscoped() if you want a hard delete,
	// otherwise GORM performs a soft delete if DeletedAt exists in the model.
	err := db.WithContext(ctx).Delete(&payment_engine_models.PaymentGatewayConfig{}, "id = ?", id).Error
	if err != nil {
		return fmt.Errorf("failed to delete gateway config %s: %w", id, err)
	}

	return nil
}
