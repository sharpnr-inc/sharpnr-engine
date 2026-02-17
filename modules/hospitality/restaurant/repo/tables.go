package restaurantrepo

import (
	"context"
	"fmt"
	dbtx "sharpnr-engine/core/dbTx"
	restaurant_models "sharpnr-engine/modules/hospitality/restaurant/models"

	"github.com/google/uuid"
)

func (r *restaurantRepo) CreateTables(ctx context.Context, resTable *restaurant_models.RestaurantTable) error {
	db, ok := dbtx.ExtractTx(ctx)
	if !ok {
		db = r.db
	}
	return db.WithContext(ctx).Create(resTable).Error
}

func (r *restaurantRepo) UpdateTableStatus(ctx context.Context, tableID uuid.UUID, status string) error {
	// We use .Model() to identify the table and .Update() for the specific column
	result := r.db.WithContext(ctx).
		Model(&restaurant_models.RestaurantTable{}).
		Where("id = ?", tableID).
		Update("status", status)

	if result.Error != nil {
		return result.Error
	}

	// Checking if the record actually existed
	if result.RowsAffected == 0 {
		return fmt.Errorf("table not found with ID: %s", tableID)
	}

	return nil
}

func (d *restaurantRepo) GetTableById(ctx context.Context, tableId uuid.UUID) (*restaurant_models.RestaurantTable, error) {

	db, ok := dbtx.ExtractTx(ctx)
	if !ok {
		db = d.db
	}
	var table restaurant_models.RestaurantTable

	err := db.WithContext(ctx).First(&table, "id = ?", tableId).Error
	if err != nil {
		return nil, err
	}

	return &table, nil
}

func (r *restaurantRepo) GetTablesByStatus(ctx context.Context, restaurantID uuid.UUID, status string) ([]*restaurant_models.RestaurantTable, error) {
	db, ok := dbtx.ExtractTx(ctx)
	if !ok {
		db = r.db
	}
	var tables []*restaurant_models.RestaurantTable

	// We filter by both the parent restaurant ID and the specific status.
	// Using .Find() ensures we get all matching tables in a slice.
	err := db.WithContext(ctx).
		Where("restaurant_id = ? AND status = ?", restaurantID, status).
		Find(&tables).Error

	if err != nil {
		// Find returns an empty slice and nil error if no matching tables are found.
		return nil, err
	}

	return tables, nil
}

func (r *restaurantRepo) GetAllTablesOfFloor(ctx context.Context, floorID uuid.UUID) ([]*restaurant_models.RestaurantTable, error) {
	db, ok := dbtx.ExtractTx(ctx)
	if !ok {
		db = r.db
	}
	var tables []*restaurant_models.RestaurantTable

	// Filtering strictly by floor_id without a specific sort order
	err := db.WithContext(ctx).
		Where("floor_id = ?", floorID).
		Find(&tables).Error

	if err != nil {
		return nil, err
	}

	return tables, nil
}
