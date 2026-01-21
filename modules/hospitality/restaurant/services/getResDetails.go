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
		CreatedAt:      restaurant.CreatedAt,
		UpdatedAt:      restaurant.UpdatedAt,
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

	return &res, nil
}
