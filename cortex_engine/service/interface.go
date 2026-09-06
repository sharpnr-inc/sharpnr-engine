package cortexEngineService

import (
	"context"
	cortexEngineDtos "sharpnr-engine/cortex_engine/dtos"

	"github.com/google/uuid"
)

type CortexEngineServiceInterface interface {
	// StartCortexServiceForOrg provisions the org's one Cortex account together
	// with its settings row.
	StartCortexServiceForOrg(ctx context.Context, orgID, emplID uuid.UUID, req cortexEngineDtos.StartCortexServiceRequestDto) error

	// GetCortexAccountStatus is the read every other Cortex flow is gated on:
	// has this org started Cortex, and is that account switched on.
	GetCortexAccountStatus(ctx context.Context, orgID, emplID uuid.UUID) (cortexEngineDtos.CortexAccountStatusResponse, error)
}
