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

func (r *restaurantService) AddFloors(
	ctx context.Context,
	AddFloorInput *restaurant_contracts.CreateFloorInput,
) (floorID uuid.UUID, err error) {

	err = r.txManager.WithTx(ctx, func(tx *gorm.DB) error {

		txCtx := dbtx.InjectTx(ctx, tx)

		input := &restaurant_models.RestaurantFloor{
			ID:           uuid.New(),
			Name:         AddFloorInput.FloorName,
			RestaurantID: AddFloorInput.RestaurantID,
			FloorNumber:  AddFloorInput.FloorNumber,
			CreatedAt:    time.Now(),
		}

		if err := r.restaurantRepo.CreateFloor(txCtx, input); err != nil {
			return err
		}

		floorID = input.ID
		return nil
	})

	if err != nil {
		return uuid.Nil, err
	}

	return floorID, nil
}

/* ---------------------------------- <***> --------------------------------- */

func (r *restaurantService) AddTablesInFloor(
	ctx context.Context,
	AddTablesInFloorInput *restaurant_contracts.CreateTablesInFloorInput,
) (tableID uuid.UUID, err error) {

	err = r.txManager.WithTx(ctx, func(tx *gorm.DB) error {

		txCtx := dbtx.InjectTx(ctx, tx)

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

		if err := r.restaurantRepo.CreateTables(txCtx, input); err != nil {
			return err
		}

		tableID = input.ID
		return nil
	})

	if err != nil {
		return uuid.Nil, err
	}

	return tableID, nil
}

/* ---------------------------------- <***> --------------------------------- */

func (r *restaurantService) SetupRestaurantLayout(
	ctx context.Context,
	orgID uuid.UUID,
	emplID uuid.UUID,
	restaurantID uuid.UUID,
	setupInput restaurant_contracts.AddLayoutRequestDto,
) error {

	return r.txManager.WithTx(ctx, func(tx *gorm.DB) error {

		txCtx := dbtx.InjectTx(ctx, tx)

		// Loop through each floor
		for _, floorInput := range setupInput.RestaurantLayout {

			floorID, err := r.AddFloors(txCtx, &restaurant_contracts.CreateFloorInput{
				RestaurantID: restaurantID,
				FloorName:    floorInput.Floor_Name,
				FloorNumber:  floorInput.Floor_Number,
			})
			if err != nil {
				return fmt.Errorf("failed to create floor %s: %w", floorInput.Floor_Name, err)
			}

			// Create tables for this floor
			for i := 1; i <= floorInput.Number_of_tables; i++ {

				tableName := fmt.Sprintf("Table %d", i)
				tableCode := fmt.Sprintf("F%d-T%d", floorInput.Floor_Number, i)

				_, err := r.AddTablesInFloor(txCtx, &restaurant_contracts.CreateTablesInFloorInput{
					RestaurantID: restaurantID,
					FloorID:      floorID,
					Name:         &tableName,
					Code:         &tableCode,
				})

				if err != nil {
					return fmt.Errorf(
						"failed to create table %d on floor %s: %w",
						i,
						floorInput.Floor_Name,
						err,
					)
				}
			}
		}

		configStatus, err := r.restaurantRepo.GetCurrentConfig(txCtx, restaurantID)
		if err != nil {
			return err
		}

		if configStatus.CurrentStatus == restaurant_constants.RestaurantStatusGeneralDetailsAdded ||
			configStatus.CurrentStatus == restaurant_constants.RestaurantStatusCreated {

			if err := r.UpdateConfigurationStatus(
				txCtx,
				orgID,
				restaurantID,
				restaurant_constants.RestaurantStatusCompleted,
			); err != nil {
				return err
			}

			if err := r.restaurantRepo.UpdateConfiguration(txCtx, restaurantID, true); err != nil {
				return err
			}
		}

		return nil
	})
}

/* ---------------------------------- <***> --------------------------------- */

func (s *restaurantService) GetGeneralDetails(
	ctx context.Context,
	orgID uuid.UUID,
	emplID uuid.UUID,
	restaurantID uuid.UUID,
) (restaurant_contracts.RestaurantGeneralInfo_resp_dto, error) {

	var response restaurant_contracts.RestaurantGeneralInfo_resp_dto

	err := s.txManager.WithTx(ctx, func(tx *gorm.DB) error {

		txCtx := dbtx.InjectTx(ctx, tx)

		// 1. Fetch restaurant
		restaurant, err := s.restaurantRepo.GetRestaurantByID(txCtx, restaurantID)
		if err != nil {
			return err
		}

		// 2. Security check (multi-tenant safety)
		if restaurant.OrganizationID != orgID {
			return errors.New("unauthorized: restaurant does not belong to this organization")
		}

		// 3. Fetch configuration
		resConfig, err := s.restaurantRepo.GetCurrentConfig(txCtx, restaurant.ID)
		if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
			return err
		}

		// 4. Map base restaurant fields
		response = restaurant_contracts.RestaurantGeneralInfo_resp_dto{
			RestaurantID: restaurant.ID,
			Name:         restaurant.Name,
			OpeningTime:  restaurant.OpeningTime,
			ClosingTime:  restaurant.ClosingTime,
			IsPureVeg:    restaurant.IsPureVeg,
			LastUpdated:  restaurant.UpdatedAt,
		}

		// 5. Map configuration fields
		if resConfig != nil {
			response.IsActive = resConfig.IsActive
			response.IsConfigured = resConfig.IsConfigured
			response.CurrentConfigStatus = resConfig.CurrentStatus
		}

		return nil
	})

	if err != nil {
		return restaurant_contracts.RestaurantGeneralInfo_resp_dto{}, err
	}

	return response, nil
}

/* ---------------------------------- <***> --------------------------------- */

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

		txCtx := dbtx.InjectTx(ctx, tx)

		if err := s.restaurantRepo.UpdateGeneralDetails(txCtx, restaurantID, updates); err != nil {
			return err
		}

		configStatus, err := s.restaurantRepo.GetCurrentConfig(txCtx, restaurantID)
		if err != nil {
			return err
		}

		if configStatus.CurrentStatus == restaurant_constants.RestaurantStatusCreated {
			if err := s.UpdateConfigurationStatus(
				txCtx,
				orgID,
				restaurantID,
				restaurant_constants.RestaurantStatusGeneralDetailsAdded,
			); err != nil {
				return err
			}
		}

		return nil
	})
}

/* ---------------------------------- <***> --------------------------------- */

func (s *restaurantService) GetRestaurantPreferredKitchenID(
	ctx context.Context,
	orgID, employeeID, restaurantID uuid.UUID,
) (restaurant_contracts.RestaurantPreferredKitchen_resp_dto, error) {

	// Default response → both nil
	res := restaurant_contracts.RestaurantPreferredKitchen_resp_dto{
		RestaurantID: nil,
		KitchenID:    nil,
	}

	err := s.txManager.WithTx(ctx, func(tx *gorm.DB) error {
		txCtx := dbtx.InjectTx(ctx, tx)

		preferredKitchen, err := s.restaurantRepo.
			GetRestaurantPreferredKitchenID(txCtx, restaurantID)

		if err != nil {
			if errors.Is(err, gorm.ErrRecordNotFound) {
				// Not found → keep nil fields
				return nil
			}
			return err
		}

		// Set values only if found
		res.RestaurantID = &preferredKitchen.RestaurantID
		res.KitchenID = &preferredKitchen.KitchenID

		return nil
	})

	if err != nil {
		return restaurant_contracts.RestaurantPreferredKitchen_resp_dto{}, err
	}

	return res, nil
}

/* ---------------------------------- <***> --------------------------------- */

func (s *restaurantService) AddPreferredKitchen(
	ctx context.Context,
	orgID, employeeID,
	restaurantID, kitchenID uuid.UUID,
) error {

	return s.txManager.WithTx(ctx, func(tx *gorm.DB) error {
		txCtx := dbtx.InjectTx(ctx, tx)

		// Check existing preferred kitchen
		existing, err := s.restaurantRepo.GetRestaurantPreferredKitchenID(txCtx, restaurantID)
		if err != nil {
			if !errors.Is(err, gorm.ErrRecordNotFound) {
				return err
			}

			// Not found → Create new
			pfk := restaurant_models.RestaurantPreferedKitchen{
				ID:           uuid.New(),
				RestaurantID: restaurantID,
				KitchenID:    kitchenID,
				CreatedAt:    time.Now(),
				UpdatedAt:    time.Now(),
			}

			return s.restaurantRepo.CreatePreferredKitchen(txCtx, &pfk)
		}

		// Found → Update KitchenID
		existing.KitchenID = kitchenID
		existing.UpdatedAt = time.Now()

		return s.restaurantRepo.UpdatePreferredKitchen(txCtx, existing)
	})
}

/* ---------------------------------- <***> --------------------------------- */

func (s *restaurantService) AddConfigurationStatus(ctx context.Context, orgID uuid.UUID, restaurantID uuid.UUID, status restaurant_constants.RestaurantConfigurationStatusType) error {

	return s.txManager.WithTx(ctx, func(tx *gorm.DB) error {

		txCtx := dbtx.InjectTx(ctx, tx)

		config := &restaurant_models.RestaurantConfiguration{
			ID:             uuid.New(),
			OrganizationID: orgID,
			RestaurantID:   restaurantID,
			IsActive:       false,
			IsConfigured:   false,
			CurrentStatus:  status,
			CreatedAt:      time.Now(),
			UpdatedAt:      time.Now(),
		}

		if err := s.restaurantRepo.CreateRestaurantConfiguration(txCtx, config); err != nil {
			return err
		}

		return nil
	})

}

/* ---------------------------------- <***> --------------------------------- */

func (s *restaurantService) UpdateConfigurationStatus(ctx context.Context, orgID uuid.UUID, restaurantID uuid.UUID, status restaurant_constants.RestaurantConfigurationStatusType) error {

	return s.txManager.WithTx(ctx, func(tx *gorm.DB) error {
		txCtx := dbtx.InjectTx(ctx, tx)

		return s.restaurantRepo.UpdateRestuarantConfigStatus(txCtx, restaurantID, status)

	})
}

/* ---------------------------------- <***> --------------------------------- */
