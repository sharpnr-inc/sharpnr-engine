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

func (r *restaurantRepo) GetRestaurantPreferredKitchenID(ctx context.Context, restaurantID uuid.UUID) (*restaurant_models.RestaurantPreferedKitchen, error) {
	db, ok := dbtx.ExtractTx(ctx)
	if !ok {
		db = r.db
	}

	var resPrefKitchen restaurant_models.RestaurantPreferedKitchen
	err := db.WithContext(ctx).
		Where("restaurant_id = ?", restaurantID).
		First(&resPrefKitchen).Error

	if err != nil {
		return nil, err
	}

	return &resPrefKitchen, nil
}

func (r *restaurantRepo) CreatePreferredKitchen(
	ctx context.Context,
	pfktcn *restaurant_models.RestaurantPreferedKitchen,
) error {

	db := r.db
	if tx, ok := dbtx.ExtractTx(ctx); ok {
		db = tx
	}

	return db.WithContext(ctx).
		Create(pfktcn).
		Error
}

func (r *restaurantRepo) UpdatePreferredKitchen(
	ctx context.Context,
	pfk *restaurant_models.RestaurantPreferedKitchen,
) error {

	db := r.db
	if tx, ok := dbtx.ExtractTx(ctx); ok {
		db = tx
	}

	return db.
		Model(&restaurant_models.RestaurantPreferedKitchen{}).
		Where("restaurant_id = ?", pfk.RestaurantID).
		Updates(map[string]interface{}{
			"kitchen_id": pfk.KitchenID,
			"updated_at": pfk.UpdatedAt,
		}).Error
}
