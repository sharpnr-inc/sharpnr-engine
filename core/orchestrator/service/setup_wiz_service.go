package orchestrator_service

import (
	"context"
	orchestrator_dtos "sharpnr-engine/core/orchestrator/dtos"
	org_setup_wizard_contracts "sharpnr-engine/core/org_setup_wizard/contracts"

	"github.com/google/uuid"
)

// These all service uses the Service from the SetupWizard
// These are service we have to implement the contracts

func (s *orchestratorService) AddOrganizationHQAddress(
	ctx context.Context,
	orgID uuid.UUID,
	req *orchestrator_dtos.SetupOrganizationHQAddressReqDTO,
) error {

	// Step 1 --> This is the mapping process, orchestrator implements the contracts of
	// setup wizard

	input := org_setup_wizard_contracts.StepOneOrganizationHQ{
		OrgID:        orgID,
		Address_line: req.AddressLine,
		Country:      req.Country,
		City:         req.City,
		State:        req.State,
		Pincode:      req.Pincode,
		Location: org_setup_wizard_contracts.StepOneOrganizationHQGeoLocation{
			Longitude: req.Longitude,
			Latitude:  req.Latitude,
			Accuracy:  req.AccuracyM,
		},
	}

	if err := s.orgSetupWizService.Step_one_AddOrgHqAddrAndGeoLocation(ctx, input); err != nil {
		return err
	}

	return nil
}

func (s *orchestratorService) StepTwoProperty(
	ctx context.Context,
	orgID uuid.UUID,
	req *orchestrator_dtos.SetupOrgPropertyReqDTO,
) error {

	s.orgSetupWizService.Step_two_AddProperty(ctx)
	return nil
}

func (s *orchestratorService) StepThreeLegalEntity(
	ctx context.Context,
	orgID uuid.UUID,
	req *orchestrator_dtos.SetupLegalEntityReqDTO,
) error {
	s.orgSetupWizService.Step_three_AddLegalEntity(ctx)
	return nil
}

func (s *orchestratorService) SetUpWizProgress(ctx context.Context, orgID uuid.UUID) error {
	s.orgSetupWizService.Progress(ctx, orgID)
	return nil
}
