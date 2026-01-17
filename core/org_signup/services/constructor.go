package org_signup_services

import (
	dbtx "sharpnr-engine/core/dbTx"
	organization_repo "sharpnr-engine/core/organization/repo"
	"sharpnr-engine/sere"
	sharp_mailer_engine "sharpnr-engine/sharp_mailer_engine"
)

type orgSignUpService struct {
	txManager         dbtx.TxManager
	sere              *sere.SERE
	tempOrgSignupRepo organization_repo.TempOrgSignUpRepositoryInterface
	mainOrgSignUpRepo organization_repo.MainOrgSignUpRepositoryInterface
	mailerEngine      *sharp_mailer_engine.SharpMailerEngine
}

func NewOrgSignUpService(
	txManager dbtx.TxManager,
	sere *sere.SERE,
	tempOrgSignupRepo organization_repo.TempOrgSignUpRepositoryInterface,
	mainOrgSignUpRepo organization_repo.MainOrgSignUpRepositoryInterface,
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
