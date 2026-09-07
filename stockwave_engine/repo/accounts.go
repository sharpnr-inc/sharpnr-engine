package stockwaveEngineRepo

import (
	"context"
	"errors"

	stockwaveEngineErrs "sharpnr-engine/stockwave_engine/errs"
	stockwaveEngineModels "sharpnr-engine/stockwave_engine/models"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

// InitializeStockwaveAccount implements [StockwaveEngineRepoInterface].
//
// The caller mints the id (uuid.New()) because it needs it for the settings row
// written in the same transaction — the two can never disagree that way.
func (r *stockwaveRepo) InitializeStockwaveAccount(ctx context.Context, acc stockwaveEngineModels.OrgStockwaveEngineAccounts) error {
	return r.conn(ctx).Create(&acc).Error
}

/* ---------------------------------- <***> --------------------------------- */

// GetStockwaveAccountByOrgID implements [StockwaveEngineRepoInterface].
//
// "Never started" is a distinct answer rather than a bare gorm miss, because
// every other stockwave flow branches on exactly that.
func (r *stockwaveRepo) GetStockwaveAccountByOrgID(ctx context.Context, orgID uuid.UUID) (*stockwaveEngineModels.OrgStockwaveEngineAccounts, error) {
	var account stockwaveEngineModels.OrgStockwaveEngineAccounts

	err := r.conn(ctx).
		Where("organization_id = ?", orgID).
		First(&account).
		Error

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, stockwaveEngineErrs.ErrNoAccountFound
		}
		return nil, err
	}

	return &account, nil
}

/* ---------------------------------- <***> --------------------------------- */

// CreateSettingForStockwaveAccount implements [StockwaveEngineRepoInterface].
func (r *stockwaveRepo) CreateSettingForStockwaveAccount(ctx context.Context, accSetting stockwaveEngineModels.StockwaveEngineAccountSetting) error {
	return r.conn(ctx).Create(&accSetting).Error
}

/* ---------------------------------- <***> --------------------------------- */

// StockwaveAccountSwitch implements [StockwaveEngineRepoInterface].
//
// The account id is ANDed with the org rather than trusted on its own: an id
// belonging to another tenant affects no rows and reads back as not found.
func (r *stockwaveRepo) StockwaveAccountSwitch(ctx context.Context, orgID, orgStockwaveAccountID uuid.UUID, isActive bool) error {
	result := r.conn(ctx).
		Model(&stockwaveEngineModels.OrgStockwaveEngineAccounts{}).
		Where("organization_id = ? AND id = ?", orgID, orgStockwaveAccountID).
		Update("is_active", isActive)

	if result.Error != nil {
		return result.Error
	}

	if result.RowsAffected == 0 {
		return stockwaveEngineErrs.ErrNoAccountFound
	}

	return nil
}
