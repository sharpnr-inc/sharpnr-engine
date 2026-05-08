package organization_services

import (
	"context"
	"errors"
	dbtx "sharpnr-engine/core/dbTx"
	organization_constants "sharpnr-engine/core/organization/constants"
	organization_contracts "sharpnr-engine/core/organization/contracts"
	"sharpnr-engine/core/organization/models"
	mailer "sharpnr-engine/sharp_mailer"
	"sharpnr-engine/sharp_mailer_engine"
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

// AddHQ implements OrganizationInternalServiceInterface.
func (r *organizationService) AddHQ(ctx context.Context, input organization_contracts.OrganizationHQ) error {

	verifiedOrg, err := r.orgMainRepo.GetOrganizationByID(ctx, input.OrgID)
	if err != nil {
		return err
	}

	office := &models.OrganizationOffice{
		ID:             uuid.New(),
		OrganizationID: verifiedOrg.ID,
		Code:           "HQ",
		Name:           "Default Org Office",
		IsHQ:           true,
		IsActive:       false,
		CreatedAt:      time.Now(),
		UpdatedAt:      time.Now(),
	}

	if err := r.orgMainRepo.CreateOrgOffice(ctx, office); err != nil {
		return err
	}

	officeAddr := &models.OrganizationOfficeAddress{
		ID:          uuid.New(),
		OfficeID:    office.ID,
		Type:        "",
		AddressLine: input.Address_line,
		City:        input.City,
		State:       input.State,
		Pincode:     input.Pincode,
		Country:     input.Country,
		IsPrimary:   true,
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}

	if err := r.orgMainRepo.CreateOrgOfficeAddress(ctx, officeAddr); err != nil {
		return err
	}

	r.mailerEngine.Enqueue(mailer.Notification, sharp_mailer_engine.NewNotificationJob(
		verifiedOrg.RootEmail, "HQ Notification", sharp_mailer_engine.NotificationData{
			Title:      "",
			Subtitle:   "",
			Message:    "",
			ActionURL:  "",
			ActionText: "",
			FooterNote: "What the heck",
		},
	))
	return nil
}

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

// GetOrgHQ implements OrganizationPublicServiceInterface.
func (r *organizationService) GetOrgHQ() {
	panic("unimplemented")
}
