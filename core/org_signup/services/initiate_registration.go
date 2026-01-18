package org_signup_services

import (
	"context"
	org_signup_dtos "sharpnr-engine/core/org_signup/dtos"
)

func (s *orgSignUpService) InitiateRegistration(
	ctx context.Context,
	req org_signup_dtos.OrganizationRegistrationDTO,
) (string, error)
