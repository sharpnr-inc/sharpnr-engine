package cortexEngineRepo

import (
	"context"
	"errors"

	cortexEngineErrs "sharpnr-engine/cortex_engine/errs"
	cortexEngineModels "sharpnr-engine/cortex_engine/models"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

// InitializeCortexAccount implements [CortexEngineRepoInterface].
//
// The caller mints the id (uuid.New()) because it needs it for the settings row
// written in the same transaction — the two can never disagree that way.
func (c *cortexRepo) InitializeCortexAccount(ctx context.Context, acc cortexEngineModels.OrgCortexAccount) error {
	return c.conn(ctx).Create(&acc).Error
}

/* ---------------------------------- <***> --------------------------------- */

// GetCortexAccountByOrgID implements [CortexEngineRepoInterface].
//
// "Never started" is a distinct answer rather than a bare gorm miss, because
// every other Cortex flow branches on exactly that.
func (c *cortexRepo) GetCortexAccountByOrgID(ctx context.Context, orgID uuid.UUID) (*cortexEngineModels.OrgCortexAccount, error) {
	var account cortexEngineModels.OrgCortexAccount

	err := c.conn(ctx).
		Where("organization_id = ?", orgID).
		First(&account).
		Error

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, cortexEngineErrs.ErrNoCortexAccountFound
		}
		return nil, err
	}

	return &account, nil
}

/* ---------------------------------- <***> --------------------------------- */

// CreateSettingForCortexAccount implements [CortexEngineRepoInterface].
func (c *cortexRepo) CreateSettingForCortexAccount(ctx context.Context, accSetting cortexEngineModels.CortexAccountSetting) error {
	return c.conn(ctx).Create(&accSetting).Error
}

/* ---------------------------------- <***> --------------------------------- */

// CortexAccountSwitch implements [CortexEngineRepoInterface].
//
// The account id is ANDed with the org rather than trusted on its own: an id
// belonging to another tenant affects no rows and reads back as not found.
func (c *cortexRepo) CortexAccountSwitch(ctx context.Context, orgID, orgCortexAccountID uuid.UUID, isActive bool) error {
	result := c.conn(ctx).
		Model(&cortexEngineModels.OrgCortexAccount{}).
		Where("organization_id = ? AND id = ?", orgID, orgCortexAccountID).
		Update("is_active", isActive)

	if result.Error != nil {
		return result.Error
	}

	if result.RowsAffected == 0 {
		return cortexEngineErrs.ErrNoCortexAccountFound
	}

	return nil
}
