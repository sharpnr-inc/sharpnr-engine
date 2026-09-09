package hrmsEngineRepo

import (
	"context"
	"errors"

	hrmsEngineErrs "sharpnr-engine/hrms_engine/errs"
	hrmsEngineModels "sharpnr-engine/hrms_engine/models"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

// InitializeHrmsAccount implements [HrmsEngineRepoInterface].
//
// The caller mints the id (uuid.New()) because it needs it for the settings row
// written in the same transaction — the two can never disagree that way.
func (d *repository) InitializeHrmsAccount(ctx context.Context, acc hrmsEngineModels.OrgHrmsAccount) error {
	return d.conn(ctx).Create(&acc).Error
}

/* ---------------------------------- <***> --------------------------------- */

// GetHrmsAccountByOrgID implements [HrmsEngineRepoInterface].
//
// "Never started" is a distinct answer rather than a bare gorm miss, because
// every other HRMS flow branches on exactly that.
func (d *repository) GetHrmsAccountByOrgID(ctx context.Context, orgID uuid.UUID) (*hrmsEngineModels.OrgHrmsAccount, error) {
	var account hrmsEngineModels.OrgHrmsAccount

	err := d.conn(ctx).
		Where("organization_id = ?", orgID).
		First(&account).
		Error

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, hrmsEngineErrs.ErrNoHRMSAccountFound
		}
		return nil, err
	}

	return &account, nil
}

/* ---------------------------------- <***> --------------------------------- */

// CreateSettingForHrmsAccount implements [HrmsEngineRepoInterface].
func (d *repository) CreateSettingForHrmsAccount(ctx context.Context, accSetting hrmsEngineModels.HrmsAccountSetting) error {
	return d.conn(ctx).Create(&accSetting).Error
}

/* ---------------------------------- <***> --------------------------------- */

// HrmsAccountSwitch implements [HrmsEngineRepoInterface].
//
// The account id is ANDed with the org rather than trusted on its own: an id
// belonging to another tenant affects no rows and reads back as not found.
func (d *repository) HrmsAccountSwitch(ctx context.Context, orgID, orgHrmsAccountID uuid.UUID, isActive bool) error {
	result := d.conn(ctx).
		Model(&hrmsEngineModels.OrgHrmsAccount{}).
		Where("organization_id = ? AND id = ?", orgID, orgHrmsAccountID).
		Update("is_active", isActive)

	if result.Error != nil {
		return result.Error
	}

	if result.RowsAffected == 0 {
		return hrmsEngineErrs.ErrNoHRMSAccountFound
	}

	return nil
}
