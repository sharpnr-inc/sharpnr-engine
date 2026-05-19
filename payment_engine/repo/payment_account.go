package payment_engine_repo

import (
	"context"
	"fmt"
	payment_engine_models "sharpnr-engine/payment_engine/models"
	dbtx "sharpnr-engine/pkg/dbTx"

	"github.com/google/uuid"
)

func (r *paymentRepo) CreatePaymentAccount(ctx context.Context, account *payment_engine_models.PaymentAccount) error {
	// 1. Extract the DB from context (Transaction or Base DB)
	db, ok := dbtx.ExtractTx(ctx)
	if !ok {
		db = r.db
	}

	// 2. Use WithContext to ensure trace propagation and timeouts are respected
	// 3. Perform the Create operation
	err := db.WithContext(ctx).Create(account).Error
	if err != nil {
		return fmt.Errorf("could not create payment account: %w", err)
	}

	return nil
}

/* ---------------------------------- <***> --------------------------------- */

func (r *paymentRepo) GetPaymentAccountByID(ctx context.Context, id uuid.UUID) (*payment_engine_models.PaymentAccount, error) {
	db, ok := dbtx.ExtractTx(ctx)
	if !ok {
		db = r.db
	}

	var account payment_engine_models.PaymentAccount
	err := db.WithContext(ctx).First(&account, "id = ?", id).Error
	if err != nil {
		return nil, fmt.Errorf("failed to get payment account by id %s: %w", id, err)
	}

	return &account, nil
}

/* ---------------------------------- <***> --------------------------------- */

func (r *paymentRepo) GetPaymentAccountByLegalEntity(ctx context.Context, tenantID, legalEntityID uuid.UUID) (*payment_engine_models.PaymentAccount, error) {
	db, ok := dbtx.ExtractTx(ctx)
	if !ok {
		db = r.db
	}

	var account payment_engine_models.PaymentAccount
	err := db.WithContext(ctx).Where("tenant_id = ? AND legal_entity_id = ?", tenantID, legalEntityID).First(&account).Error
	if err != nil {
		return nil, fmt.Errorf("failed to find account for legal entity %s: %w", legalEntityID, err)
	}

	return &account, nil
}

/* ---------------------------------- <***> --------------------------------- */

func (r *paymentRepo) ListPaymentAccounts(ctx context.Context, tenantID uuid.UUID) ([]*payment_engine_models.PaymentAccount, error) {
	db, ok := dbtx.ExtractTx(ctx)
	if !ok {
		db = r.db
	}

	var accounts []*payment_engine_models.PaymentAccount
	err := db.WithContext(ctx).Where("tenant_id = ?", tenantID).Find(&accounts).Error
	if err != nil {
		return nil, fmt.Errorf("failed to list payment accounts for tenant %s: %w", tenantID, err)
	}

	return accounts, nil
}

/* ---------------------------------- <***> --------------------------------- */

func (r *paymentRepo) UpdatePaymentAccountStatus(ctx context.Context, id uuid.UUID, isActive bool) error {
	db, ok := dbtx.ExtractTx(ctx)
	if !ok {
		db = r.db
	}

	result := db.WithContext(ctx).Model(&payment_engine_models.PaymentAccount{}).
		Where("id = ?", id).
		Update("is_active", isActive)

	if result.Error != nil {
		return fmt.Errorf("failed to update status for account %s: %w", id, result.Error)
	}

	if result.RowsAffected == 0 {
		return fmt.Errorf("no account found with id %s to update", id)
	}

	return nil
}

/* ---------------------------------- <***> --------------------------------- */
