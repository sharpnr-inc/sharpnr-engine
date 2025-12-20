package restaurantrepo

import (
	"context"
	hospitalitymodels "sharpnr-engine/modules/hospitality/hospitality_models"

	"github.com/google/uuid"
)

func (r *restaurantRepo) GetRestaurantByOrgId(ctx context.Context, orgId uuid.UUID) (*hospitalitymodels.Restaurant, error) {
	var restaurant hospitalitymodels.Restaurant

	err := r.db.WithContext(ctx).
		Preload("Organization"). // <-- load organization
		Preload("Kitchen").      // <-- load kitchen
		Preload("Floors").
		Preload("Floors.Tables").
		Where("organization_id = ?", orgId).
		First(&restaurant).Error

	if err != nil {
		return nil, err
	}

	return &restaurant, nil
}

func (r *restaurantRepo) GetRestaurantFloorById(
	ctx context.Context,
	floorId uuid.UUID,
) (*hospitalitymodels.RestaurantFloor, error) {

	var floor hospitalitymodels.RestaurantFloor

	err := r.db.WithContext(ctx).
		Where("id = ?", floorId).
		First(&floor).Error

	if err != nil {
		return nil, err
	}

	return &floor, nil
}

func (d *restaurantRepo) GetTableById(ctx context.Context, tableId uuid.UUID) (*hospitalitymodels.RestaurantTable, error) {

	var table hospitalitymodels.RestaurantTable

	err := d.db.WithContext(ctx).First(&table, "id = ?", tableId).Error
	if err != nil {
		return nil, err
	}

	return &table, nil
}
