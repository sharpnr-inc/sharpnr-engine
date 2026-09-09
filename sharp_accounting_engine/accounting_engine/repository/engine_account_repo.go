package accounting_repo

import (
	"context"
	"errors"

	dbtx "sharpnr-engine/pkg/dbTx"
	accountingEngineErrs "sharpnr-engine/sharp_accounting_engine/accounting_engine/errs"
	accounting_models "sharpnr-engine/sharp_accounting_engine/accounting_engine/models"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

// InitializeAccountingEngineAccount implements [EngineAccountRepositoryInterface].
//
// The caller mints the id (uuid.New()) because it needs it for the settings row
// written in the same transaction — the two can never disagree that way.
func (r *accountingRepository) InitializeAccountingEngineAccount(ctx context.Context, acc accounting_models.OrgAccountineEngineAccount) error {
	db, ok := dbtx.ExtractTx(ctx)
	if !ok {
		db = r.db
	}
	return db.WithContext(ctx).Create(&acc).Error
}

/* ---------------------------------- <***> --------------------------------- */

// GetAccountingEngineAccountByOrgID implements [EngineAccountRepositoryInterface].
//
// "Never started" is a distinct answer rather than a bare gorm miss, because
// every accounting-owned flow branches on exactly that.
func (r *accountingRepository) GetAccountingEngineAccountByOrgID(ctx context.Context, orgID uuid.UUID) (*accounting_models.OrgAccountineEngineAccount, error) {
	db, ok := dbtx.ExtractTx(ctx)
	if !ok {
		db = r.db
	}

	var account accounting_models.OrgAccountineEngineAccount

	// First, not Find: a miss has to be an error here so the caller can tell
	// "never started" from "started", rather than reading back a zero struct.
	err := db.WithContext(ctx).
		Where("organization_id = ?", orgID).
		First(&account).
		Error

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, accountingEngineErrs.ErrAccountingServiceNotInitialized
		}
		return nil, err
	}

	return &account, nil
}

/* ---------------------------------- <***> --------------------------------- */

// CreateSettingForAccountingEngineAccount implements [EngineAccountRepositoryInterface].
func (r *accountingRepository) CreateSettingForAccountingEngineAccount(ctx context.Context, accSetting accounting_models.AccountingEngineAccountSetting) error {
	db, ok := dbtx.ExtractTx(ctx)
	if !ok {
		db = r.db
	}
	return db.WithContext(ctx).Create(&accSetting).Error
}

/* ---------------------------------- <***> --------------------------------- */

// AccountingEngineAccountSwitch implements [EngineAccountRepositoryInterface].
//
// The account id is ANDed with the org rather than trusted on its own: an id
// belonging to another tenant affects no rows and reads back as not
// initialized.
func (r *accountingRepository) AccountingEngineAccountSwitch(ctx context.Context, orgID, orgAccountingAccountID uuid.UUID, isActive bool) error {
	db, ok := dbtx.ExtractTx(ctx)
	if !ok {
		db = r.db
	}

	result := db.WithContext(ctx).
		Model(&accounting_models.OrgAccountineEngineAccount{}).
		Where("organization_id = ? AND id = ?", orgID, orgAccountingAccountID).
		Update("is_active", isActive)

	if result.Error != nil {
		return result.Error
	}

	if result.RowsAffected == 0 {
		return accountingEngineErrs.ErrAccountingServiceNotInitialized
	}

	return nil
}
