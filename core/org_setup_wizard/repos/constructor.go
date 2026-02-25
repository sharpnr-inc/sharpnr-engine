package org_setup_wizard_repos

import "gorm.io/gorm"

type orgSetupWizardRepo struct {
	db *gorm.DB
}

func NewOrgSetupWizardRepository(db *gorm.DB) *orgSetupWizardRepo {
	return &orgSetupWizardRepo{
		db: db,
	}
}
