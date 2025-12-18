package accounting_repo

import (
	"context"
	"errors"
	"sharpnr-engine/modules/accounting/accounting_models"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

// CreateAccount implements AccountingRepositoryInterface.
func (r *accountingRepository) CreateAccount(ctx context.Context, account *accounting_models.Account) error {
	return r.db.WithContext(ctx).Create(account).Error
}

// GetAccountByID implements AccountingRepositoryInterface.
func (r *accountingRepository) GetAccountByID(ctx context.Context, accountID uuid.UUID) (*accounting_models.Account, error) {
	var account accounting_models.Account

	err := r.db.WithContext(ctx).
		Where("id = ?", accountID).
		First(&account).
		Error

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}

	return &account, nil
}

// GetAccountsByGroupID implements AccountingRepositoryInterface.
func (r *accountingRepository) GetAccountsByGroupID(ctx context.Context, groupID uuid.UUID) ([]accounting_models.Account, error) {
	var accounts []accounting_models.Account

	err := r.db.WithContext(ctx).
		Where("account_group_id = ?", groupID).
		Order("created_at ASC").
		Find(&accounts).
		Error

	if err != nil {
		return nil, err
	}

	return accounts, nil
}
