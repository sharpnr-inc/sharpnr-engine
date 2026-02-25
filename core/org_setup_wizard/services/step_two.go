package org_setup_wizard_services

import (
	"context"
	"errors"
	org_property_contracts "sharpnr-engine/core/org_property/contracts"
	org_setup_wizard_contracts "sharpnr-engine/core/org_setup_wizard/contracts"
	org_setup_wizard_models "sharpnr-engine/core/org_setup_wizard/models"
	"time"

	"github.com/google/uuid"
)

func (s *orgSetupWizardService) Step_two_AddProperty(ctx context.Context, input org_setup_wizard_contracts.PropertySetupInput) error {

	orgVerify, err := s.orgService.VerifyOrgByID(ctx, input.OrgID)
	if err != nil {
		return err // Handle error first
	}

	if !orgVerify {
		return errors.New("organization not verified")
	}

	mappedData := org_property_contracts.CreateOrgPropertyInput{
		OrgID:       input.OrgID,
		Name:        input.Name,
		AddressLine: input.AddressLine,
		City:        input.City,
		State:       input.State,
		Country:     input.Country,
		Pincode:     input.Pincode,
	}

	if err := s.orgPropertySvc.CreatePropertyDefault(ctx, mappedData); err != nil {
		return err
	}

	if err := s.orgSetWizRepo.UpdateCurrentStepStatus(ctx, "COMPLETE", input.OrgID, input.StepID); err != nil {
		return err
	}

	step := &org_setup_wizard_models.OrganizationSetup{
		ID:             uuid.New(),
		OrganizationID: input.OrgID,
		Status:         "IN_PROGRESS",
		CurrentStep:    "LEGAL_ENTITY_SETUP",
		Version:        1,
		StartedAt:      time.Now(),
		CreatedAt:      time.Now(),
	}

	if err := s.orgSetWizRepo.Create(ctx, step); err != nil {
		return err
	}

	return nil
}
