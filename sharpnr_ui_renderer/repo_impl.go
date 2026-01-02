package sharpnruirenderer

import (
	"gorm.io/gorm"
)

type orgUIRendererRepository struct {
	db *gorm.DB
}

// GetOrgUIRenderer implements UIRendererRepository.
func (o *orgUIRendererRepository) GetOrgUIRenderer() {
	panic("unimplemented")
}

func NewOrgUIRendererRepo(db *gorm.DB) *orgUIRendererRepository {
	return &orgUIRendererRepository{
		db: db,
	}
}
