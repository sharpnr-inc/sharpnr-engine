package hospitality_org_factory_repository

import "gorm.io/gorm"

type hospitalityOrgFactoryRepository struct {
	db *gorm.DB
}

func NewHospitalityOrgFactoryRepository(db *gorm.DB) *hospitalityOrgFactoryRepository {
	return &hospitalityOrgFactoryRepository{
		db: db,
	}
}
