package app

import (
	org_setup_wizard_repos "sharpnr-engine/core/org_setup_wizard/repos"
	org_setup_wizard_services "sharpnr-engine/core/org_setup_wizard/services"

	"gorm.io/gorm"
)

func bootstrapOrgSetupWizard(db *gorm.DB) org_setup_wizard_services.OrgSetupWizardServiceInterface {
	repo := org_setup_wizard_repos.NewOrgSetupWizardRepository(db)
	return org_setup_wizard_services.NewOrgSetupWizardService(repo)
}
