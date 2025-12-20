package hospitality_setup

import (
	hospitalitymodels "sharpnr-engine/modules/hospitality/hospitality_models"
	"time"

	"github.com/google/uuid"
	"golang.org/x/net/context"
)

type HotelSetupStep struct{}

func (s *HotelSetupStep) Name() string { return "HOTEL" }

func (s *HotelSetupStep) Execute(ctx context.Context, c SetupContext) error {

	hotel := &hospitalitymodels.Hotel{
		ID:             uuid.New(),
		OrganizationID: c.Org.ID,
		Organization:   c.Org,
		Name:           "Default Hotels",
		IsKitchen:      true,
		KitchenID:      &c.Kitchen.ID,
		CreatedAt:      time.Now(),
		UpdatedAt:      time.Now(),
	}

	if err := c.HotelRepo.CreateHotel(ctx, hotel); err != nil {
		return err
	}

	menu := &hospitalitymodels.MenuGroup{
		ID:             uuid.New(),
		OrganizationID: c.Org.ID,
		Name:           "Room Service Menu",
		Description:    "Default hotel menu",
		Type:           "HOTEL",
		CreatedAt:      time.Now(),
		UpdatedAt:      time.Now(),
	}

	return c.MenuRepo.CreateMenuGroup(ctx, menu)
}
