package org_signup_services

import (
	employee_services "sharpnr-engine/core/employees/services"
	org_setup_wizard_services "sharpnr-engine/core/org_setup_wizard/services"
	organization_services "sharpnr-engine/core/organization/services"

	dbtx "sharpnr-engine/pkg/dbTx"
	sentinelservice "sharpnr-engine/sentinel/service"
	sharp_mailer_engine "sharpnr-engine/sharp_mailer_engine"
)

type orgSignUpService struct {
	txManager           dbtx.TxManager
	mailerEngine        *sharp_mailer_engine.SharpMailerEngine
	employeeSvc         employee_services.EmployeesIdentityServiceInterface
	organizationService organization_services.OrganizationServiceInterface
	setUpWizService     org_setup_wizard_services.OrgSetupWizardV1_ServiceInterface
	sentinelService     sentinelservice.SentinelServiceInterface
}

func NewOrgSignUpService(
	txManager dbtx.TxManager,
	mailerEngine *sharp_mailer_engine.SharpMailerEngine,
	employeeSvc employee_services.EmployeesIdentityServiceInterface,
	organizationService organization_services.OrganizationServiceInterface,
	setUpWizService org_setup_wizard_services.OrgSetupWizardV1_ServiceInterface,
	sentinelService sentinelservice.SentinelServiceInterface,
) *orgSignUpService {
	return &orgSignUpService{
		txManager:           txManager,
		mailerEngine:        mailerEngine,
		employeeSvc:         employeeSvc,
		organizationService: organizationService,
		setUpWizService:     setUpWizService,
		sentinelService:     sentinelService,
	}
}
