package barrepo

import "gorm.io/gorm"

type barRepository struct {
	db *gorm.DB
}

func NewBarRepository(db *gorm.DB) *barRepository {
	return &barRepository{
		db: db,
	}
}
