package org_setup_wizard_services

import (
	"context"
	legal_entity_contracts "sharpnr-engine/core/org_legal_entity/contracts"
	org_setup_wizard_contracts "sharpnr-engine/core/org_setup_wizard/contracts"
)

func (s *orgSetupWizardService) Step_three_AddLegalEntity(
	ctx context.Context,
	step_three_input org_setup_wizard_contracts.SetupLegalEntityInput,
) error {

	inputData := legal_entity_contracts.NewLegalEntityInput{
		OrgID:       step_three_input.OrgID,
		LegalName:   step_three_input.LegalName,
		Type:        step_three_input.Type,
		AddressLine: step_three_input.AddressLine,
		City:        step_three_input.City,
		State:       step_three_input.State,
		Country:     step_three_input.Country,
		Pincode:     step_three_input.Pincode,
	}

	if _, err := s.orgLegalEntitySvc.CreateNewLegalEntity(ctx, inputData); err != nil {
		return err
	}

	if err := s.orgSetWizRepo.UpdateCurrentStepStatus(ctx, "COMPLETED", step_three_input.OrgID, step_three_input.StepID); err != nil {
		return err
	}
	return nil
}
