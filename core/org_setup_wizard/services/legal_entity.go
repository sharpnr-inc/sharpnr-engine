package org_setup_wizard_services

import (
	"context"
	legal_entity_contracts "sharpnr-engine/core/legal_entity/contracts"
	org_setup_wizard_contracts "sharpnr-engine/core/org_setup_wizard/contracts"
)

func (s *orgSetupWizardService) Step_three_AddLegalEntity(
	ctx context.Context,
	step_three_input org_setup_wizard_contracts.SetupLegalEntityInput,
) error {

	inputData := legal_entity_contracts.NewLegalEntityRequestDto{
		OrgID:     step_three_input.OrgID,
		LegalName: step_three_input.LegalName,
	}

	if _, err := s.orgLegalEntitySvc.CreateNewLegalEntity(ctx, inputData); err != nil {
		return err
	}

	if err := s.orgSetWizRepo.UpdateCurrentStepStatus(ctx, "COMPLETED", step_three_input.OrgID, step_three_input.StepID); err != nil {
		return err
	}
	return nil
}
