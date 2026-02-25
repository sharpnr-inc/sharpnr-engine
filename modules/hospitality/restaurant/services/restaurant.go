package restaurantservices

import (
	"context"
	"errors"
	restaurant_contracts "sharpnr-engine/modules/hospitality/restaurant/contracts"
	restaurant_models "sharpnr-engine/modules/hospitality/restaurant/models"
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

func (r *restaurantService) SetupDefaultRestaurant(ctx context.Context, orgID uuid.UUID, propertyID uuid.UUID) (restuarantID uuid.UUID, err error) {

	defRes := &restaurant_models.Restaurant{
		ID:             uuid.New(),
		Name:           "Default Restaurant",
		OrganizationID: orgID,
		PropertyID:     propertyID,
		CreatedAt:      time.Now(),
	}

	if err := r.restaurantRepo.CreateRestaurant(ctx, defRes); err != nil {
		return uuid.Nil, err
	}

	return defRes.ID, nil

}

/* ----------------------------------- *** ---------------------------------- */

func (r *restaurantService) GetAllRestaurantOfProperty(
	ctx context.Context,
	propertyID uuid.UUID,
) ([]*restaurant_contracts.RestaurantDetails, error) {

	restaurants, err := r.restaurantRepo.GetRestaurantByPropertyID(ctx, propertyID)
	if err != nil {
		// IMPORTANT: record not found → return empty slice
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return []*restaurant_contracts.RestaurantDetails{}, nil
		}
		return nil, err
	}

	result := make([]*restaurant_contracts.RestaurantDetails, 0, len(restaurants))

	for _, rs := range restaurants {
		result = append(result, &restaurant_contracts.RestaurantDetails{
			ID:           rs.ID,
			Name:         rs.Name,
			Status:       false,
			IsConfigured: rs.IsConfigured,
		})
	}

	return result, nil
}

/* ---------------------------------- <**> ---------------------------------- */

func (r *restaurantService) GetRestaurantLayout(
	ctx context.Context,
	orgID uuid.UUID,
	restaurantID uuid.UUID,
) (layout restaurant_contracts.RestaurantLayout_Response_dto, err error) {

	floors, err := r.restaurantRepo.GetAllFloorOfRestaurant(ctx, restaurantID)
	if err != nil {
		return layout, err
	}

	var floorLayouts []restaurant_contracts.RestaurantFloorLayout_Response_dto

	for _, floor := range floors {
		tables, err := r.restaurantRepo.GetAllTablesOfFloor(ctx, floor.ID)
		if err != nil {
			return layout, err
		}

		var tableDTOs []restaurant_contracts.RestaurantTableLayout_Response_dto

		for _, table := range tables {
			// 1. Fetch the current session for this table
			session, err := r.restaurantRepo.GetDiningSessionByTableID(ctx, table.ID)

			var sessionID *uuid.UUID
			if err == nil && session != nil {
				// Session found
				sessionID = &session.ID
			} else if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
				// Actual DB error (connection, etc.), stop the execution
				return layout, err
			}
			// If err is ErrRecordNotFound, sessionID remains nil, which is correct.

			tableDTOs = append(tableDTOs, restaurant_contracts.RestaurantTableLayout_Response_dto{
				TableName:            table.Name,
				TableCode:            table.Code,
				FloorID:              table.FloorID.String(),
				TableID:              table.ID.String(),
				Status:               table.Status,
				Capacity:             table.Capacity,
				CurrentDineSessionID: sessionID, // 2. Attach the ID here
			})
		}

		floorLayouts = append(floorLayouts, restaurant_contracts.RestaurantFloorLayout_Response_dto{
			Floor_Name:   floor.Name,
			Floor_Number: floor.FloorNumber,
			Tables:       tableDTOs,
		})
	}

	layout.Layout = floorLayouts
	return layout, nil
}

/* -------------------------------- **<***>** ------------------------------- */
