package barservices

import (
	"context"
	hospitalitydtos "sharpnr-engine/modules/hospitality/hospitality_dtos"

	"github.com/google/uuid"
)

type BarServices interface {
	InitializeDefaultBar()
	GetBarDetails(ctx context.Context, orgId uuid.UUID) (*hospitalitydtos.GetBarDetails_RESPONSE, error)
}
