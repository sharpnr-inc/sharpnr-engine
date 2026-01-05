package org_signup_services

import (
	"context"
	"crypto/subtle"
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

	// 2️⃣ Compare OTP (constant-time)
	if subtle.ConstantTimeCompare(
		[]byte(stagingOrg.Otp),
		[]byte(otp.Otp),
	) != 1 {
		return errors.New("invalid otp")
	}
	return nil
}
