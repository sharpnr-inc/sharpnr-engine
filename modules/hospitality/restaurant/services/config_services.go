package restaurantservices

import (
	"context"
	"errors"
	"fmt"
	dbtx "sharpnr-engine/core/dbTx"
	restaurant_constants "sharpnr-engine/modules/hospitality/restaurant/constants"
	restaurant_contracts "sharpnr-engine/modules/hospitality/restaurant/contracts"
	restaurant_models "sharpnr-engine/modules/hospitality/restaurant/models"
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
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
	ctx context.Context, orgID uuid.UUID, emplID uuid.UUID, restaurantID uuid.UUID, setupInput restaurant_contracts.AddLayoutRequestDto,
) error {

	// 1. Loop through each floor in the setup input
	for _, floorInput := range setupInput.RestaurantLayout {

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

/* ---------------------------------- <***> --------------------------------- */

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

/* ---------------------------------- <***> --------------------------------- */

func (s *restaurantService) GetGeneralDetails(
	ctx context.Context,
	orgID uuid.UUID,
	emplID uuid.UUID,
	restaurantID uuid.UUID,
) (restaurant_contracts.RestaurantGeneralInfo_resp_dto, error) {

	// 1. Fetch the restaurant
	restaurant, err := s.restaurantRepo.GetRestaurantByID(ctx, restaurantID)
	if err != nil {
		return restaurant_contracts.RestaurantGeneralInfo_resp_dto{}, err
	}

	// 2. Security Check: Ensure the restaurant belongs to the requested Org
	// This prevents cross-tenant data leakage.
	if restaurant.OrganizationID != orgID {
		return restaurant_contracts.RestaurantGeneralInfo_resp_dto{}, errors.New("unauthorized: restaurant does not belong to this organization")
	}

	// 3. Mapping (Consider adding a .ToDTO() method on your entity later)
	response := restaurant_contracts.RestaurantGeneralInfo_resp_dto{
		RestaurantID: restaurant.ID,
		Name:         restaurant.Name,
		OpeningTime:  restaurant.OpeningTime,
		ClosingTime:  restaurant.ClosingTime,
		IsPureVeg:    restaurant.IsPureVeg,
		IsActive:     restaurant.IsActive,
		IsConfigured: restaurant.IsConfigured,
		LastUpdated:  restaurant.UpdatedAt,
		// Future: add flags like HasBar: restaurant.HasBar, etc.
	}

	return response, nil
}

func (s *restaurantService) UpdateGeneralDetails(
	ctx context.Context,
	orgID, emplID, restaurantID uuid.UUID,
	req restaurant_contracts.UpdateBodyRequestDto,
) error {
	// 1. Convert struct pointers to a map
	updates := make(map[string]any)

	if req.Name != nil {
		updates["name"] = *req.Name
	}
	if req.OpeningTime != nil {
		updates["opening_time"] = *req.OpeningTime
	}
	if req.ClosingTime != nil {
		updates["closing_time"] = *req.ClosingTime
	}

	// 2. Check if there's actually anything to update
	if len(updates) == 0 {
		return nil // Or return a "No changes provided" error
	}

	// 3. Execute within Transaction
	return s.txManager.WithTx(ctx, func(tx *gorm.DB) error {
		// Use the context that contains the transaction
		return s.restaurantRepo.UpdateGeneralDetails(ctx, restaurantID, updates)
	})
}

func (s *restaurantService) GetRestaurantPreferredKitchenID(
	ctx context.Context,
	restaurantID uuid.UUID,
) (restaurant_contracts.RestaurantPreferredKitchen_resp_dto, error) {

	var res restaurant_contracts.RestaurantPreferredKitchen_resp_dto

	err := s.txManager.WithTx(ctx, func(tx *gorm.DB) error {
		txCtx := dbtx.InjectTx(ctx, tx)

		preferredKitchen, err := s.restaurantRepo.
			GetRestaurantPreferredKitchenID(txCtx, restaurantID)
		if err != nil {
			return err
		}

		res = restaurant_contracts.RestaurantPreferredKitchen_resp_dto{
			RestaurantID: preferredKitchen.RestaurantID,
			KitchenID:    preferredKitchen.KitchenID,
		}

		return nil
	})

	if err != nil {
		return restaurant_contracts.RestaurantPreferredKitchen_resp_dto{}, err
	}

	return res, nil
}
