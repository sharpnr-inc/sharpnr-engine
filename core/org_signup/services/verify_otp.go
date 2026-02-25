package org_signup_services

import (
	"context"

	"github.com/google/uuid"
)

func (r *orgSignUpService) VerifyOtp(
	ctx context.Context,
	otp string,
	slug uuid.UUID,
) (orgID uuid.UUID, emp uuid.UUID, err error) {

	// Step one is give the Slug and otp for the Verification
	// and if the verification is complete then we

	dataForEmp, err := r.organizationService.VerifySlugCreateDefaultOrg(ctx, otp, slug)
	if err != nil {
		return uuid.Nil, uuid.Nil, err
	}

	output, err := r.organizationService.CreateNewOrganization(ctx, dataForEmp)
	if err != nil {
		return uuid.Nil, uuid.Nil, err
	}

	// *This is Add Progress in Setup Wizard
	// *it Set the Org and put in the In progress so when the user reach the orchestrator
	// * and fetch the setup progress it will put in the org_set
	if err := r.setUpWizService.AddStepProgress(ctx, output.OrgID, "ORG_HQ_SETUP", "IN_PROGRESS"); err != nil {
		return uuid.Nil, uuid.Nil, err
	}

	empId, err := r.employeeSvc.CreateDefaultSuperAdmin(ctx, output.OrgID, output.Email, output.HashedPass)
	if err != nil {
		return uuid.Nil, uuid.Nil, err
	}

	return output.OrgID, empId, err
}
