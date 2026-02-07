package restaurantrepo

import (
	"context"
	"fmt"
	restaurant_models "sharpnr-engine/modules/hospitality/restaurant/models"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

func (r *restaurantRepo) CreateRestaurant(ctx context.Context, res *restaurant_models.Restaurant) error {
	if res == nil {
		return fmt.Errorf("restaurant data is nil")
	}
	// Insert record into database
	if err := r.db.WithContext(ctx).Create(res).Error; err != nil {
		return fmt.Errorf("failed to create restaurant: %w", err)
	}

	return nil
}

func (r *restaurantRepo) GetRestaurantByID(ctx context.Context, restaurantID uuid.UUID) (*restaurant_models.Restaurant, error) {
	var restaurant restaurant_models.Restaurant

	// We use .First() because we are looking for one specific record by ID.
	// This method returns gorm.ErrRecordNotFound if the ID does not exist.
	err := r.db.WithContext(ctx).
		Where("id = ?", restaurantID).
		First(&restaurant).Error

	if err != nil {
		// We return the error (including ErrRecordNotFound) so the service layer
		// can handle it appropriately (e.g., returning a 404 Not Found).
		return nil, err
	}

	return &restaurant, nil
}

func (r *restaurantRepo) GetRestaurantByOrgID(ctx context.Context, orgId uuid.UUID) ([]*restaurant_models.Restaurant, error) {
	// 1. Define the slice to hold the results
	var restaurants []*restaurant_models.Restaurant

	// 2. Use .Find() to fetch all records matching the organization_id
	err := r.db.WithContext(ctx).
		Where("organization_id = ?", orgId).
		Find(&restaurants).Error

	if err != nil {
		// Find returns an empty slice and nil error if no records are found.
		// We only return an error if the database connection actually fails.
		return nil, err
	}

	return restaurants, nil
}

func (r *restaurantRepo) GetRestaurantByPropertyID(ctx context.Context, propertyID uuid.UUID) ([]*restaurant_models.Restaurant, error) {
	var restaurants []*restaurant_models.Restaurant

	// We filter by property_id to get all restaurants belonging to that property
	err := r.db.WithContext(ctx).
		Where("property_id = ?", propertyID).
		Find(&restaurants).Error

	if err != nil {
		// GORM .Find() returns an empty slice and no error if nothing is found.
		// We only return an error for actual database failures.
		return nil, err
	}

	return restaurants, nil
}

func (r *restaurantRepo) UpdateGeneralDetails(ctx context.Context, restaurantID uuid.UUID, updates map[string]interface{}) error {
	// .Model() tells GORM which table to use
	// .Where() ensures we only update the specific restaurant
	result := r.db.WithContext(ctx).
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
