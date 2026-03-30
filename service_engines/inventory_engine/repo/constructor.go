package inventory_engine_repo

import "gorm.io/gorm"

type inventortyEngineRepo struct {
	db *gorm.DB
}

func NewInventoryEngineRepo(db *gorm.DB) *inventortyEngineRepo {
	return &inventortyEngineRepo{
		db: db,
	}
}
