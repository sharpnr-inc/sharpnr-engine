package restaurantservices

import (
	"context"
	restaurant_models "sharpnr-engine/modules/hospitality/restaurant/models"
	"time"

	"github.com/google/uuid"
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
