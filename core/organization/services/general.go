package organization_services

import (
	"context"
	"errors"
	dbtx "sharpnr-engine/core/dbTx"
	organization_constants "sharpnr-engine/core/organization/constants"
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

/* ---------------------------------- <***> --------------------------------- */

func (r *organizationService) GetStatus(ctx context.Context, orgID uuid.UUID) (organization_constants.OrgSetupStatusType, error) {
	org, err := r.orgMainRepo.GetOrganizationByID(ctx, orgID)
	if err != nil {
		return "", err
	}

	return org.Status, nil
}

/* ---------------------------------- <***> --------------------------------- */

func (r *organizationService) UpdateStatus(ctx context.Context, orgID uuid.UUID, status organization_constants.OrgSetupStatusType) error {
	org, err := r.orgMainRepo.GetOrganizationByID(ctx, orgID)
	if err != nil {
		return err
	}

	org.Status = status
	org.UpdatedAt = time.Now()

	if err := r.orgMainRepo.UpdateOrganizationStatus(ctx, org.ID, status); err != nil {
		return err
	}

	return nil
}

/* ---------------------------------- <***> --------------------------------- */

// This is the method which Verify the oranisation
func (r *organizationService) VerifyOrgByID(ctx context.Context, orgID uuid.UUID) (bool, error) {

	err := r.txManager.WithTx(ctx, func(tx *gorm.DB) error {

		dbTx := dbtx.InjectTx(ctx, tx)

		org, err := r.orgMainRepo.GetOrganizationByID(dbTx, orgID)
		if err != nil {
			return err
		}
		if org.ID != orgID {
			return errors.New("org do not exist")
		}
		return nil
	})

	if err != nil {
		return false, err
	}
	return true, nil
}
