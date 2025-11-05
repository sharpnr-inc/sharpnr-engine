package sharpnruirenderer

import (
	"gorm.io/gorm"
)

type orgUIRendererRepository struct {
	db *gorm.DB
}

func NewOrgUIRendererRepo(db *gorm.DB) *orgUIRendererRepository {
	return &orgUIRendererRepository{
		db: db,
	}
}
