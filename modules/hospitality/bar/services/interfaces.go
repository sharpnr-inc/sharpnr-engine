package barservices

import (
	"context"
	hospitalitydtos "sharpnr-engine/modules/hospitality/hospitality_dtos"

	"github.com/google/uuid"
)

type BarServices interface {
	InitializeDefaultBar()
	GetBarDetails(ctx context.Context, orgId uuid.UUID) (*hospitalitydtos.GetBarDetails_RESPONSE, error)

	CreateTablesInBar(ctx context.Context, orgId uuid.UUID, barId uuid.UUID, floorId uuid.UUID) (*hospitalitydtos.BarTablesResponse, error)
	GetBarDetailsConfig(ctx context.Context, orgId uuid.UUID) (*hospitalitydtos.BAR_UpdateBarDetails_Response, error)
}
