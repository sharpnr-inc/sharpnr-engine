package orgCostCenterRepo

import "gorm.io/gorm"

type orgCostCenterRepoConstructor struct {
	db *gorm.DB
}

func NewOrgCostCenterRepo(db *gorm.DB) *orgCostCenterRepoConstructor {
	return &orgCostCenterRepoConstructor{
		db: db,
	}
}
