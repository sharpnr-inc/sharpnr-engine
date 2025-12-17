package hospitality_setup

import (
	"context"
	"time"

	hospitalitymodels "sharpnr-engine/modules/hospitality/hospitality_models"

	"github.com/google/uuid"
)

type RestaurantSetupStep struct{}

func (s *RestaurantSetupStep) Name() string { return "RESTAURANT" }

func (s *RestaurantSetupStep) Execute(ctx context.Context, c SetupContext) error {

	restaurant := &hospitalitymodels.Restaurant{
		ID:             uuid.New(),
		Name:           "Default Restaurant",
		OrganizationID: c.Org.ID,
		Organization:   c.Org,
		IsKitchen:      true,
		KitchenID:      &c.Kitchen.ID,
		CreatedAt:      time.Now(),
		UpdatedAt:      time.Now(),
	}

	if err := c.RestaurantRepo.CreateRestaurant(ctx, restaurant); err != nil {
		return err
	}

	menu := &hospitalitymodels.MenuGroup{
		ID:             uuid.New(),
		OrganizationID: c.Org.ID,
		Name:           "Restaurant Menu",
		Description:    "Default restaurant menu",
		Type:           "RESTAURANT",
		CreatedAt:      time.Now(),
		UpdatedAt:      time.Now(),
	}

	return c.MenuRepo.CreateMenuGroup(ctx, menu)
}
