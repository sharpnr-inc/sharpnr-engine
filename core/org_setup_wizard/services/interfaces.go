package org_setup_wizard_services

import (
	"context"
	org_setup_wizard_contracts "sharpnr-engine/core/org_setup_wizard/contracts"
)

// This si version V1 interface
// to Upgrade it Just add another interface making the
type OrgSetupWizardServiceInterface interface {
	Step_one_AddOrgHqAddrAndGeoLocation(ctx context.Context, input org_setup_wizard_contracts.StepOneOrganizationHQ) error
	Step_two_AddProperty(ctx context.Context)
	Step_three_AddLegalEntity(ctx context.Context)
}
