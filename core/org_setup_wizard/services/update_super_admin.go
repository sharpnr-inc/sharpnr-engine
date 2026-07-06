package org_setup_wizard_services

import (
	"context"
	"log"
	coreEmployeeContracts "sharpnr-engine/core/employees/contracts"
	org_setup_wizards_constants "sharpnr-engine/core/org_setup_wizard/constants"
	org_setup_wizard_contracts "sharpnr-engine/core/org_setup_wizard/contracts"

	"github.com/google/uuid"
)

func (s *orgSetupWizardService) UpdateSuperAdminDetailsStep(
	ctx context.Context,
	orgID uuid.UUID,
	emplID uuid.UUID,
	input org_setup_wizard_contracts.SuperAdminDetailsInput,
) error {

	log.Printf(
		"[ORG_SETUP] Starting UpdateSuperAdminDetailsStep | org_id=%s employee_id=%s step_id=%s",
		orgID,
		emplID,
		input.StepID,
	)

	stepStatus, err := s.StepVerication(ctx, orgID, emplID, input.StepID)
	if err != nil {
		log.Printf(
			"[ORG_SETUP] Step verification failed | org_id=%s employee_id=%s step_id=%s error=%v",
			orgID,
			emplID,
			input.StepID,
			err,
		)
		return err
	}

	log.Printf(
		"[ORG_SETUP] Step verified successfully | org_id=%s step_id=%s status=%s",
		orgID,
		input.StepID,
		*stepStatus,
	)

	if *stepStatus != org_setup_wizards_constants.SetupStatusInProgress {
		log.Printf(
			"[ORG_SETUP] Invalid step status | org_id=%s expected=%s actual=%s",
			orgID,
			org_setup_wizards_constants.SetupStatusInProgress,
			*stepStatus,
		)

		return ErrorInvalidStepStatus
	}

	log.Printf(
		"[ORG_SETUP] Updating super admin details | org_id=%s employee_id=%s",
		orgID,
		emplID,
	)

	if err := s.coreEmployeeService.UpdateSuperAdminEmployee(
		ctx,
		orgID,
		emplID,
		&coreEmployeeContracts.UpdateSuperAdminEmployeeInput{
			OrganizationID: orgID,
			EmployeeID:     &emplID,
			FirstName:      &input.FirstName,
			LastName:       &input.FirstName,
			SecondaryEmail: &input.SecondaryEmail,
		},
	); err != nil {

		log.Printf(
			"[ORG_SETUP] Failed to update super admin details | org_id=%s employee_id=%s error=%v",
			orgID,
			emplID,
			err,
		)

		return err
	}

	log.Printf(
		"[ORG_SETUP] Super admin details updated successfully | org_id=%s",
		orgID,
	)

	log.Printf(
		"[ORG_SETUP] Marking current step as completed | org_id=%s step_id=%s",
		orgID,
		input.StepID,
	)

	if err := s.orgSetWizRepo.UpdateCurrentStepStatus(
		ctx,
		org_setup_wizards_constants.SetupStatusCompleted,
		orgID,
		input.StepID,
	); err != nil {

		log.Printf(
			"[ORG_SETUP] Failed to update current step status | org_id=%s step_id=%s error=%v",
			orgID,
			input.StepID,
			err,
		)

		return err
	}

	log.Printf(
		"[ORG_SETUP] Current step marked as completed | org_id=%s step_id=%s",
		orgID,
		input.StepID,
	)

	log.Printf(
		"[ORG_SETUP] Moving next step to in-progress | org_id=%s next_step=%s",
		orgID,
		org_setup_wizards_constants.SetupStepOrgDetailsUpdate,
	)

	if err := s.AddStepProgress(
		ctx,
		orgID,
		org_setup_wizards_constants.SetupStepOrgDetailsUpdate,
		org_setup_wizards_constants.SetupStatusInProgress,
	); err != nil {

		log.Printf(
			"[ORG_SETUP] Failed to initialize next step | org_id=%s next_step=%s error=%v",
			orgID,
			org_setup_wizards_constants.SetupStepOrgDetailsUpdate,
			err,
		)

		return err
	}

	log.Printf(
		"[ORG_SETUP] UpdateSuperAdminDetailsStep completed successfully | org_id=%s",
		orgID,
	)

	return nil
}

/* --------------------------------- <****> --------------------------------- */

// This is the step in V1 this step simply updates Organization name
// as we go further we can add the further details
func (s *orgSetupWizardService) UpdateOrgDetailsStep(
	ctx context.Context,
	orgID uuid.UUID, emplID uuid.UUID,
	input org_setup_wizard_contracts.OrganizationDetailsSetupInput,
) error {

	// 1. Verification of the Step if it belongs to same organization

	stepstatus, err := s.StepVerication(ctx, orgID, emplID, input.StepID)
	if err != nil {
		return err
	}

	if *stepstatus != org_setup_wizards_constants.SetupStatusInProgress {
		return ErrorInvalidStepStatus
	}

	// 2. Update the Organization name
	if err := s.orgService.UpdateOrganizationName(ctx, orgID, emplID, input.Name); err != nil {
		return err
	}

	if err := s.orgSetWizRepo.UpdateCurrentStepStatus(ctx,
		org_setup_wizards_constants.SetupStatusCompleted,
		orgID, input.StepID); err != nil {
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

/* --------------------------------- <****> --------------------------------- */
