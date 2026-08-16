package cables_engine_models

import (
	cablesEngineConstants "sharpnr-engine/cables_engine/constants"
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

// CablesConversation is THE container every message belongs to.
//
// This is the one design call worth understanding before anything else:
// a channel is not a message container, a conversation is. A channel is just
// the *named, discoverable* face of a conversation.
//
//	channel #general  -> cables_channels row  -> conversation (type CHANNEL)
//	DM with Ali       -> no channel row       -> conversation (type DIRECT)
//	group DM of 4     -> no channel row       -> conversation (type GROUP)
//
// Everything downstream (messages, participants, unread state, reactions,
// pins, search) hangs off conversation_id only. One query path, one repo,
// instead of writing channel_messages and dm_messages twice.
//
// indexes:
//
//	index  (organization_id, type)
//	unique (channel_id) where channel_id is not null
//	unique (org_cables_account_id, direct_key) where direct_key is not null
type CablesConversation struct {
	ID                 uuid.UUID
	OrganizationID     uuid.UUID
	OrgCablesAccountID uuid.UUID

	Type      cablesEngineConstants.ConversationType // CHANNEL | DIRECT | GROUP | MEETING
	ChannelID *uuid.UUID                             // set only when Type == CHANNEL
	MeetingID *uuid.UUID                             // set only when Type == MEETING

	// DirectKey stops duplicate DMs. For a DIRECT conversation it is the two
	// member ids sorted and joined ("<uuidA>:<uuidB>"); for GROUP it's a hash
	// of the sorted member set at creation time. Unique per cables account, so
	// "open DM with X" is an upsert instead of a scan-for-existing.
	// Null for channels.
	DirectKey *string

	// GROUP DMs can be renamed; DIRECT ones take their title from the other
	// member, so this stays null there.
	Title *string

	// ---- hot counters, updated on every send ----
	// LastSeq is the sequence allocator for this conversation. Every message
	// takes LastSeq+1. Unread math is then just:
	//     unread = conversation.last_seq - participant.last_read_seq
	// which is one subtraction instead of a COUNT(*) over messages.
	LastSeq          int64
	LastMessageID    *uuid.UUID
	LastMessageAt    *time.Time // sidebar ordering
	LastMessageBrief *string    // "Ali: sounds good" — avoids a join per sidebar row
	MessageCount     int64
	ParticipantCount int

	// Live call badge in the sidebar ("🔴 3 in call") without querying
	// cables_calls for every row. Set on call start, cleared on call end.
	ActiveCallID *uuid.UUID

	CreatedBy *uuid.UUID // cables_members.id — null for system-created
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt
}

func (t *CablesConversation) TableName() string {
	return "cables.cables_conversations"
}

/* ---------------------------------- <***> --------------------------------- */
