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

func (r *restaurantService) GetRestaurantLayout(
	ctx context.Context,
	orgID uuid.UUID, // PropertyID
	restaurantID uuid.UUID, // RestaurantID
) (layout restaurant_contracts.RestaurantLayout_Response_dto, err error) {

	// 1. Fetch all floors for the restaurant
	floors, err := r.restaurantRepo.GetAllFloorOfRestaurant(ctx, restaurantID)
	if err != nil {
		return layout, err
	}

	// 2. Initialize the result slice
	var floorLayouts []restaurant_contracts.RestaurantFloorLayout_Response_dto

	// 3. For each floor, fetch its tables and map to DTO
	for _, floor := range floors {
		tables, err := r.restaurantRepo.GetAllTablesOfFloor(ctx, floor.ID)
		if err != nil {
			return layout, err
		}

		var tableDTOs []restaurant_contracts.RestaurantTableLayout_Response_dto
		for _, table := range tables {
			tableDTOs = append(tableDTOs, restaurant_contracts.RestaurantTableLayout_Response_dto{
				TableName: table.Name,
				TableCode: table.Code,
				FloorID:   table.FloorID.String(),
				TableID:   table.ID.String(),
				Status:    table.Status,
				Capacity:  table.Capacity,
			})
		}

		floorLayouts = append(floorLayouts, restaurant_contracts.RestaurantFloorLayout_Response_dto{
			Floor_Name:   floor.Name,
			Floor_Number: floor.FloorNumber,
			Tables:       tableDTOs,
		})
	}

	// 4. Wrap in final DTO
	layout.Layout = floorLayouts
	return layout, nil
}
