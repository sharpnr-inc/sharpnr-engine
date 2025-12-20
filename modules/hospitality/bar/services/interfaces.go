package barservices

import (
	"context"
	hospitalitydtos "sharpnr-engine/modules/hospitality/hospitality_dtos"

	"github.com/google/uuid"
)

type BarServices interface {
	GetBarLayout(ctx context.Context, orgId uuid.UUID) (*hospitalitydtos.GetBarDetails_RESPONSE, error)

	CreateTablesInBar(ctx context.Context, orgId uuid.UUID, req hospitalitydtos.Bar_CreateTablesInBar_Request) error
	GetBarDetailsConfig(ctx context.Context, orgId uuid.UUID) (*hospitalitydtos.BAR_UpdateBarDetails_Response, error)
	CreateFloor(ctx context.Context, orgId uuid.UUID, req hospitalitydtos.BAR_Add_Floor_Req) error

	UpdateBar(ctx context.Context, orgId uuid.UUID, req hospitalitydtos.BAR_BarUpdateReq) error
}
