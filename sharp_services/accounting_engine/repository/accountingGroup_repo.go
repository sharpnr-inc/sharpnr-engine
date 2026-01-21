package accounting_repo

import (
	"context"
	"errors"
	accounting_models "sharpnr-engine/sharp_services/accounting_engine/models"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

// CreateAccountGroup implements AccountingRepositoryInterface.
func (r *accountingRepository) CreateAccountGroup(ctx context.Context, group *accounting_models.AccountGroup) error {
	return r.db.WithContext(ctx).Create(group).Error
}

// GetAccountGroupByID implements AccountingRepositoryInterface.
func (r *accountingRepository) GetAccountGroupByID(ctx context.Context, groupID uuid.UUID) (*accounting_models.AccountGroup, error) {
	var group accounting_models.AccountGroup

	err := r.db.WithContext(ctx).
		Where("id = ?", groupID).
		First(&group).
		Error

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}

	return &group, nil
}

// GetAccountGroupsByBookID implements AccountingRepositoryInterface.
func (r *accountingRepository) GetAccountGroupsByBookID(ctx context.Context, bookID uuid.UUID) ([]accounting_models.AccountGroup, error) {
	var groups []accounting_models.AccountGroup

	err := r.db.WithContext(ctx).
		Where("accounting_book_id = ?", bookID).
		Order("created_at ASC").
		Find(&groups).
		Error

	if err != nil {
		return nil, err
	}

	return groups, nil
}
