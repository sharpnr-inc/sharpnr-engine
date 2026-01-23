package org_setup_wizard_services

import (
	"context"
	org_setup_wizard_models "sharpnr-engine/core/org_setup_wizard/models"
	"time"

	"github.com/google/uuid"
)

func (s *orgSetupWizardService) Progress(ctx context.Context, orgID uuid.UUID) error {
	_, err := s.orgSetWizRepo.GetAllSetupStepByOrgID(ctx, orgID)
	if err != nil {
		return err
	}

	return nil
}

func (s *orgSetupWizardService) AddStepProgress(ctx context.Context, orgID uuid.UUID, step string, stepStatus string) error {

	setupStep := &org_setup_wizard_models.OrganizationSetup{
		ID:             uuid.New(),
		OrganizationID: orgID,
		Status:         stepStatus,
		CurrentStep:    step,
		Version:        1,
		StartedAt:      time.Now(),
		CreatedAt:      time.Now(),
	}

	if err := s.orgSetWizRepo.Create(ctx, setupStep); err != nil {
		return err
	}

	return nil
}
