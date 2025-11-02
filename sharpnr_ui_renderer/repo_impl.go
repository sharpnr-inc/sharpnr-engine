package sharpnruirenderer

import (
	"sharpnr-engine/shared/database"

	"gorm.io/gorm"
)

type orgUIRendererRepository struct {
	db *gorm.DB
}

func NewOrgUIRendererRepo(db *gorm.DB) *orgUIRendererRepository {
	return &orgUIRendererRepository{
		db: database.DB,
	}
}
