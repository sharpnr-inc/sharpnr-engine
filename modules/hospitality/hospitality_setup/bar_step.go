package hospitality_setup

import (
	"context"
	hospitalitymodels "sharpnr-engine/modules/hospitality/hospitality_models"
	"time"

	"github.com/google/uuid"
)

type BarSetupStep struct{}

func (s *BarSetupStep) Name() string { return "BAR" }

func (s *BarSetupStep) Execute(ctx context.Context, c SetupContext) error {
	bar := &hospitalitymodels.Bar{
		ID:             uuid.New(),
		Name:           "Default Bar",
		OrganizationID: c.Org.ID,
		Organization:   c.Org,
		IsKitchen:      true,
		KitchenID:      &c.Kitchen.ID,
		OpeningTime:    time.Now(),
		ClosingTime:    time.Now(),
		CreatedAt:      time.Now(),
		UpdatedAt:      time.Now(),
	}

	if err := c.BarRepo.CreateBar(ctx, bar); err != nil {
		return err
	}

	menu := &hospitalitymodels.MenuGroup{
		ID:             uuid.New(),
		OrganizationID: c.Org.ID,
		Name:           "Bar Menu",
		Description:    "Default bar menu",
		Type:           "BAR",
		CreatedAt:      time.Now(),
		UpdatedAt:      time.Now(),
	}

	return c.MenuRepo.CreateMenuGroup(ctx, menu)
}
