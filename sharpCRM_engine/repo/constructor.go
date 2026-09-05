package sharpCrmEngineRepo

import "gorm.io/gorm"

type repository struct {
	db *gorm.DB
}

func NewSharpCrmEngineRepo(db *gorm.DB) *repository {
	return &repository{
		db: db,
	}
}
