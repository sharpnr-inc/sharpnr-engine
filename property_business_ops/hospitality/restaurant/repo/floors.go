package restaurantrepo

import (
	"context"
	"fmt"
	restaurant_models "sharpnr-engine/property_business_ops/hospitality/restaurant/models"
	dbtx "sharpnr-engine/pkg/dbTx"

	"github.com/google/uuid"
)

func (r *restaurantRepo) CreateFloor(ctx context.Context, floor *restaurant_models.RestaurantFloor) error {
	db, ok := dbtx.ExtractTx(ctx)
	if !ok {
		db = r.db
	}

	if floor == nil {
		return fmt.Errorf("floor data is nil")
	}

	// .Create performs the SQL INSERT
	// It returns an error if constraints (like foreign keys) are violated
	return db.WithContext(ctx).Create(floor).Error
}

func (r *restaurantRepo) GetFloorByID(ctx context.Context, floorID uuid.UUID) (*restaurant_models.RestaurantFloor, error) {
	db, ok := dbtx.ExtractTx(ctx)
	if !ok {
		db = r.db
	}

	var floor restaurant_models.RestaurantFloor

	// .First() adds 'LIMIT 1' and returns gorm.ErrRecordNotFound if the ID doesn't exist
	err := db.WithContext(ctx).
		Where("id = ?", floorID).
		First(&floor).Error

	if err != nil {
		return nil, err
	}

	return &floor, nil
}

func (r *restaurantRepo) GetAllFloorOfRestaurant(ctx context.Context, restaurantID uuid.UUID) ([]*restaurant_models.RestaurantFloor, error) {
	db, ok := dbtx.ExtractTx(ctx)
	if !ok {
		db = r.db
	}

	var floors []*restaurant_models.RestaurantFloor

	// We filter the restaurant_floors table by the restaurant_id
	// Adding an Order clause is usually best practice for floor lists
	err := db.WithContext(ctx).
		Where("restaurant_id = ?", restaurantID).
		Order("floor_number ASC").
		Find(&floors).Error

	if err != nil {
		return nil, err
	}

	return floors, nil
}
