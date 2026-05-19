package org_signup_services

import (
	"context"
	org_signup_dtos "sharpnr-engine/core/org_signup/dtos"

	"github.com/google/uuid"
)

// This Interface is simply SignUp Service
type OrganizationSignUpServiceInterfaces interface {

	// This is Service Start registration Process
	// This takes email, hashed password, mobile number
	// This method use service from Employee and Organization and check both
	// if we email exist in any of the repo we return the error
	// this method also have to return Temporary OrganizationID that will be used to create
	// slug or token and put in cookie to set by the controller
	InitiateRegistration(
		ctx context.Context,
		req org_signup_dtos.OrganizationRegistrationDTO,
	) (uuid.UUID, error)

	// This is method simple verify the slug or token in cookie
	// which was set by the *org_sign_controller*
	VerifyInitiationRegCookie(ctx context.Context, tmpOrgId uuid.UUID) (string, error)

	// This method simply Verify the Otp
	// After Verify the Otp
	// We create the a Default Super Admin and Organization From temporary to main table
	// and then we remove the cookie which was set by InitiateRegistration()
	// and set the New Cookie
	// Then from here things will be move on the orchestrator
	VerifyOtp(ctx context.Context, otp string, slug uuid.UUID) (orgID uuid.UUID, empID uuid.UUID, err error)
}
