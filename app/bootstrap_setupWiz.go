package app

import (
	dbtx "sharpnr-engine/core/dbTx"
	legal_entity_service "sharpnr-engine/core/legal_entity/service"
	org_property_services "sharpnr-engine/core/org_property/services"
	org_setup_wizard_repos "sharpnr-engine/core/org_setup_wizard/repos"
	org_setup_wizard_services "sharpnr-engine/core/org_setup_wizard/services"
	organization_services "sharpnr-engine/core/organization/services"
	sharpnr_identity_auth_services "sharpnr-engine/core/sharpnr_identity/services"
	sharp_mailer_engine "sharpnr-engine/sharp_mailer_engine"

	"gorm.io/gorm"
)

func bootstrapOrgSetupWizard(
	db *gorm.DB,
	orgLegalEntitySvc legal_entity_service.OrgLegalEntityServiceInterface,
	orgPropertySvc org_property_services.OrganizationPropertyServicesInterfaces,
	orgService organization_services.OrganizationServiceInterface,
	sharpIdentity sharpnr_identity_auth_services.SharpnrIdentityAuthServiceInterface,
	txManager dbtx.TxManager,
	mailerEngine *sharp_mailer_engine.SharpMailerEngine,
) org_setup_wizard_services.OrgSetupWizardV1_ServiceInterface {
	repo := org_setup_wizard_repos.NewOrgSetupWizardRepository(db)
	return org_setup_wizard_services.NewOrgSetupWizardService(
		repo,
		orgLegalEntitySvc,
		orgPropertySvc,
		orgService,
		sharpIdentity,
		txManager,
		mailerEngine,
	)
}
