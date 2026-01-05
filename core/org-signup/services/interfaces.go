package org_signup_services

import (
	"context"
	org_signup_dtos "sharpnr-engine/core/org-signup/dtos"
)

type OrgSignUpServiceInterface interface {

	// ─────────────────────────────
	// STEP 1: INITIATE SIGNUP
	// ─────────────────────────────
	// - Validate input
	// - Check main DB (email / phone / slug)
	// - Create temp org in Redis (TTL)
	// - Generate & send OTP
	// - Return signup_id (tracking key)
	InitiateSignup(ctx context.Context, req org_signup_dtos.InitiateOrgSignUpRequest_Dto) (string, error)

	VerifyOtp(
		ctx context.Context,
		orgSignUpSlug string,
		otp org_signup_dtos.VerifyOtpRequest_Dto,
	) error

	VerifySlug(ctx context.Context, slug string) (*org_signup_dtos.StagingOrgResponseDto, error)
}
