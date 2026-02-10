package app

import (
	dbtx "sharpnr-engine/core/dbTx"
	org_setup_wizard_services "sharpnr-engine/core/org_setup_wizard/services"
	org_signup_services "sharpnr-engine/core/org_signup/services"
	organization_services "sharpnr-engine/core/organization/services"
	employee_services "sharpnr-engine/modules/hr/employees/services"
	"sharpnr-engine/sere"
	sharp_mailer_engine "sharpnr-engine/sharp_mailer_engine"
)

func bootstrapOrgSignup(
	txManager dbtx.TxManager,
	sereEngine *sere.SERE,
	mailerEngine *sharp_mailer_engine.SharpMailerEngine,
	employeeSvc employee_services.EmployeeServiceInterface,
	organizationService organization_services.OrganizationServiceInterface,
	setUpWizService org_setup_wizard_services.OrgSetupWizardServiceInterface,
) org_signup_services.OrganizationSignUpServiceInterfaces {

	return org_signup_services.NewOrgSignUpService(
		txManager,
		sereEngine,
		mailerEngine,
		employeeSvc,
		organizationService,
		setUpWizService,
	)
}
