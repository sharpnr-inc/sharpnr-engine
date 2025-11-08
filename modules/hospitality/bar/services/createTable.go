package barservices

import (
	"context"
	hospitalitydtos "sharpnr-engine/modules/hospitality/hospitality_dtos"

	"github.com/google/uuid"
)

func (s *barServices) CreateTablesInBar(ctx context.Context, orgId uuid.UUID, barId uuid.UUID, floorId uuid.UUID) (*hospitalitydtos.BarTablesResponse, error) {
	panic("unimplemented")
}
