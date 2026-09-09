package stockwaveEngineServices

import (
	"context"
	stockwaveEngineDtos "sharpnr-engine/stockwave_engine/dtos"

	"github.com/google/uuid"
)

// StockwaveEngineServicesInterface is the stockwave engine's public surface.
//
// It carries the account layer only. Every method it used to have read or wrote
// the inventory tables migration 000062 dropped; the stock flows that replace
// them join this interface as they are designed, all gated on the account
// below.
type StockwaveEngineServicesInterface interface {
	// StartStockwaveServiceForOrg provisions the org's one stockwave account
	// together with its settings row.
	StartStockwaveServiceForOrg(ctx context.Context, orgID, emplID uuid.UUID, req stockwaveEngineDtos.StartStockwaveEngineRequestDto) error

	// GetStockwaveAccountStatus is the read every other stockwave flow is gated
	// on: has this org started stockwave, and is that account switched on.
	GetStockwaveAccountStatus(ctx context.Context, orgID, emplID uuid.UUID) (stockwaveEngineDtos.StockwaveEngineStatusResponseDto, error)
}
