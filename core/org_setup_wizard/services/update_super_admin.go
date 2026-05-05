package org_setup_wizard_services

import (
	"context"
	org_setup_wizards_constants "sharpnr-engine/core/org_setup_wizard/constants"
	org_setup_wizard_contracts "sharpnr-engine/core/org_setup_wizard/contracts"
	sharpnr_identity_auth_dtos "sharpnr-engine/core/sharpnr_identity/dtos"

	"github.com/google/uuid"
)

func (s *orgSetupWizardService) UpdateSuperAdminDetailsStep(ctx context.Context, orgID uuid.UUID, emplID uuid.UUID, input org_setup_wizard_contracts.SuperAdminDetailsInput) error {

	// before updating the super admin details, we need to check if the organization setup wizard is in the correct step and status. If not, we should return an error.
	err := s.sharpIdentity.UpdateSuperAdminDetails(ctx, orgID, emplID, sharpnr_identity_auth_dtos.UpdateSuperAdminDetailsInput{
		OrgID:          input.OrgID,
		FirstName:      input.FirstName,
		LastName:       input.LastName,
		EmployeeID:     input.EmployeeID,
		SecondaryEmail: input.SecondaryEmail,
		PhoneNumber:    input.SecondaryPhone,
	})

	if err != nil {
		return err
	}

	// if the super admin details are updated successfully, we can update the organization setup wizard status to completed for this step and move to the next step.
	if err := s.AddStepProgress(
		ctx, orgID,
		org_setup_wizards_constants.SetupStepSuperAdminProfile,
		org_setup_wizards_constants.SetupStatusCompleted,
	); err != nil {
		return err
	}

	// We make the next step as in progress, so that the frontend can show the next step to the user.
	if err := s.AddStepProgress(
		ctx, orgID,
		org_setup_wizards_constants.SetupStepOperations,
		org_setup_wizards_constants.SetupStatusInProgress,
	); err != nil {
		return err
	}

	return nil
}
