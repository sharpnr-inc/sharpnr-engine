package org_setup_wizard_services

import (
	"context"
	"errors"
	org_property_contracts "sharpnr-engine/core/org_property/contracts"
	org_setup_wizard_contracts "sharpnr-engine/core/org_setup_wizard/contracts"
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
	return nil
}
