package org_signup_services

import (
	"context"
	org_signup_dtos "sharpnr-engine/core/org_signup/dtos"
	organization_contracts "sharpnr-engine/core/organization/contracts"
	organizationErrs "sharpnr-engine/core/organization/errs"
	"sharpnr-engine/sharp_mailer_engine"
	"strings"

	"github.com/google/uuid"
	"golang.org/x/sync/errgroup"
)

func (s *orgSignUpService) InitiateRegistration(
	ctx context.Context,
	req org_signup_dtos.OrganizationRegistrationDTO,
) (uuid.UUID, error) {

	// --- Step 1 & 2: Parallel Uniqueness Check ---
	// We use errgroup to run both database lookups at the same time.

	g, groupCtx := errgroup.WithContext(ctx)

	g.Go(func() error {
		userExists, err := s.employeeSvc.IsUserExists(groupCtx, req.Email)
		if err != nil {
			return err
		}
		if userExists {
			return organizationErrs.OrganizationAlreadyExistsErr
		}
		return nil
	})

	g.Go(func() error {
		orgExists, err := s.organizationService.IsOrgExists(groupCtx, req.Email)
		if err != nil {
			return err
		}
		if orgExists {
			return organizationErrs.OrganizationAlreadyExistsErr
		}
		return nil
	})

	// Wait for both checks to complete. If either returns an error, Wait() returns it.
	if err := g.Wait(); err != nil {
		return uuid.Nil, err
	}

	// --- Step 3: Create Draft Organization (Redis) ---
	input := organization_contracts.DraftOrganizationInput{
		Email:    req.Email,
		Mobile:   &req.Mobile,
		Password: req.Password,
	}

	slug, otp, err := s.organizationService.CreateDraftOrgForSignup(ctx, input)
	if err != nil {
		return uuid.Nil, err
	}

	// Signup collects no name yet, so the mail is addressed to the email local
	// part rather than the phone number — the template greets the recipient
	// with Name, and a phone number reads as a mistake there.
	s.mailerEngine.Enqueue(sharp_mailer_engine.NewOTPJob(
		req.Email, "Sharpnr Signup Verification",
		sharp_mailer_engine.OTPData{
			Name:          recipientName(req.Email),
			Code:          otp,
			Email:         req.Email,
			ExpiryMinutes: int(s.stagingTTL.Minutes()),
		},
	))

	return slug, nil
}

// recipientName derives a human-ish greeting from an email address for mails
// sent before the user has told us their name. Falls back to the full address
// if there is nothing usable in front of the "@".
func recipientName(email string) string {
	local, _, found := strings.Cut(email, "@")
	if !found || local == "" {
		return email
	}
	return local
}
