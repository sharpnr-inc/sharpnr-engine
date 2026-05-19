package org_setup_wizard_services

import (
	"context"
	org_setup_wizards_constants "sharpnr-engine/core/org_setup_wizard/constants"
	org_setup_wizard_contracts "sharpnr-engine/core/org_setup_wizard/contracts"

	"github.com/google/uuid"
)

// This si version V1 interface
// to Upgrade it Just add another interface making the
type OrgSetupWizardV1_ServiceInterface interface {
	Step_two_AddProperty(ctx context.Context, orgID, emplID uuid.UUID, input org_setup_wizard_contracts.PropertySetupInput) error

	AddStepProgress(
		ctx context.Context,
		orgID uuid.UUID,
		step org_setup_wizards_constants.SetupStepType,
		stepStatus org_setup_wizards_constants.SetupStatus,
	) error

	StepVerication(ctx context.Context, orgID uuid.UUID, emplID uuid.UUID, stepID uuid.UUID) (*org_setup_wizards_constants.SetupStatus, error)

	UpdateSuperAdminDetailsStep(
		ctx context.Context,
		orgID uuid.UUID, emplID uuid.UUID,
		input org_setup_wizard_contracts.SuperAdminDetailsInput,
	) error

	UpdateOrgDetailsStep(
		ctx context.Context,
		orgID uuid.UUID, emplID uuid.UUID,
		input org_setup_wizard_contracts.OrganizationDetailsSetupInput,
	) error

	Progress(ctx context.Context, orgID uuid.UUID) (org_setup_wizard_contracts.SetupProgress, error)
}
