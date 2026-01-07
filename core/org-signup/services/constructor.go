package org_signup_services

import (
	dbtx "sharpnr-engine/core/dbTx"
	hospitality_org_factory_repository "sharpnr-engine/core/hospitality-org-factory/repository"
	org_signup_repositories "sharpnr-engine/core/org-signup/repositories"
	hr_employee_services "sharpnr-engine/modules/hr/employees/services"
	"sharpnr-engine/sere"
)

type orgSignUpService struct {
	txManager              dbtx.TxManager
	sere                   *sere.SERE
	tempOrgSignupRepo      org_signup_repositories.TempOrgSignUpRepositoryInterface
	mainOrgSignUpRepo      org_signup_repositories.MainOrgSignUpRepositoryInterface
	hospitalityFactoryRepo hospitality_org_factory_repository.HospitalityOrgFactoryRepoInterface

	hrEmployeeService hr_employee_services.HrEmployeeServiceInterface
}

func NewOrgSignUpService(
	txManager dbtx.TxManager,
	sere *sere.SERE,
	tempOrgSignupRepo org_signup_repositories.TempOrgSignUpRepositoryInterface,
	mainOrgSignUpRepo org_signup_repositories.MainOrgSignUpRepositoryInterface,
	hospitalityFactoryRepo hospitality_org_factory_repository.HospitalityOrgFactoryRepoInterface,
	hrEmployeeService hr_employee_services.HrEmployeeServiceInterface,
) *orgSignUpService {
	return &orgSignUpService{
		txManager:              txManager,
		sere:                   sere,
		tempOrgSignupRepo:      tempOrgSignupRepo,
		mainOrgSignUpRepo:      mainOrgSignUpRepo,
		hospitalityFactoryRepo: hospitalityFactoryRepo,
		hrEmployeeService:      hrEmployeeService,
	}
}
