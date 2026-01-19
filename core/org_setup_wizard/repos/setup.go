package org_setup_wizard_repos

import (
	"context"
	"errors"
	org_setup_wizard_models "sharpnr-engine/core/org_setup_wizard/models"
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

func (d *orgSetupWizardRepo) Create(
	ctx context.Context,
	step *org_setup_wizard_models.OrganizationSetup,
) error {

	if step == nil {
		return errors.New("organization setup is nil")
	}

	db := d.db.WithContext(ctx)

	if err := db.Create(step).Error; err != nil {
		return err
	}

	return nil

}

//--------------------------------------------------------------------//

func (d *orgSetupWizardRepo) GetAllSetupStepByOrgID(
	ctx context.Context,
	orgID uuid.UUID,
) ([]org_setup_wizard_models.OrganizationSetup, error) {
	var setups []org_setup_wizard_models.OrganizationSetup

	err := d.db.
		WithContext(ctx).
		Where("organization_id = ?", orgID).
		Order("created_at ASC").
		Find(&setups).Error

	if err != nil {
		return nil, err
	}

	return setups, nil
}

//--------------------------------------------------------------------//

func (d *orgSetupWizardRepo) UpdateCurrentStepStatus(
	ctx context.Context,
	stepStatus string,
	orgID uuid.UUID,
	ID uuid.UUID,
) error {
	if stepStatus == "" {
		return errors.New("step cannot be empty")
	}

	tx := d.db.
		WithContext(ctx).
		Model(&org_setup_wizard_models.OrganizationSetup{}).
		Where("organization_id = ? AND id = ?", orgID, ID).
		Updates(map[string]interface{}{
			"status":     stepStatus,
			"updated_at": time.Now(),
		})

	if tx.Error != nil {
		return tx.Error
	}

	if tx.RowsAffected == 0 {
		return gorm.ErrRecordNotFound
	}

	return nil

}

//--------------------------------------------------------------------//
