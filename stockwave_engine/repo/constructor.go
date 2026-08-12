package stockwaveEngineRepo

import "gorm.io/gorm"

type inventoryEngineRepo struct {
	db *gorm.DB
}

func NewInventoryEngineRepo(db *gorm.DB) *inventoryEngineRepo {
	return &inventoryEngineRepo{
		db: db,
	}
}
