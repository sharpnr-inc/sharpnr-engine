package barrepo

import (
	"context"
	hospitalitymodels "sharpnr-engine/modules/hospitality/hospitality_models"

	"github.com/google/uuid"
)

func (r *barRepository) GetFloorById(ctx context.Context, floorId uuid.UUID) (*hospitalitymodels.BarFloor, error) {
	var floor hospitalitymodels.BarFloor

	err := r.db.WithContext(ctx).
		Where("id = ?", floorId).
		First(&floor).Error

	if err != nil {
		return nil, err
	}

	return &floor, nil
}
