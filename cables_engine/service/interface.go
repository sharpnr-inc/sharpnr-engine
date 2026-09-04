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

	// JoinChannel puts the caller into a public channel. Private channels are
	// not joinable — somebody already inside has to add you.
	JoinChannel(ctx context.Context, orgID, emplID, channelID uuid.UUID) (cablesEngineDtos.ConversationSummaryDto, error)

	/* ------------------------------- workspace ------------------------------- */

	// Bootstrap is the client's first call: identity, roster and sidebar in one
	// round trip. It also reconciles the org roster into cables_members, so an
	// employee hired since the last visit gets their cables identity here.
	Bootstrap(ctx context.Context, orgID, emplID uuid.UUID) (cablesEngineDtos.BootstrapResponseDto, error)

	ListConversations(ctx context.Context, orgID, emplID uuid.UUID) ([]cablesEngineDtos.ConversationSummaryDto, error)

	// OpenDirectConversation gets or creates the DM with one other member.
	// Idempotent — the same pair always resolves to the same conversation.
	OpenDirectConversation(ctx context.Context, orgID, emplID, otherMemberID uuid.UUID) (cablesEngineDtos.ConversationSummaryDto, error)

	/* -------------------------------- messages ------------------------------- */

	// GetMessages returns one page of history newest-first. beforeSeq is a
	// cursor rather than an offset, because history grows at the end being read.
	GetMessages(ctx context.Context, orgID, emplID, conversationID uuid.UUID, beforeSeq int64, limit int) (cablesEngineDtos.MessagePageResponseDto, error)

	// SendMessage writes the message and then fans it out. Writes go over HTTP,
	// never up the socket — this is where membership and idempotency are
	// enforced.
	SendMessage(ctx context.Context, orgID, emplID, conversationID uuid.UUID, request cablesEngineDtos.SendMessageRequestDto) (cablesEngineDtos.MessageResponseDto, error)

	// MarkRead moves the caller's read cursor, clearing the sidebar badge.
	MarkRead(ctx context.Context, orgID, emplID, conversationID uuid.UUID, seq int64) error
}
