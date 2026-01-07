package org_signup_services

import (
	"sharpnr-engine/app"
	dbtx "sharpnr-engine/core/dbTx"
	hospitality_org_factory_repository "sharpnr-engine/core/hospitality-org-factory/repository"
	org_signup_repositories "sharpnr-engine/core/org-signup/repositories"
	"sharpnr-engine/sere"
)

type orgSignUpService struct {
	txManager              dbtx.TxManager
	sere                   *sere.SERE
	tempOrgSignupRepo      org_signup_repositories.TempOrgSignUpRepositoryInterface
	mainOrgSignUpRepo      org_signup_repositories.MainOrgSignUpRepositoryInterface
	hospitalityFactoryRepo hospitality_org_factory_repository.HospitalityOrgFactoryRepoInterface
	sharpModules           app.ModulesProvider
}

func NewOrgSignUpService(
	txManager dbtx.TxManager,
	sere *sere.SERE,
	tempOrgSignupRepo org_signup_repositories.TempOrgSignUpRepositoryInterface,
	mainOrgSignUpRepo org_signup_repositories.MainOrgSignUpRepositoryInterface,
	hospitalityFactoryRepo hospitality_org_factory_repository.HospitalityOrgFactoryRepoInterface,
	sharpModules app.ModulesProvider,
) *orgSignUpService {
	return &orgSignUpService{
		txManager:              txManager,
		sere:                   sere,
		tempOrgSignupRepo:      tempOrgSignupRepo,
		mainOrgSignUpRepo:      mainOrgSignUpRepo,
		hospitalityFactoryRepo: hospitalityFactoryRepo,
		sharpModules:           sharpModules,
	}
}
