package org_signup_services

import (
	"context"
	org_signup_dtos "sharpnr-engine/core/org-signup/dtos"
)

type OrgSignUpServiceInterface interface {
	InitiateSignup(
		ctx context.Context,
		req org_signup_dtos.InitiateOrgSignUpRequest_Dto,
	) (string, error)

	VerifyOtp(
		ctx context.Context,
		orgSignUpSlug string,
		otp org_signup_dtos.VerifyOtpRequest_Dto,
	) error

	VerifySlug(
		ctx context.Context,
		slug string,
	) (*org_signup_dtos.StagingOrgResponseDto, error)

	// Create Organization
	StartOrgOnBoarding(
		ctx context.Context,
		stagingSlug string,
		req org_signup_dtos.OrgOnBoardingRequest_dto,
	) (*org_signup_dtos.OrgOnBoardingSuccessResponse_dto, error)
}
