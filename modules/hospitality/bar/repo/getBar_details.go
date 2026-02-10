package barrepo

import (
	"context"
	"errors"
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

func (r *barRepository) GetBarByPropertyId(ctx context.Context, propertyID uuid.UUID) (*bar_models.Bar, error) {
	var bar bar_models.Bar

	err := r.db.
		WithContext(ctx).
		Where("property_id = ?", propertyID).
		First(&bar).Error

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil // or return a domain-specific error
		}
		return nil, err
	}

	return &bar, nil
}
