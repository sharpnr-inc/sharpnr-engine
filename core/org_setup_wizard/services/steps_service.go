package org_setup_wizard_services

import (
	"context"
	"errors"
	org_setup_wizards_constants "sharpnr-engine/core/org_setup_wizard/constants"
	orgSetUpWizardErrs "sharpnr-engine/core/org_setup_wizard/errs"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

func (s *orgSetupWizardService) StepVerication(ctx context.Context, orgID uuid.UUID, emplID uuid.UUID, stepID uuid.UUID) (*org_setup_wizards_constants.SetupStatus, error) {

	step, err := s.orgSetWizRepo.GetStepByID(ctx, orgID, stepID)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, orgSetUpWizardErrs.StepDoesntExistErr
		}
		return nil, err
	}

	if step.OrganizationID != orgID {
		return nil, orgSetUpWizardErrs.StepDoesntExistErr
	}

	// we can also check if the employee is authorized to access this step, but for now we will just return the step status.
	return &step.Status, nil
}
