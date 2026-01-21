package restaurantrepo

import (
	"context"
	restaurant_models "sharpnr-engine/modules/hospitality/restaurant/models"

	"github.com/google/uuid"
)

func (r *restaurantRepo) CreateFloor(ctx context.Context, floor *restaurant_models.RestaurantFloor) error {
	// .Create performs the SQL INSERT
	// It returns an error if constraints (like foreign keys) are violated
	return r.db.WithContext(ctx).Create(floor).Error
}

func (r *restaurantRepo) GetFloorByID(ctx context.Context, floorID uuid.UUID) (*restaurant_models.RestaurantFloor, error) {
	var floor restaurant_models.RestaurantFloor

	// .First() adds 'LIMIT 1' and returns gorm.ErrRecordNotFound if the ID doesn't exist
	err := r.db.WithContext(ctx).
		Where("id = ?", floorID).
		First(&floor).Error

	if err != nil {
		return nil, err
	}

	return &floor, nil
}

func (r *restaurantRepo) GetAllFloorOfRestaurant(ctx context.Context, restaurantID uuid.UUID) ([]*restaurant_models.RestaurantFloor, error) {
	var floors []*restaurant_models.RestaurantFloor

	// We filter the restaurant_floors table by the restaurant_id
	// Adding an Order clause is usually best practice for floor lists
	err := r.db.WithContext(ctx).
		Where("restaurant_id = ?", restaurantID).
		Order("floor_number ASC").
		Find(&floors).Error

	if err != nil {
		// .Find() returns an empty slice and no error if no floors exist.
		// We return the error only for database-level failures.
		return nil, err
	}

	return floors, nil
}
