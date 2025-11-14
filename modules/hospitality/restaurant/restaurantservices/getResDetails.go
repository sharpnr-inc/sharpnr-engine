package restaurantservices

import (
	"context"
	hospitalitydtos "sharpnr-engine/modules/hospitality/hospitality_dtos"

	"github.com/google/uuid"
)

func (s *restaurantService) GetRestaurantDetails(ctx context.Context, orgId uuid.UUID) (*hospitalitydtos.Restaurant_GetRestaurant_ResponseDTO, error) {
	restaurant, err := s.restaurantRepo.GetRestaurant(ctx, orgId)
	if err != nil {
		return nil, err
	}

	res := &hospitalitydtos.Restaurant_GetRestaurant_ResponseDTO{
		RestaurantID:   restaurant.ID,
		ResTaurantName: restaurant.Name,
	}
	return res, nil

}

func (s *restaurantService) InitializeRestaurant() {

}
