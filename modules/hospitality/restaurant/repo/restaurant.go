package restaurantrepo

import (
	"context"
	"fmt"
	dbtx "sharpnr-engine/core/dbTx"
	restaurant_models "sharpnr-engine/modules/hospitality/restaurant/models"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

func (r *restaurantRepo) CreateRestaurant(ctx context.Context, res *restaurant_models.Restaurant) error {
	db, ok := dbtx.ExtractTx(ctx)
	if !ok {
		db = r.db
	}
	if res == nil {
		return fmt.Errorf("restaurant data is nil")
	}
	// Insert record into database
	if err := db.WithContext(ctx).Create(res).Error; err != nil {
		return fmt.Errorf("failed to create restaurant: %w", err)
	}

	return nil
}

func (r *restaurantRepo) GetRestaurantByID(ctx context.Context, restaurantID uuid.UUID) (*restaurant_models.Restaurant, error) {
	db, ok := dbtx.ExtractTx(ctx)
	if !ok {
		db = r.db
	}

	var restaurant restaurant_models.Restaurant
	err := db.WithContext(ctx).
		Where("id = ?", restaurantID).
		First(&restaurant).Error

	if err != nil {
		return nil, err
	}

	return &restaurant, nil
}

func (r *restaurantRepo) GetRestaurantByOrgID(ctx context.Context, orgId uuid.UUID) ([]*restaurant_models.Restaurant, error) {
	db, ok := dbtx.ExtractTx(ctx)
	if !ok {
		db = r.db
	}

	var restaurants []*restaurant_models.Restaurant
	err := db.WithContext(ctx).
		Where("organization_id = ?", orgId).
		Find(&restaurants).Error

	if err != nil {
		return nil, err
	}

	return restaurants, nil
}

func (r *restaurantRepo) GetRestaurantByPropertyID(ctx context.Context, propertyID uuid.UUID) ([]*restaurant_models.Restaurant, error) {
	db, ok := dbtx.ExtractTx(ctx)
	if !ok {
		db = r.db
	}

	var restaurants []*restaurant_models.Restaurant
	err := db.WithContext(ctx).
		Where("property_id = ?", propertyID).
		Find(&restaurants).Error

	if err != nil {
		return nil, err
	}

	return restaurants, nil
}

func (r *restaurantRepo) UpdateGeneralDetails(ctx context.Context, restaurantID uuid.UUID, updates map[string]interface{}) error {
	db, ok := dbtx.ExtractTx(ctx)
	if !ok {
		db = r.db
	}

	result := db.WithContext(ctx).
		Model(&restaurant_models.Restaurant{}).
		Where("id = ?", restaurantID).
		Updates(updates)

	if result.Error != nil {
		return result.Error
	}

	if result.RowsAffected == 0 {
		return gorm.ErrRecordNotFound
	}

	return nil
}
