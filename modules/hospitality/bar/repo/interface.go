package barrepo

import (
	"context"
	bar_models "sharpnr-engine/modules/hospitality/bar/models"

	"github.com/google/uuid"
)

type BarRepository interface {
	CreateBar(ctx context.Context, bar *bar_models.Bar) error
	GetBar(ctx context.Context, barId uuid.UUID, orgId uuid.UUID) (*bar_models.Bar, error)
	GetBarByOrgId(ctx context.Context, orgId uuid.UUID) (*bar_models.Bar, error)
	CreateFloor(ctx context.Context, floor *bar_models.BarFloor) error
	UpdateBarDetails(ctx context.Context, barID uuid.UUID, updates map[string]interface{}) error
	CreateTable(ctx context.Context, batTable *bar_models.BarTable) error
	GetFloorById(ctx context.Context, floorId uuid.UUID) (*bar_models.BarFloor, error)
}
