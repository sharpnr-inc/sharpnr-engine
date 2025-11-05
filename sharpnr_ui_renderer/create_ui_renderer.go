package sharpnruirenderer

import (
	"context"
	"sharpnr-engine/core/models"

	"gorm.io/gorm"
)

func (o *orgUIRendererRepository) CreateOrgUIRenderer(ctx context.Context, renderer *models.OrgUIRendererTree) error {
	// Defensive nil check
	if renderer == nil {
		return gorm.ErrInvalidData
	}

	// Use WithContext to ensure request-level cancellation and tracing
	if err := o.db.WithContext(ctx).Create(renderer).Error; err != nil {
		return err
	}

	return nil
}
