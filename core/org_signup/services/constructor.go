package org_signup_services

import (
	org_setup_wizard_services "sharpnr-engine/core/org_setup_wizard/services"
	organization_services "sharpnr-engine/core/organization/services"
	employee_services "sharpnr-engine/hrms_engine/employees/services"
	dbtx "sharpnr-engine/pkg/dbTx"
	sentinelservice "sharpnr-engine/sentinel_engine/service"
	sharp_mailer_engine "sharpnr-engine/sharp_mailer_engine"
)

type orgSignUpService struct {
	txManager           dbtx.TxManager
	mailerEngine        *sharp_mailer_engine.SharpMailerEngine
	employeeSvc         employee_services.EmployeesIdentityServiceInterface
	organizationService organization_services.OrganizationServiceInterface
	setUpWizService     org_setup_wizard_services.OrgSetupWizardV1_ServiceInterface
	sentinelService     sentinelservice.SentinelService
}

func NewOrgSignUpService(
	txManager dbtx.TxManager,
	mailerEngine *sharp_mailer_engine.SharpMailerEngine,
	employeeSvc employee_services.EmployeesIdentityServiceInterface,
	organizationService organization_services.OrganizationServiceInterface,
	setUpWizService org_setup_wizard_services.OrgSetupWizardV1_ServiceInterface,
	sentinelService sentinelservice.SentinelService,
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
