package cablesEngineRepo

import "gorm.io/gorm"

type cablesRepo struct {
	db *gorm.DB
}

func NewCablesRepo(db *gorm.DB) *cablesRepo {
	return &cablesRepo{
		db: db,
	}
}
