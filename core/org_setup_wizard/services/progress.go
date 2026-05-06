package org_setup_wizard_services

import (
	"context"
	org_setup_wizards_constants "sharpnr-engine/core/org_setup_wizard/constants"
	org_setup_wizard_contracts "sharpnr-engine/core/org_setup_wizard/contracts"
	org_setup_wizard_models "sharpnr-engine/core/org_setup_wizard/models"
	org_setup_wizard_resolvers "sharpnr-engine/core/org_setup_wizard/resolvers"
	"time"

	"github.com/google/uuid"
)

func (s *orgSetupWizardService) Progress(ctx context.Context, orgID uuid.UUID) (org_setup_wizard_contracts.SetupProgress, error) {
	steps, err := s.orgSetWizRepo.GetAllSetupStepByOrgID(ctx, orgID)
	if err != nil {
		return org_setup_wizard_contracts.SetupProgress{}, err
	}

	return org_setup_wizard_resolvers.ResolveSetupProgress(steps), nil

}

/* ---------------------------------- <***> --------------------------------- */
// Adds a new step progress entry for the organization in the database. This function is called whenever a new step in the setup wizard is initiated or completed. It records the current step, its status (e.g., IN_PROGRESS, COMPLETED), and timestamps to track the organization's progress through the setup process. This allows us to provide real-time feedback to users about their setup progress and helps in orchestrating the next steps based on the current state of the setup wizard.
func (s *orgSetupWizardService) AddStepProgress(
	ctx context.Context, orgID uuid.UUID,
	step org_setup_wizards_constants.SetupStepType,
	stepStatus org_setup_wizards_constants.SetupStatus,
) error {

	setupStep := &org_setup_wizard_models.OrganizationSetup{
		ID:             uuid.New(),
		OrganizationID: orgID,
		Status:         stepStatus,
		CurrentStep:    step,
		Version:        1,
		StartedAt:      time.Now(),
		CreatedAt:      time.Now(),
		UpdatedAt:      time.Now(),
	}

	if err := s.orgSetWizRepo.Create(ctx, setupStep); err != nil {
		return err
	}

	return nil
}

/* ---------------------------------- <***> --------------------------------- */
