package uisetup

// import (
// 	"context"
// 	"sharpnr-engine/core/models"
// 	"time"

// 	"github.com/google/uuid"
// )

// type HospitalityUIRendererStep struct{}

// func (s *HospitalityUIRendererStep) Name() string {
// 	return "HOSPITALITY_UI_RENDERER"
// }

// func (s *HospitalityUIRendererStep) Execute(ctx context.Context, c SetupContext) error {

// 	rendererID := uuid.New()

// 	uiTree, err := SetupContext.(
// 		ctx,
// 		c.Org.ID,
// 		c.Modules,
// 		rendererID,
// 	)
// 	if err != nil {
// 		return err
// 	}

// 	tree := &models.OrgUIRendererTree{
// 		ID:             rendererID,
// 		UIRendererTree: uiTree,
// 		OrganizationID: c.Org.ID,
// 		CreatedAt:      time.Now(),
// 		UpdatedAt:      time.Now(),
// 	}

// 	return c.UIRenderRepo.CreateOrgUIRenderer(ctx, tree)
// }
