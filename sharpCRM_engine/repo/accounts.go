package sharpCrmEngineRepo

import (
	"context"
	"errors"

	sharpCrmEngineErrs "sharpnr-engine/sharpCRM_engine/errs"
	sharpCrmEngineModels "sharpnr-engine/sharpCRM_engine/models"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

// InitializeCrmAccount implements [SharpCrmEngineRepoInterace].
//
// The caller mints the id (uuid.New()) because it needs it for the settings row
// written in the same transaction — the two can never disagree that way.
func (d *repository) InitializeCrmAccount(ctx context.Context, acc sharpCrmEngineModels.OrgSharpCrmAccount) error {
	return d.conn(ctx).Create(&acc).Error
}

/* ---------------------------------- <***> --------------------------------- */

// GetCrmAccountByOrgID implements [SharpCrmEngineRepoInterace].
//
// "Never started" is a distinct answer rather than a bare gorm miss, because
// every other CRM flow branches on exactly that.
func (d *repository) GetCrmAccountByOrgID(ctx context.Context, orgID uuid.UUID) (*sharpCrmEngineModels.OrgSharpCrmAccount, error) {
	var account sharpCrmEngineModels.OrgSharpCrmAccount

	err := d.conn(ctx).
		Where("organization_id = ?", orgID).
		First(&account).
		Error

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, sharpCrmEngineErrs.ErrCrmServiceNotInitialized
		}
		return nil, err
	}

	return &account, nil
}

/* ---------------------------------- <***> --------------------------------- */

// CreateSettingForCrmAccount implements [SharpCrmEngineRepoInterace].
func (d *repository) CreateSettingForCrmAccount(ctx context.Context, acc_setting sharpCrmEngineModels.SharpCrmAccountSetting) error {
	return d.conn(ctx).Create(&acc_setting).Error
}

/* ---------------------------------- <***> --------------------------------- */

// CrmAccountSwitch implements [SharpCrmEngineRepoInterace].
//
// The account id is ANDed with the org rather than trusted on its own: an id
// belonging to another tenant affects no rows and reads back as not
// initialized.
func (d *repository) CrmAccountSwitch(ctx context.Context, orgID, orgCrmAccountID uuid.UUID, isActive bool) error {
	result := d.conn(ctx).
		Model(&sharpCrmEngineModels.OrgSharpCrmAccount{}).
		Where("organization_id = ? AND id = ?", orgID, orgCrmAccountID).
		Updates(map[string]any{
			"is_active":  isActive,
			"updated_at": gorm.Expr("now()"),
		})

	if result.Error != nil {
		return result.Error
	}

	if result.RowsAffected == 0 {
		return sharpCrmEngineErrs.ErrCrmServiceNotInitialized
	}

	return nil
}
