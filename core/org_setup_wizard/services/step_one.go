package org_setup_wizard_services

import (
	"context"
	org_setup_wizard_contracts "sharpnr-engine/core/org_setup_wizard/contracts"
	org_setup_wizard_models "sharpnr-engine/core/org_setup_wizard/models"
	organization_contracts "sharpnr-engine/core/organization/contracts"
	"time"

	"github.com/google/uuid"
)

func (s *orgSetupWizardService) Step_one_AddOrgHqAddrAndGeoLocation(
	ctx context.Context,
	input org_setup_wizard_contracts.StepOneOrganizationHQ,
) error {

	// This is the mapping process here
	// which implement the contract of organization
	// it is not exposed to the Repo. of organization

	data_input := organization_contracts.OrganizationHQ{
		OrgID:        input.OrgID,
		Address_line: input.Address_line,
		Country:      input.Country,
		City:         input.City,
		State:        input.State,
		Pincode:      input.Pincode,
		Location: organization_contracts.OrganizationHQGeoLocation{
			Longitude: input.Location.Latitude,
			Latitude:  input.Location.Latitude,
			Accuracy:  input.Location.Accuracy,
		},
	}

	if err := s.orgService.AddHQ(ctx, data_input); err != nil {
		return err
	}

	if err := s.orgSetWizRepo.UpdateCurrentStepStatus(ctx, "COMPLETE", input.OrgID, input.StepID); err != nil {
		return err
	}

	step := &org_setup_wizard_models.OrganizationSetup{
		ID:             uuid.New(),
		OrganizationID: input.OrgID,
		Status:         "IN_PROGRESS",
		CurrentStep:    "PROPERTY_SETUP",
		Version:        1,
		StartedAt:      time.Now(),
		CreatedAt:      time.Now(),
	}

	if err := s.orgSetWizRepo.Create(ctx, step); err != nil {
		return err
	}

	return nil
}
