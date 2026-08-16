package org_signup_services

import (
	"time"

	coreEmployeeServices "sharpnr-engine/core/employees/services"
	org_setup_wizard_services "sharpnr-engine/core/org_setup_wizard/services"
	organization_services "sharpnr-engine/core/organization/services"

	dbtx "sharpnr-engine/pkg/dbTx"
	sentinelservice "sharpnr-engine/sentinel/service"
	sharp_mailer_engine "sharpnr-engine/sharp_mailer_engine"
)

type orgSignUpService struct {
	txManager           dbtx.TxManager
	mailerEngine        *sharp_mailer_engine.SharpMailerEngine
	employeeSvc         coreEmployeeServices.EmployeeSignUpServiceInterface
	organizationService organization_services.OrganizationServiceInterface
	setUpWizService     org_setup_wizard_services.OrgSetupWizardV1_ServiceInterface
	sentinelService     sentinelservice.SentinelServiceInterface

	// stagingTTL is the lifetime of the Redis staging record, which is also the
	// OTP validity window. It is injected rather than hardcoded so the expiry
	// quoted in the verification email cannot drift away from the real one.
	stagingTTL time.Duration
}

func NewOrgSignUpService(
	txManager dbtx.TxManager,
	mailerEngine *sharp_mailer_engine.SharpMailerEngine,
	employeeSvc coreEmployeeServices.EmployeeSignUpServiceInterface,
	organizationService organization_services.OrganizationServiceInterface,
	setUpWizService org_setup_wizard_services.OrgSetupWizardV1_ServiceInterface,
	sentinelService sentinelservice.SentinelServiceInterface,
	stagingTTL time.Duration,
) *orgSignUpService {
	return &orgSignUpService{
		txManager:           txManager,
		mailerEngine:        mailerEngine,
		employeeSvc:         employeeSvc,
		organizationService: organizationService,
		setUpWizService:     setUpWizService,
		sentinelService:     sentinelService,
		stagingTTL:          stagingTTL,
	}
}
