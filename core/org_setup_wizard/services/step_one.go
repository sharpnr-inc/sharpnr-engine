package org_setup_wizard_services

import (
	"context"
	org_setup_wizard_contracts "sharpnr-engine/core/org_setup_wizard/contracts"
	organization_contracts "sharpnr-engine/core/organization/contracts"
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

	// Here we call the Repo method to update the Setup progress

	// !!! Remember to Add the Repo method to Update the Setup progress ///

	return nil
}
