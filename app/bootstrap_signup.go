package app

import (
	org_setup_wizard_services "sharpnr-engine/core/org_setup_wizard/services"
	org_signup_services "sharpnr-engine/core/org_signup/services"
	organization_services "sharpnr-engine/core/organization/services"
	employee_services "sharpnr-engine/hrms_engine/employees/services"
	dbtx "sharpnr-engine/pkg/dbTx"
	sentinelservice "sharpnr-engine/sentinel_engine/service"
	sharp_mailer_engine "sharpnr-engine/sharp_mailer_engine"
)

func bootstrapOrgSignup(
	txManager dbtx.TxManager,
	mailerEngine *sharp_mailer_engine.SharpMailerEngine,
	employeeSvc employee_services.EmployeesIdentityServiceInterface,
	organizationService organization_services.OrganizationServiceInterface,
	setUpWizService org_setup_wizard_services.OrgSetupWizardV1_ServiceInterface,
	sentinelService sentinelservice.SentinelService,
) org_signup_services.OrganizationSignUpServiceInterfaces {

	return org_signup_services.NewOrgSignUpService(
		txManager,
		mailerEngine,
		employeeSvc,
		organizationService,
		setUpWizService,
		sentinelService,
	)
}
