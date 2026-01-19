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

	r.employeeSvc.CreateDefaultSuperAdmin(ctx, output.OrgID, output.Email, output.HashedPass)

	return uuid.Nil, uuid.Nil, err
}
