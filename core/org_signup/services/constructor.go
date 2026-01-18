package org_signup_services

import (
	dbtx "sharpnr-engine/core/dbTx"
	organization_services "sharpnr-engine/core/organization/services"
	employee_services "sharpnr-engine/modules/hr/employees/services"
	"sharpnr-engine/sere"
	sharp_mailer_engine "sharpnr-engine/sharp_mailer_engine"
)

type orgSignUpService struct {
	txManager           dbtx.TxManager
	sere                *sere.SERE
	mailerEngine        *sharp_mailer_engine.SharpMailerEngine
	employeeSvc         employee_services.EmployeeServiceInterface
	organizationService organization_services.OrganizationServiceInterface
}

func NewOrgSignUpService(
	txManager dbtx.TxManager,
	sere *sere.SERE,
	mailerEngine *sharp_mailer_engine.SharpMailerEngine,
	employeeSvc employee_services.EmployeeServiceInterface,
	organizationService organization_services.OrganizationServiceInterface,
) *orgSignUpService {
	return &orgSignUpService{
		txManager:           txManager,
		sere:                sere,
		mailerEngine:        mailerEngine,
		employeeSvc:         employeeSvc,
		organizationService: organizationService,
	}
}
