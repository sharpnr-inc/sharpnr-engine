package org_signup_services

import (
	"context"
	"errors"
	org_signup_dtos "sharpnr-engine/core/org-signup/dtos"
)

func (s *orgSignUpService) VerifyOtp(
	ctx context.Context,
	orgSignUpSlug string,
	otp org_signup_dtos.VerifyOtpRequest_Dto,
) error {

	// Get Org from Signing Stage
	exists, stagingOrg, err := s.tempOrgSignupRepo.Exists(ctx, orgSignUpSlug)

	if err != nil {
		return err
	}

	if !exists {
		return errors.New("signup session expired")
	}

	if stagingOrg.Otp != otp.Otp {
		return errors.New("invalid otp")
	}
	return nil
}
