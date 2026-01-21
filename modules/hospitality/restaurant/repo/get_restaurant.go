package restaurantrepo

import (
	"context"
	restaurant_models "sharpnr-engine/modules/hospitality/restaurant/models"

	"github.com/google/uuid"
)

func (r *restaurantRepo) GetRestaurantByOrgId(ctx context.Context, orgId uuid.UUID) (*restaurant_models.Restaurant, error) {
	var restaurant restaurant_models.Restaurant

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
) (*restaurant_models.RestaurantFloor, error) {

	var floor restaurant_models.RestaurantFloor

	err := r.db.WithContext(ctx).
		Where("id = ?", floorId).
		First(&floor).Error

	if err != nil {
		return nil, err
	}

	return &floor, nil
}

func (d *restaurantRepo) GetTableById(ctx context.Context, tableId uuid.UUID) (*restaurant_models.RestaurantTable, error) {

	var table restaurant_models.RestaurantTable

	err := d.db.WithContext(ctx).First(&table, "id = ?", tableId).Error
	if err != nil {
		return nil, err
	}

	return &table, nil
}
