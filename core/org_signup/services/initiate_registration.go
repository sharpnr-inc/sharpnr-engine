package org_signup_services

import (
	"context"
	"errors"
	org_signup_dtos "sharpnr-engine/core/org_signup/dtos"
	organization_contracts "sharpnr-engine/core/organization/contracts"
	"sharpnr-engine/sharp_mailer_engine"
)

var (
	ErrUserAlreadyExists = errors.New("user already exists")
	ErrOrgAlreadyExists  = errors.New("organization already exists")
)

func (s *orgSignUpService) InitiateRegistration(
	ctx context.Context,
	req org_signup_dtos.OrganizationRegistrationDTO,
) (string, error) {

	// Step 1: Check if employee/user already exists
	userExists, err := s.employeeSvc.IsUserExists(ctx, req.Email)
	if err != nil {
		return "", err
	}

	if userExists {
		return "", ErrUserAlreadyExists
	}

	// Step 2: Check if organization already exists
	orgExists, err := s.organizationService.IsOrgExists(ctx, req.Email, req.Mobile)
	if err != nil {
		return "", err
	}

	if orgExists {
		return "", ErrOrgAlreadyExists
	}

	// Step 3: Create the Organization temporary Organization in the Redis

	input := organization_contracts.DraftOrganizationInput{
		Email:    req.Email,
		Mobile:   &req.Mobile,
		Password: req.Password,
	}

	slug, otp, err := s.organizationService.CreateDraftOrgForSignup(ctx, input)
	if err != nil {
		return "", err
	}

	otpMailData := map[string]interface{}{
		"Code":   otp,
		"Expiry": "10 minutes",
	}

	s.mailerEngine.Queues[sharp_mailer_engine.OTP] <- sharp_mailer_engine.EmailJob{
		To:       req.Email,
		Subject:  "Email Verification",
		Template: sharp_mailer_engine.TmplVerification,
		Data:     otpMailData,
	}

	return slug, nil
}
