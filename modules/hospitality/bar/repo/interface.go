package barrepo

import (
	"context"
	hospitalitymodels "sharpnr-engine/modules/hospitality/hospitality_models"

	"github.com/google/uuid"
)

type BarRepository interface {
	CreateBar(ctx context.Context, bar *hospitalitymodels.Bar) error

	GetBar(ctx context.Context, barId uuid.UUID, orgId uuid.UUID) (*hospitalitymodels.Bar, error)

	GetBarByOrgId(ctx context.Context, orgId uuid.UUID) (*hospitalitymodels.Bar, error)

	CreateFloor(ctx context.Context, floor *hospitalitymodels.BarFloor) error

	UpdateBarDetails(ctx context.Context, barID uuid.UUID, updates map[string]interface{}) error

	CreateTable(ctx context.Context, batTable *hospitalitymodels.BarTable) error

	GetFloorById(ctx context.Context, floorId uuid.UUID) (*hospitalitymodels.BarFloor, error)
}
