package sharpCrmEngineRepo

import "gorm.io/gorm"

type Repository struct {
	db *gorm.DB
}

func NewSharpCrmEngineRepo(db *gorm.DB) *Repository {
	return &Repository{
		db: db,
	}
}
