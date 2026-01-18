package org_signup_services

import (
	"context"
	"errors"
	org_signup_dtos "sharpnr-engine/core/org_signup/dtos"
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
	orgExists, err := s.organizationService.IsOrgExists(ctx, req.Email)
	if err != nil {
		return "", err
	}

	if orgExists {
		return "", ErrOrgAlreadyExists
	}

	// Step 3: Initiate registration flow
	// Example: generate temporary registration token / request id
	registrationID := "temp-registration-id" // TODO: replace with UUID / token generator

	return registrationID, nil
}
