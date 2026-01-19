package org_signup_services

import (
	"context"
	"fmt"

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

	fmt.Println(dataForEmp)

	return uuid.Nil, uuid.Nil, err
}
