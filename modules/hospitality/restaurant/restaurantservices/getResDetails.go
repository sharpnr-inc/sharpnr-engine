package restaurantservices

import (
	"context"
	hospitalitydtos "sharpnr-engine/modules/hospitality/hospitality_dtos"

	"github.com/google/uuid"
)

func (s *restaurantService) GetRestaurantDetails(
	ctx context.Context,
	orgId uuid.UUID,
) (*hospitalitydtos.Restaurant_GetRestaurant_ResponseDTO, error) {

	restaurant, err := s.restaurantRepo.GetRestaurant(ctx, orgId)
	if err != nil {
		return nil, err
	}

	// ---- MAP MAIN RESTAURANT ----
	res := &hospitalitydtos.Restaurant_GetRestaurant_ResponseDTO{
		ID:             restaurant.ID,
		Name:           restaurant.Name,
		OrganizationID: restaurant.OrganizationID,
		IsPureVeg:      restaurant.IsPureVeg,
		OpeningTime:    restaurant.OpeningTime,
		ClosingTime:    restaurant.ClosingTime,

		LegalName:   restaurant.LegalName,
		DisplayName: restaurant.DisplayName,
		GSTNumber:   restaurant.GSTNumber,
		FSSAINumber: restaurant.FSSAINumber,

		Address: restaurant.Address,
		City:    restaurant.City,
		State:   restaurant.State,
		Pincode: restaurant.Pincode,

		IsKitchen: restaurant.IsKitchen,
		KitchenID: restaurant.KitchenID,
		CreatedAt: restaurant.CreatedAt,
		UpdatedAt: restaurant.UpdatedAt,
	}

	// ---- MAP FLOORS ----
	for _, floor := range restaurant.Floors {
		floorDto := hospitalitydtos.RestaurantFloorDTO{
			ID:             floor.ID,
			RestaurantID:   floor.RestaurantID,
			OrganizationID: floor.OrganizationID,
			Name:           floor.Name,
			Description:    floor.Description,
			FloorNumber:    floor.FloorNumber,
			CreatedAt:      floor.CreatedAt,
			UpdatedAt:      floor.UpdatedAt,
		}

		// ---- MAP TABLES INSIDE FLOOR ----
		for _, tbl := range floor.Tables {
			floorDto.Tables = append(floorDto.Tables, hospitalitydtos.RestaurantTableDTO{
				ID:           tbl.ID,
				RestaurantID: tbl.RestaurantID,
				FloorID:      tbl.FloorID,
				Name:         tbl.Name,
				Capacity:     tbl.Capacity,
				Status:       tbl.Status,
				CreatedAt:    tbl.CreatedAt,
				UpdatedAt:    tbl.UpdatedAt,
			})
		}

		res.Floors = append(res.Floors, floorDto)
	}

	// ---- MAP DIRECT RESTAURANT TABLES (if any) ----
	for _, tbl := range restaurant.Tables {
		res.Tables = append(res.Tables, hospitalitydtos.RestaurantTableDTO{
			ID:           tbl.ID,
			RestaurantID: tbl.RestaurantID,
			FloorID:      tbl.FloorID,
			Name:         tbl.Name,
			Capacity:     tbl.Capacity,
			Status:       tbl.Status,
			CreatedAt:    tbl.CreatedAt,
			UpdatedAt:    tbl.UpdatedAt,
		})
	}

	// ---- MAP KITCHEN (nullable) ----
	if restaurant.Kitchen != nil {
		res.Kitchen = &hospitalitydtos.KitchenDTO{
			ID:   restaurant.Kitchen.ID,
			Name: restaurant.Kitchen.Name,
		}
	}

	return res, nil
}

func (s *restaurantService) InitializeRestaurant() {

}
