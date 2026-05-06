package org_setup_wizard_repos

import (
	"context"
	org_setup_wizards_constants "sharpnr-engine/core/org_setup_wizard/constants"
	org_setup_wizard_models "sharpnr-engine/core/org_setup_wizard/models"

	"github.com/google/uuid"
)

type OrganizationSetupWizardRepositoryInterface interface {
	Create(ctx context.Context, step *org_setup_wizard_models.OrganizationSetup) error
	GetAllSetupStepByOrgID(ctx context.Context, orgID uuid.UUID) ([]*org_setup_wizard_models.OrganizationSetup, error)

	UpdateCurrentStepStatus(
		ctx context.Context,
		stepStatus org_setup_wizards_constants.SetupStatus,
		OrgID uuid.UUID, stepID uuid.UUID,
	) error
	
	GetStepByID(ctx context.Context, orgID uuid.UUID, stepID uuid.UUID) (*org_setup_wizard_models.OrganizationSetup, error)
}
