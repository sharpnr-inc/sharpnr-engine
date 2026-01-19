package org_setup_wizard_repos

import (
	"context"
	org_setup_wizard_models "sharpnr-engine/core/org_setup_wizard/models"

	"github.com/google/uuid"
)

// This interface

type OrganizationSetupWizardRepositoryInterface interface {

	// This method Creates the single entry with step status, current step name
	Create(ctx context.Context, step *org_setup_wizard_models.OrganizationSetup) error

	GetAllSetupStepByOrgID(ctx context.Context, orgID uuid.UUID) ([]org_setup_wizard_models.OrganizationSetup, error)

	UpdateCurrentStepStatus(ctx context.Context, stepStatus string, OrgID uuid.UUID, ID uuid.UUID) error
}
