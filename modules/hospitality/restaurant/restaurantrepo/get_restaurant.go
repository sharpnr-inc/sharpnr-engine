package restaurantrepo

import (
	"context"
	hospitalitymodels "sharpnr-engine/modules/hospitality/hospitality_models"

	"github.com/google/uuid"
)

func (r *restaurantRepo) GetRestaurant(ctx context.Context, orgId uuid.UUID) (*hospitalitymodels.Restaurant, error) {
	var restaurant hospitalitymodels.Restaurant

	// Query: SELECT * FROM restaurants WHERE org_id = ?
	err := r.db.WithContext(ctx).
		Where("organization_id = ?", orgId).
		First(&restaurant).Error

	if err != nil {
		return nil, err
	}

	return &restaurant, nil
}
