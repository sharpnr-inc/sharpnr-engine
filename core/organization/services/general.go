package organization_services

import (
	"context"
	"errors"
	organization_contracts "sharpnr-engine/core/organization/contracts"
	"sharpnr-engine/core/organization/models"
	"sharpnr-engine/sharp_mailer_engine"
	"time"

	"github.com/google/uuid"
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

	notificationMailJob := map[string]interface{}{
		"Title":      "Kya Baat hai!! Naya Office",
		"Subtitle":   "Andha Paisa",
		"FooterNote": "Party kab de raha hai",
	}
	r.mailerEngine.Queues[sharp_mailer_engine.Notification] <- sharp_mailer_engine.EmailJob{
		To:       verifiedOrg.RootEmail,
		Subject:  "HQ Updates",
		Template: sharp_mailer_engine.TmplNotification,
		Data:     notificationMailJob,
	}

	return nil
}

func (r *organizationService) GetStatus(ctx context.Context, orgID uuid.UUID) (string, error) {
	org, err := r.orgMainRepo.GetOrganizationByID(ctx, orgID)
	if err != nil {
		return "", err
	}

	return org.Status, nil
}

func (r *organizationService) VerifyOrgByID(ctx context.Context, orgID uuid.UUID) (bool, error) {
	org, err := r.orgMainRepo.GetOrganizationByID(ctx, orgID)
	if err != nil {
		return false, err
	}
	if org.ID != orgID {
		return false, errors.New("org do not exist")
	}

	return true, nil
}
