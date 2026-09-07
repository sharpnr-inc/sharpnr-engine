package orchestrator_service

import (
	"context"

	stockwaveEngineDtos "sharpnr-engine/stockwave_engine/dtos"
	stockwaveEngineServices "sharpnr-engine/stockwave_engine/services"

	"github.com/google/uuid"
)

// stockwaveOrchestrator implements StockwaveOrchestratorServiceInterface.
//
// Formerly inventoryOrchestrator: the engine was renamed when migration 000062
// moved the `inventory` schema to `stockwave_engine` and dropped the tables the
// old methods read.
//
// Every method is a single delegated call, so there is no transaction here
// (rule 2 in doc.go) — the stockwave service owns its own consistency, and it
// already commits the account, its settings row and the audit entry describing
// them in one transaction.
//
// ⚠ No guard.Authorize calls yet (doc.go rule 1): stockwave has no Sentinel
// resource registered, so there is no action id to check against — guarding
// against an unregistered type would deny every request rather than protect it.
// Starting the org's stockwave service is currently open to any authenticated
// employee of the org; wire the guard here the day that registration lands, the
// way departmentOrchestrator does.
type stockwaveOrchestrator struct {
	stockwaveEngine stockwaveEngineServices.StockwaveEngineServicesInterface
}

/* ---------------------------------- <***> --------------------------------- */

func (s *stockwaveOrchestrator) GetStockwaveAccountStatus(ctx context.Context, orgID, emplID uuid.UUID) (stockwaveEngineDtos.StockwaveEngineStatusResponseDto, error) {
	return s.stockwaveEngine.GetStockwaveAccountStatus(ctx, orgID, emplID)
}

/* ---------------------------------- <***> --------------------------------- */

func (s *stockwaveOrchestrator) StartStockwaveService(ctx context.Context, orgID, emplID uuid.UUID, request stockwaveEngineDtos.StartStockwaveEngineRequestDto) error {
	return s.stockwaveEngine.StartStockwaveServiceForOrg(ctx, orgID, emplID, request)
}
