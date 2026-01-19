package org_setup_wizard_services

import (
	"context"
	org_setup_wizard_contracts "sharpnr-engine/core/org_setup_wizard/contracts"
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

// this method with fetch the current steps
func (r *orgSetupWizardService) GetSetupProgress(
	ctx context.Context,
	orgID uuid.UUID,
) (*org_setup_wizard_contracts.SetupProgressOutput, error) {

	steps, err := r.orgSetWizRepo.GetAllSetupStepByOrgID(ctx, orgID)
	if err != nil {
		return nil, err
	}

	// 👇 No steps = onboarding not started
	if len(steps) == 0 {
		return &org_setup_wizard_contracts.SetupProgressOutput{
			Status: "NOT_STARTED",
		}, nil
	}

	return resolveProgress(steps), nil

}

func resolveProgress(
	steps []org_setup_wizard_models.OrganizationSetup,
) *org_setup_wizard_contracts.SetupProgressOutput {

	var (
		completedSteps []string
		currentStep    string
		currentStepID  uuid.UUID
		completedCount int
	)

	for _, step := range steps {

		switch step.Status {
		case "COMPLETED":
			completedSteps = append(completedSteps, step.CurrentStep)
			completedCount++

		case "IN_PROGRESS":
			currentStep = step.CurrentStep
			currentStepID = step.ID
		}
	}

	// Determine overall status
	status := "IN_PROGRESS"
	if completedCount == len(steps) {
		status = "COMPLETED"
	}

	return &org_setup_wizard_contracts.SetupProgressOutput{
		StepID:         currentStepID,
		Status:         status,
		CurrentStep:    currentStep,
		CompletedSteps: completedSteps,
		TotalSteps:     len(steps),
		CompletedCount: completedCount,
	}
}

func (r *orgSetupWizardService) UpdateCurrentStepStatus(
	ctx context.Context,
	orgID uuid.UUID,
	stepID uuid.UUID,
	status string,
) error
