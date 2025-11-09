package menurepo

import "gorm.io/gorm"

type menuRepo struct {
	db *gorm.DB
}

func NewHospitalityMenuRepository(db *gorm.DB) *menuRepo {
	return &menuRepo{
		db: db,
	}
}
