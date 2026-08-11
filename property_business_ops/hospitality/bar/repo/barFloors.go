package barrepo

import (
	"context"
	bar_models "sharpnr-engine/property_business_ops/hospitality/bar/models"

	"github.com/google/uuid"
)

func (r *barRepository) GetFloorById(ctx context.Context, floorId uuid.UUID) (*bar_models.BarFloor, error) {
	var floor bar_models.BarFloor

	err := r.db.WithContext(ctx).
		Where("id = ?", floorId).
		First(&floor).Error

	if err != nil {
		return nil, err
	}

	return &floor, nil
}
