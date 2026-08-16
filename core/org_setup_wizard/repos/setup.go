package org_setup_wizard_repos

import (
	"context"
	"errors"
	org_setup_wizards_constants "sharpnr-engine/core/org_setup_wizard/constants"
	orgSetUpWizardErrs "sharpnr-engine/core/org_setup_wizard/errs"
	org_setup_wizard_models "sharpnr-engine/core/org_setup_wizard/models"
	dbtx "sharpnr-engine/pkg/dbTx"
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

func (d *orgSetupWizardRepo) Create(
	ctx context.Context,
	step *org_setup_wizard_models.OrganizationSetup,
) error {

	if step == nil {
		return orgSetUpWizardErrs.NilSetupStepErr
	}

	conn := d.db
	if tx, ok := dbtx.ExtractTx(ctx); ok {
		conn = tx
	}

	if err := conn.WithContext(ctx).Create(step).Error; err != nil {
		return err
	}

	return nil

}

//--------------------------------------------------------------------//

func (d *orgSetupWizardRepo) GetAllSetupStepByOrgID(
	ctx context.Context,
	orgID uuid.UUID,
) ([]*org_setup_wizard_models.OrganizationSetup, error) {
	var setups []*org_setup_wizard_models.OrganizationSetup

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
	stepStatus org_setup_wizards_constants.SetupStatus,
	orgID uuid.UUID,
	ID uuid.UUID,
) error {
	if stepStatus == "" {
		return orgSetUpWizardErrs.EmptyStepStatusErr
	}

	now := time.Now()

	updates := map[string]interface{}{
		"status":     stepStatus,
		"updated_at": now,
	}

	// A skipped step is treated as complete, so both terminal statuses stamp the
	// completion time.
	if stepStatus == org_setup_wizards_constants.SetupStatusCompleted ||
		stepStatus == org_setup_wizards_constants.SetupStatusSkipped {
		updates["completed_at"] = now
	}

	conn := d.db
	if txConn, ok := dbtx.ExtractTx(ctx); ok {
		conn = txConn
	}

	tx := conn.
		WithContext(ctx).
		Model(&org_setup_wizard_models.OrganizationSetup{}).
		Where("organization_id = ? AND id = ?", orgID, ID).
		Updates(updates)

	if tx.Error != nil {
		return tx.Error
	}

	if tx.RowsAffected == 0 {
		return gorm.ErrRecordNotFound
	}

	return nil

}

//--------------------------------------------------------------------//

func (d *orgSetupWizardRepo) GetStepByID(
	ctx context.Context,
	orgID uuid.UUID,
	stepID uuid.UUID,
) (*org_setup_wizard_models.OrganizationSetup, error) {

	db, ok := dbtx.ExtractTx(ctx)
	if !ok {
		db = d.db
	}

	var step org_setup_wizard_models.OrganizationSetup

	err := db.
		WithContext(ctx).
		Where("id = ? AND organization_id = ?", stepID, orgID).
		First(&step).Error

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, gorm.ErrRecordNotFound
		}
		return nil, err
	}

	return &step, nil
}
