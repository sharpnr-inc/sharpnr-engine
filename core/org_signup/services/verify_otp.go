package org_signup_services

import (
	"context"
	org_setup_wizards_constants "sharpnr-engine/core/org_setup_wizard/constants"

	"github.com/google/uuid"
)

func (r *orgSignUpService) VerifyOtp(
	ctx context.Context,
	otp string,
	slug uuid.UUID,
) (orgID uuid.UUID, emp uuid.UUID, err error) {

	dataForEmp, err := r.organizationService.VerifySlugCreateDefaultOrg(ctx, otp, slug)
	if err != nil {
		return uuid.Nil, uuid.Nil, err
	}

	output, err := r.organizationService.CreateNewOrganization(ctx, dataForEmp)
	if err != nil {
		return uuid.Nil, uuid.Nil, err
	}

	empId, err := r.employeeSvc.CreateDefaultSuperAdmin(ctx, output.OrgID, output.Email, output.HashedPass)
	if err != nil {
		return uuid.Nil, uuid.Nil, err
	}

	// Authorization: give the founding employee a full-access "Super Admin" role
	// for this org (role + wildcard permission + assignment, in one transaction).
	// if _, err := r.sentinelService.ProvisionSuperAdmin(ctx, output.OrgID, empId); err != nil {
	// 	return uuid.Nil, uuid.Nil, err
	// }

	// * This is Add Progress in Setup Wizard
	// * and fetch the setup progress it will put in the org_setup_wizard table and then it will be show in the UI that the user is in the Super Admin Profile step
	if err := r.setUpWizService.AddStepProgress(
		ctx,
		output.OrgID,
		org_setup_wizards_constants.SetupStepSuperAdminProfile,
		org_setup_wizards_constants.SetupStatusInProgress); err != nil {
		return uuid.Nil, uuid.Nil, err
	}

	return output.OrgID, empId, err
}
