package org_signup_services

import (
	dbtx "sharpnr-engine/core/dbTx"
	org_signup_repositories "sharpnr-engine/core/org-signup/repositories"
	"sharpnr-engine/sere"
)

type orgSignUpService struct {
	txManager         dbtx.TxManager
	sere              *sere.SERE
	tempOrgSignupRepo org_signup_repositories.TempOrgSignUpRepositoryInterface
	mainOrgSignUpRepo org_signup_repositories.MainOrgSignUpRepositoryInterface
}

func NewOrgSignUpService(
	txManager dbtx.TxManager,
	sere *sere.SERE,
	tempOrgSignupRepo org_signup_repositories.TempOrgSignUpRepositoryInterface,
	mainOrgSignUpRepo org_signup_repositories.MainOrgSignUpRepositoryInterface,
) *orgSignUpService {
	return &orgSignUpService{
		txManager:         txManager,
		sere:              sere,
		tempOrgSignupRepo: tempOrgSignupRepo,
		mainOrgSignUpRepo: mainOrgSignUpRepo,
	}
}
