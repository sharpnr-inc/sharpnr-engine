package org_signup_services

import (
	"context"

	"github.com/google/uuid"
)

func (r *orgSignUpService) VerifyInitiationRegCookie(
	ctx context.Context,
	tmpOrgID uuid.UUID) (string, error) {

	email, err := r.organizationService.GetSignupSlug(ctx, tmpOrgID)
	if err != nil {
		return "", err
	}

	return email, nil

}
