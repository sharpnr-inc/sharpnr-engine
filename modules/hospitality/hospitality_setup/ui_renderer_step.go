package hospitality_setup

import (
	"context"
	"time"

	"sharpnr-engine/core/models"
	sharpnruirenderer "sharpnr-engine/sharpnr_ui_renderer"

	"github.com/google/uuid"
)

type UIRendererSetupStep struct{}

func (s *UIRendererSetupStep) Name() string { return "UI_RENDERER" }

func (s *UIRendererSetupStep) Execute(ctx context.Context, c SetupContext) error {

	rendererID := uuid.New()

	uiTree, err := sharpnruirenderer.BarHotelRestaurantFactory(
		ctx,
		c.Org.ID,
		c.Org.OrgOperationType,
		rendererID,
	)
	if err != nil {
		return err
	}

	tree := &models.OrgUIRendererTree{
		ID:             rendererID,
		UIRendererTree: uiTree,
		OrganizationID: c.Org.ID,
		CreatedAt:      time.Now(),
		UpdatedAt:      time.Now(),
	}

	return c.UIRenderRepo.CreateOrgUIRenderer(ctx, tree)
}
