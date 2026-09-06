package folioEngineRepo

import (
	"context"
	"errors"

	folioEngineErrs "sharpnr-engine/folio_engine/errs"
	folioEngineModels "sharpnr-engine/folio_engine/models"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

// InitializeFolioAccount implements [FolioEngineRepoInterface].
//
// The caller mints the id (uuid.New()) because it needs it for the settings row
// written in the same transaction — the two can never disagree that way.
func (f *folioRepo) InitializeFolioAccount(ctx context.Context, acc folioEngineModels.OrgFolioAccount) error {
	return f.conn(ctx).Create(&acc).Error
}

/* ---------------------------------- <***> --------------------------------- */

// GetFolioAccountByOrgID implements [FolioEngineRepoInterface].
//
// "Never started" is a distinct answer rather than a bare gorm miss, because
// every other Folio flow branches on exactly that.
func (f *folioRepo) GetFolioAccountByOrgID(ctx context.Context, orgID uuid.UUID) (*folioEngineModels.OrgFolioAccount, error) {
	var account folioEngineModels.OrgFolioAccount

	err := f.conn(ctx).
		Where("organization_id = ?", orgID).
		First(&account).
		Error

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, folioEngineErrs.ErrFolioServiceNotInitialized
		}
		return nil, err
	}

	return &account, nil
}

/* ---------------------------------- <***> --------------------------------- */

// CreateSettingForFolioAccount implements [FolioEngineRepoInterface].
func (f *folioRepo) CreateSettingForFolioAccount(ctx context.Context, accSetting folioEngineModels.FolioAccountSetting) error {
	return f.conn(ctx).Create(&accSetting).Error
}

/* ---------------------------------- <***> --------------------------------- */

// FolioAccountSwitch implements [FolioEngineRepoInterface].
//
// The account id is ANDed with the org rather than trusted on its own: an id
// belonging to another tenant affects no rows and reads back as not
// initialized.
func (f *folioRepo) FolioAccountSwitch(ctx context.Context, orgID, orgFolioAccountID uuid.UUID, isActive bool) error {
	result := f.conn(ctx).
		Model(&folioEngineModels.OrgFolioAccount{}).
		Where("organization_id = ? AND id = ?", orgID, orgFolioAccountID).
		Update("is_active", isActive)

	if result.Error != nil {
		return result.Error
	}

	if result.RowsAffected == 0 {
		return folioEngineErrs.ErrFolioServiceNotInitialized
	}

	return nil
}
