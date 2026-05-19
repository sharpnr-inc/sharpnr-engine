package org_setup_wizard_services

import (
	"context"
	"errors"
	org_setup_wizards_constants "sharpnr-engine/core/org_setup_wizard/constants"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

var (
	ErrorInvalidStepStatus = errors.New("invalid step status for this operation")
	ErrorInvalidStep       = errors.New("invalid step for this operation")
	ErrorStepDoesntExist   = errors.New("step does not exist for this organization")
)

func (s *orgSetupWizardService) StepVerication(ctx context.Context, orgID uuid.UUID, emplID uuid.UUID, stepID uuid.UUID) (*org_setup_wizards_constants.SetupStatus, error) {

	step, err := s.orgSetWizRepo.GetStepByID(ctx, orgID, stepID)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, ErrorStepDoesntExist
		}
		return nil, err
	}

	if step.OrganizationID != orgID {
		return nil, ErrorStepDoesntExist
	}

	// we can also check if the employee is authorized to access this step, but for now we will just return the step status.
	return &step.Status, nil
}
