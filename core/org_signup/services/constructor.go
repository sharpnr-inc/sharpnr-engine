package org_signup_services

import (
	dbtx "sharpnr-engine/core/dbTx"
	org_signup_repositories "sharpnr-engine/core/org_signup/repositories"
	"sharpnr-engine/sere"
	sharp_mailer_engine "sharpnr-engine/sharp_mailer_engine"
)

type orgSignUpService struct {
	txManager         dbtx.TxManager
	sere              *sere.SERE
	tempOrgSignupRepo org_signup_repositories.TempOrgSignUpRepositoryInterface
	mainOrgSignUpRepo org_signup_repositories.MainOrgSignUpRepositoryInterface
	mailerEngine      *sharp_mailer_engine.SharpMailerEngine
}

func NewOrgSignUpService(
	txManager dbtx.TxManager,
	sere *sere.SERE,
	tempOrgSignupRepo org_signup_repositories.TempOrgSignUpRepositoryInterface,
	mainOrgSignUpRepo org_signup_repositories.MainOrgSignUpRepositoryInterface,
	mailerEngine *sharp_mailer_engine.SharpMailerEngine,
) *orgSignUpService {
	return &orgSignUpService{
		txManager:         txManager,
		sere:              sere,
		tempOrgSignupRepo: tempOrgSignupRepo,
		mainOrgSignUpRepo: mainOrgSignUpRepo,
		mailerEngine:      mailerEngine,
	}
}
