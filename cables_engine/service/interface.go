package cablesEngineService

import (
	"context"
	cablesEngineDtos "sharpnr-engine/cables_engine/dtos"

	"github.com/google/uuid"
)

type CablesEngineServiceInterface interface {
	// this will start the cables service
	StartCablesService(ctx context.Context, orgID, empID uuid.UUID, request cablesEngineDtos.StartCablesServiceRequestsDto) error
	GetCablesServiceDetails(ctx context.Context, orgID, emplID uuid.UUID) (cablesEngineDtos.CablesServiceResponseDto, error)
	AddNewChannel(ctx context.Context, orgID, emplID uuid.UUID)
}
