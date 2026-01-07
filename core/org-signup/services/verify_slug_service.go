package org_signup_services

import (
	"context"
	org_signup_dtos "sharpnr-engine/core/org-signup/dtos"
)

func (c *orgSignUpService) VerifySlug(ctx context.Context, slug string) (*org_signup_dtos.StagingOrgResponseDto, error) {

	stageOrg, err := c.tempOrgSignupRepo.Get(ctx, slug)
	if err != nil {
		return nil, err
	}

	stagingOrg := org_signup_dtos.StagingOrgResponseDto{
		OrgName:  stageOrg.Name,
		OrgEmail: stageOrg.Email,
		OrgPhone: stageOrg.Phone,
	}

	return &stagingOrg, nil
}
