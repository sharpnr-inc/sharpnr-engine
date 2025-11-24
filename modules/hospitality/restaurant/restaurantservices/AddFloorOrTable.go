package restaurantservices

import (
	"context"
	"fmt"
	hospitalitydtos "sharpnr-engine/modules/hospitality/hospitality_dtos"
	hospitalitymodels "sharpnr-engine/modules/hospitality/hospitality_models"

	"github.com/google/uuid"
)

func (r *restaurantService) CreateFloor(ctx context.Context) {}

func (r *restaurantService) CreateOrAddTable(ctx context.Context, orgid uuid.UUID, req hospitalitydtos.Restaurant_CreateOrAddTables_Request) error {

	fmt.Println("---- BEFORE SERVICE ----")

	// 1️⃣ Check restaurant owned by this org
	restaurant, err := r.restaurantRepo.GetRestaurantByOrgId(ctx, orgid)
	if err != nil {
		return fmt.Errorf("restaurant not found: %w", err)
	}
	// Optional validation
	if restaurant.ID != req.RestaurantId {
		return fmt.Errorf("restaurant mismatch")
	}

	// 2️⃣ Validate floor exists
	floor, err := r.restaurantRepo.GetRestaurantFloorById(ctx, req.FloorId)
	if err != nil {
		return fmt.Errorf("floor not found: %w", err)
	}

	if floor.RestaurantID != restaurant.ID {
		return fmt.Errorf("floor does not belong to the restaurant")
	}

	// 3️⃣ Convert DTO → Model list
	var tables []*hospitalitymodels.RestaurantTable

	for _, t := range req.Tables {

		tables = append(tables, &hospitalitymodels.RestaurantTable{
			ID:           uuid.New(),
			FloorID:      req.FloorId,
			RestaurantID: restaurant.ID,
			Restaurant:   restaurant,
			Floor:        floor,
			Name:         t.TableName,
			Capacity:     t.TableCapacity,
			Status:       "AVAILABLE",
		})
	}

	fmt.Println(tables)

	// 4️⃣ Save all tables (bulk insert)
	if err := r.restaurantRepo.CreateOrAddTablesInRestuarant(ctx, tables); err != nil {
		return fmt.Errorf("failed to create tables: %w", err)
	}

	return nil
}
