package cablesEngineService

import (
	"context"
	cablesEngineDtos "sharpnr-engine/cables_engine/dtos"

	"github.com/google/uuid"
)

// CablesEngineServiceInterface is the cables module's public surface.
//
// Everything below the account flows assumes a live cables account: a channel
// call against an org that never started cables, or switched it off, fails with
// ErrCablesServiceNotInitialized / ErrCablesAccountIsDisabled rather than
// quietly returning nothing.
type CablesEngineServiceInterface interface {
	// StartCablesService provisions the org's cables account. Starting an
	// account that exists but is switched off re-activates it; starting a live
	// one is an error, not a second account.
	StartCablesService(ctx context.Context, orgID, empID uuid.UUID, request cablesEngineDtos.StartCablesServiceRequestsDto) error

	// GetCablesServiceDetails is the status read behind the "is cables on for
	// us?" question. An org that never started cables is a false/false answer,
	// not an error — the client asks this before it knows.
	GetCablesServiceDetails(ctx context.Context, orgID, emplID uuid.UUID) (cablesEngineDtos.CablesServiceResponseDto, error)

	/* -------------------------------- channels ------------------------------- */

	AddNewChannel(ctx context.Context, orgID, emplID uuid.UUID, request cablesEngineDtos.CreateChannelRequestDto) error
	ListChannels(ctx context.Context, orgID, emplID uuid.UUID) ([]cablesEngineDtos.ChannelResponseDto, error)
	GetChannelDetails(ctx context.Context, orgID, emplID, channelID uuid.UUID) (cablesEngineDtos.ChannelResponseDto, error)
}
