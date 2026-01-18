package org_setup_wizard_services

import org_setup_wizard_repos "sharpnr-engine/core/org_setup_wizard/repos"

type orgSetupWizardService struct {
	orgSetWizRepo org_setup_wizard_repos.OrganizationSetupWizardRepositoryInterface
}

func NewOrgSetupWizardService(
	orgSetWizRepo org_setup_wizard_repos.OrganizationSetupWizardRepositoryInterface,
) *orgSetupWizardService {
	return &orgSetupWizardService{
		orgSetWizRepo: orgSetWizRepo,
	}
}
