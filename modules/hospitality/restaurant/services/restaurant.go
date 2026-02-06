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
			IsConfigured: false,
		})
	}

	return result, nil
}
