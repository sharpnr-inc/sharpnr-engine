package restaurantservices

import (
	"context"
	"fmt"
	restaurant_constants "sharpnr-engine/modules/hospitality/restaurant/constants"
	restaurant_contracts "sharpnr-engine/modules/hospitality/restaurant/contracts"
	restaurant_models "sharpnr-engine/modules/hospitality/restaurant/models"
	"time"

	"github.com/google/uuid"
)

// AddFloors implements RestaurantServiceInterface.
func (r *restaurantService) AddFloors(ctx context.Context, AddFloorInput *restaurant_contracts.CreateFloorInput) (floorID uuid.UUID, err error) {
	input := &restaurant_models.RestaurantFloor{
		ID:           uuid.New(),
		Name:         AddFloorInput.FloorName,
		RestaurantID: AddFloorInput.RestaurantID,
		FloorNumber:  AddFloorInput.FloorNumber,
		CreatedAt:    time.Now(),
	}

	if err := r.restaurantRepo.CreateFloor(ctx, input); err != nil {
		return uuid.Nil, err
	}

	return input.ID, nil
}

// AddTablesInFloor implements RestaurantServiceInterface.
func (r *restaurantService) AddTablesInFloor(ctx context.Context, AddTablesInFloorInput *restaurant_contracts.CreateTablesInFloorInput) (tableID uuid.UUID, err error) {
	input := &restaurant_models.RestaurantTable{
		ID:           uuid.New(),
		RestaurantID: AddTablesInFloorInput.RestaurantID,
		FloorID:      AddTablesInFloorInput.FloorID,
		Code:         *AddTablesInFloorInput.Code,
		Name:         *AddTablesInFloorInput.Name,
		Capacity:     4,
		Status:       restaurant_constants.TableStatusAvailable,
		CreatedAt:    time.Now(),
		UpdatedAt:    time.Now(),
	}

	if err := r.restaurantRepo.CreateTables(ctx, input); err != nil {
		return uuid.Nil, err
	}

	return input.ID, nil
}

func (r *restaurantService) SetupRestaurantLayout(
	ctx context.Context,
	propertyID uuid.UUID,
	restaurantID uuid.UUID,
	setupInput []restaurant_contracts.SetupRestaurantLayoutInput,
) error {

	// 1. Loop through each floor in the setup input
	for _, floorInput := range setupInput {

		// 2. Call your existing AddFloors method
		floorID, err := r.AddFloors(ctx, &restaurant_contracts.CreateFloorInput{
			RestaurantID: restaurantID,
			FloorName:    floorInput.Floor_Name,
			FloorNumber:  floorInput.Floor_Number,
		})
		if err != nil {
			return fmt.Errorf("failed to create floor %s: %w", floorInput.Floor_Name, err)
		}

		// 3. Loop to create the specified number of tables for this floor
		for i := 1; i <= floorInput.Number_of_tables; i++ {
			tableName := fmt.Sprintf("Table %d", i)
			tableCode := fmt.Sprintf("F%d-T%d", floorInput.Floor_Number, i)

			// 4. Call your existing AddTablesInFloor method
			_, err := r.AddTablesInFloor(ctx, &restaurant_contracts.CreateTablesInFloorInput{
				RestaurantID: restaurantID,
				FloorID:      floorID,
				Name:         &tableName, // Passing pointers as your method requires
				Code:         &tableCode,
			})
			if err != nil {
				return fmt.Errorf("failed to create table %d on floor %s: %w", i, floorInput.Floor_Name, err)
			}
		}
	}

	return nil
}

func (r *restaurantService) UpdateRestaurantGeneralDetails(
	ctx context.Context,
	restaurantID uuid.UUID,
	input *restaurant_contracts.UpdateGeneralDetailsInput,
) error {
	// We use a map to ensure boolean 'false' values are not ignored by GORM
	updates := make(map[string]interface{})

	if input.Name != "" {
		updates["name"] = input.Name
	}

	// For Booleans, we directly assign them
	// (Assuming the Input contract uses pointers or the Orchestrator validates these)
	updates["is_configured"] = input.IsConfigured
	updates["is_pure_veg"] = input.IsPureVeg

	if input.OpeningTime != "" {
		updates["opening_time"] = input.OpeningTime
	}

	if input.ClosingTime != "" {
		updates["closing_time"] = input.ClosingTime
	}

	// Call the repository method
	return r.restaurantRepo.UpdateGeneralDetails(ctx, restaurantID, updates)
}
