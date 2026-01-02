package org_signup_services

import (
	org_signup_repositories "sharpnr-engine/core/org-signup/repositories"
	"sharpnr-engine/sere"
)

type orgSignUpService struct {
	sere              *sere.SERE
	tempOrgSignupRepo org_signup_repositories.TempOrgSignUpRepositoryInterface
	mainOrgSignUpRepo org_signup_repositories.MainOrgSignUpRepositoryInterface
}

func NewOrgSignUpService(
	sere *sere.SERE,
	tempOrgSignupRepo org_signup_repositories.TempOrgSignUpRepositoryInterface,
	mainOrgSignUpRepo org_signup_repositories.MainOrgSignUpRepositoryInterface,
) *orgSignUpService {
	return &orgSignUpService{
		sere:              sere,
		tempOrgSignupRepo: tempOrgSignupRepo,
		mainOrgSignUpRepo: mainOrgSignUpRepo,
	}
}
