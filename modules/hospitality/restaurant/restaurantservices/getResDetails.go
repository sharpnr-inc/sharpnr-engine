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

	restaurant, err := s.restaurantRepo.GetRestaurantByOrgId(ctx, orgId)
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
		CINNumber:   restaurant.CINNumber,
		FSSAINumber: restaurant.FSSAINumber,

		ContactEmail:  restaurant.ContactEmail,
		ContactNumber: restaurant.ContactNumber,
		ManagerName:   restaurant.ManagerName,

		Address: restaurant.Address,
		City:    restaurant.City,
		State:   restaurant.State,
		Country: restaurant.Country,
		Pincode: restaurant.Pincode,

		IsKitchen: restaurant.IsKitchen,
		KitchenID: restaurant.KitchenID,
		CreatedAt: restaurant.CreatedAt,
		UpdatedAt: restaurant.UpdatedAt,
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

func (s *restaurantService) GetRestaurantLayout(ctx context.Context, orgid uuid.UUID) (*hospitalitydtos.Restaurant_GetRestaurantLayout_Response, error) {
	restaurant, err := s.restaurantRepo.GetRestaurantByOrgId(ctx, orgid)
	if err != nil {
		return nil, err
	}

	res := hospitalitydtos.Restaurant_GetRestaurantLayout_Response{
		RestaurantId: restaurant.ID,
		Floor:        make([]hospitalitydtos.RestaurantFloorDTO, 0),
	}

	// 3️⃣ Iterate floors
	for _, flr := range restaurant.Floors {

		floorDTO := hospitalitydtos.RestaurantFloorDTO{
			ID:             flr.ID,
			RestaurantID:   flr.RestaurantID,
			OrganizationID: flr.OrganizationID,
			Name:           flr.Name,
			Description:    flr.Description,
			FloorNumber:    flr.FloorNumber,
			CreatedAt:      flr.CreatedAt,
			UpdatedAt:      flr.UpdatedAt,
			Tables:         make([]hospitalitydtos.RestaurantTableDTO, 0),
		}

		// 4️⃣ Add tables inside floor
		for _, tbl := range flr.Tables {

			tableDTO := hospitalitydtos.RestaurantTableDTO{
				ID:           tbl.ID,
				RestaurantID: tbl.RestaurantID,
				FloorID:      &tbl.FloorID,
				Name:         tbl.Name,
				Capacity:     tbl.Capacity,
				Status:       tbl.Status,
				CreatedAt:    tbl.CreatedAt,
				UpdatedAt:    tbl.UpdatedAt,
			}

			floorDTO.Tables = append(floorDTO.Tables, tableDTO)
		}

		// Add floorDTO to response
		res.Floor = append(res.Floor, floorDTO)
	}

	return &res, nil
}

func (s *restaurantService) InitializeRestaurant() {

}
