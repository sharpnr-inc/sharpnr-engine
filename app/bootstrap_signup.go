package app

import (
	org_setup_wizard_services "sharpnr-engine/core/org_setup_wizard/services"
	org_signup_services "sharpnr-engine/core/org_signup/services"
	organization_services "sharpnr-engine/core/organization/services"
	dbtx "sharpnr-engine/pkg/dbTx"
	employee_services "sharpnr-engine/sharp_hrms_engine/employees/services"
	sharp_mailer_engine "sharpnr-engine/sharp_mailer_engine"
)

func bootstrapOrgSignup(
	txManager dbtx.TxManager,
	mailerEngine *sharp_mailer_engine.SharpMailerEngine,
	employeeSvc employee_services.EmployeesIdentityServiceInterface,
	organizationService organization_services.OrganizationServiceInterface,
	setUpWizService org_setup_wizard_services.OrgSetupWizardV1_ServiceInterface,
) org_signup_services.OrganizationSignUpServiceInterfaces {

	return org_signup_services.NewOrgSignUpService(
		txManager,
		mailerEngine,
		employeeSvc,
		organizationService,
		setUpWizService,
	)
}
