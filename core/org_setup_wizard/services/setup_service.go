package org_setup_wizard_services

import (
	"context"
	org_setup_wizard_models "sharpnr-engine/core/org_setup_wizard/models"
	"time"

	"github.com/google/uuid"
)

func (r *orgSetupWizardService) CreateAndStartSetupStep(
	ctx context.Context,
	orgID uuid.UUID,
	currentStep string,
	status string,
) error {

	setUpStep := &org_setup_wizard_models.OrganizationSetup{
		ID:             uuid.New(),
		OrganizationID: orgID,
		Status:         status,
		CurrentStep:    currentStep,
		Version:        1,
		StartedAt:      time.Now(),
		CreatedAt:      time.Now(),
	}

	if err := r.orgSetWizRepo.Create(ctx, setUpStep); err != nil {
		return err
	}

	return nil
}

// // this method with fetch the current steps
// func (r *orgSetupWizardService) GetSetupProgress(
// 	ctx context.Context,
// 	orgID uuid.UUID,
// ) error

// func (r *orgSetupWizardService) UpdateCurrentStepStatus(
// 	ctx context.Context,

// ) error
