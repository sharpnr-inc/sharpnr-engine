package siloEngineService

import (
	"context"
	siloEngineDtos "sharpnr-engine/silo_engine/dtos"

	"github.com/google/uuid"
)

type SiloEngineServiceInterface interface {
	// StartSiloServiceForOrg provisions the org's one Silo account together with
	// its settings row.
	StartSiloServiceForOrg(ctx context.Context, orgID, emplID uuid.UUID, req siloEngineDtos.StartSiloServiceRequestDto) error

	// GetSiloAccountStatus is the read every other Silo flow is gated on: has
	// this org started Silo, and is that account switched on.
	GetSiloAccountStatus(ctx context.Context, orgID, emplID uuid.UUID) (siloEngineDtos.SiloAccountStatusResponse, error)
}
