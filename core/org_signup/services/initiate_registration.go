package org_signup_services

import (
	"context"
	"errors"
	org_signup_dtos "sharpnr-engine/core/org_signup/dtos"
	organization_contracts "sharpnr-engine/core/organization/contracts"
	"sharpnr-engine/sharp_mailer_engine"

	"golang.org/x/sync/errgroup"
)

var (
	ErrUserAlreadyExists = errors.New("user already exists")
	ErrOrgAlreadyExists  = errors.New("organization already exists")
)

func (s *orgSignUpService) InitiateRegistration(
	ctx context.Context,
	req org_signup_dtos.OrganizationRegistrationDTO,
) (string, error) {

	// --- Step 1 & 2: Parallel Uniqueness Check ---
	// We use errgroup to run both database lookups at the same time.

	g, groupCtx := errgroup.WithContext(ctx)

	g.Go(func() error {
		userExists, err := s.employeeSvc.IsUserExists(groupCtx, req.Email)
		if err != nil {
			return err
		}
		if userExists {
			return ErrUserAlreadyExists
		}
		return nil
	})

	g.Go(func() error {
		orgExists, err := s.organizationService.IsOrgExists(groupCtx, req.Email)
		if err != nil {
			return err
		}
		if orgExists {
			return ErrOrgAlreadyExists
		}
		return nil
	})

	// Wait for both checks to complete. If either returns an error, Wait() returns it.
	if err := g.Wait(); err != nil {
		return "", err
	}

	// --- Step 3: Create Draft Organization (Redis) ---
	input := organization_contracts.DraftOrganizationInput{
		Email:    req.Email,
		Mobile:   &req.Mobile,
		Password: req.Password,
	}

	slug, otp, err := s.organizationService.CreateDraftOrgForSignup(ctx, input)
	if err != nil {
		return "", err
	}

	s.mailerEngine.Enqueue(sharp_mailer_engine.MailTypeNotification, sharp_mailer_engine.NewOTPJob(
		req.Email, "Sharpnr Signup Verification",
		sharp_mailer_engine.OTPData{
			Name:          req.Mobile,
			Code:          otp,
			Email:         req.Email,
			ExpiryMinutes: 10,
		},
	))

	return slug, nil
}
