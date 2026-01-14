package barrepo

import (
	"context"
	bar_models "sharpnr-engine/modules/hospitality/bar/models"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

func (r *barRepository) GetBar(ctx context.Context, barId uuid.UUID, orgId uuid.UUID) (*bar_models.Bar, error) {

	panic("unimplemented")
}

func (r *barRepository) GetBarByOrgId(ctx context.Context, orgId uuid.UUID) (*bar_models.Bar, error) {

	var bar bar_models.Bar

	err := r.db.WithContext(ctx).
		Preload("Floors.Tables").
		Where("organization_id = ?", orgId).
		First(&bar).Error

	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, nil // no bar found for org
		}
		return nil, err // db or query error
	}

	return &bar, nil
}
