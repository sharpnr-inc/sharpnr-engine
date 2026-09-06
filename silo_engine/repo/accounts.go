package siloEngineRepo

import (
	"context"
	"errors"

	siloEngineErrs "sharpnr-engine/silo_engine/errs"
	siloEngineModels "sharpnr-engine/silo_engine/models"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

// InitializeSiloAccount implements [SiloEngineRepoInterface].
//
// The caller mints the id (uuid.New()) because it needs it for the settings row
// written in the same transaction — the two can never disagree that way.
func (s *siloRepo) InitializeSiloAccount(ctx context.Context, acc siloEngineModels.OrgSiloAccount) error {
	return s.conn(ctx).Create(&acc).Error
}

/* ---------------------------------- <***> --------------------------------- */

// GetSiloAccountByOrgID implements [SiloEngineRepoInterface].
//
// "Never started" is a distinct answer rather than a bare gorm miss, because
// every other Silo flow branches on exactly that.
func (s *siloRepo) GetSiloAccountByOrgID(ctx context.Context, orgID uuid.UUID) (*siloEngineModels.OrgSiloAccount, error) {
	var account siloEngineModels.OrgSiloAccount

	err := s.conn(ctx).
		Where("organization_id = ?", orgID).
		First(&account).
		Error

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, siloEngineErrs.ErrSiloServiceNotInitialized
		}
		return nil, err
	}

	return &account, nil
}

/* ---------------------------------- <***> --------------------------------- */

// CreateSettingForSiloAccount implements [SiloEngineRepoInterface].
func (s *siloRepo) CreateSettingForSiloAccount(ctx context.Context, accSetting siloEngineModels.SiloAccountSetting) error {
	return s.conn(ctx).Create(&accSetting).Error
}

/* ---------------------------------- <***> --------------------------------- */

// SiloAccountSwitch implements [SiloEngineRepoInterface].
//
// The account id is ANDed with the org rather than trusted on its own: an id
// belonging to another tenant affects no rows and reads back as not
// initialized.
func (s *siloRepo) SiloAccountSwitch(ctx context.Context, orgID, orgSiloAccountID uuid.UUID, isActive bool) error {
	result := s.conn(ctx).
		Model(&siloEngineModels.OrgSiloAccount{}).
		Where("organization_id = ? AND id = ?", orgID, orgSiloAccountID).
		Update("is_active", isActive)

	if result.Error != nil {
		return result.Error
	}

	if result.RowsAffected == 0 {
		return siloEngineErrs.ErrSiloServiceNotInitialized
	}

	return nil
}
