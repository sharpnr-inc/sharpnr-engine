package org_signup_services

import (
	"context"

	"github.com/google/uuid"
)

func (r *orgSignUpService) VerifyInitiationRegCookie(ctx context.Context, tmpOrgId uuid.UUID) error
