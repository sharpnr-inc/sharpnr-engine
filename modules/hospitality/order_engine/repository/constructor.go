package hospitality_order_engine_repository

import (
	"gorm.io/gorm"
)

type orderEngineRepo struct {
	db *gorm.DB
}

func NewHospitalityOrderEngine(db *gorm.DB) *orderEngineRepo {
	return &orderEngineRepo{
		db: db,
	}
}
