package cables_engine_models

import (
	cablesEngineConstants "sharpnr-engine/cables_engine/constants"
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

// CablesParticipant = "this member is in this conversation".
//
// It is the membership row AND the per-member read state, deliberately in one
// table. Every client sync ("give me my sidebar with unread badges") is then a
// single indexed read of this table for one member, no joins.
//
// Works the same for a channel and for a DM, because both are conversations.
//
// indexes:
//
//	unique (conversation_id, cables_member_id)
//	index  (cables_member_id, is_hidden, last_message_at)  -- the sidebar query
//	index  (conversation_id) where left_at is null         -- member list / fan-out
type CablesParticipant struct {
	ID             uuid.UUID
	OrganizationID uuid.UUID
	ConversationID uuid.UUID
	CablesMemberID uuid.UUID

	Role cablesEngineConstants.ParticipantRole // OWNER | ADMIN | MEMBER | GUEST

	// ---- read state ----
	// LastReadSeq is compared against conversation.last_seq. Storing the seq
	// (not a timestamp) makes unread counting exact and monotonic even when two
	// messages land in the same millisecond.
	LastReadSeq       int64
	LastReadMessageID *uuid.UUID
	LastReadAt        *time.Time

	// UnreadMentionCount can't be derived from seq math (it depends on whether
	// *this* member was mentioned), so it's a real counter: incremented on
	// mention insert, zeroed on read.
	UnreadMentionCount int

	// ---- per-conversation preferences (override CablesMemberSettings) ----
	NotificationLevel cablesEngineConstants.NotificationLevel
	IsMuted           bool
	MutedUntil        *time.Time
	IsPinned          bool // favourited to the top of the sidebar
	IsHidden          bool // DM closed from the sidebar; reopens on next message

	// ---- lifecycle ----
	// Leaving is soft so system messages ("Ali left") and old posts still resolve
	// the member. Rejoining clears LeftAt rather than inserting a new row.
	JoinedAt  time.Time
	InvitedBy *uuid.UUID
	LeftAt    *time.Time

	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt
}

func (t *CablesParticipant) TableName() string {
	return "cables_engine.cables_participants"
}

/* ---------------------------------- <***> --------------------------------- */

// Who gets notified about replies on one thread.
//
// A thread is not its own conversation (see CablesMessage.ParentMessageID) —
// it's just replies pointing at a root message. But notification targeting for
// a thread is different from the channel: only people who posted in the thread
// or explicitly followed it should be pinged. Hence this small table.
//
// indexes: unique (root_message_id, cables_member_id)
type CablesThreadFollower struct {
	ID             uuid.UUID
	OrganizationID uuid.UUID
	ConversationID uuid.UUID
	RootMessageID  uuid.UUID
	CablesMemberID uuid.UUID

	// auto = they replied in the thread, manual = they hit "follow"
	IsAutoFollowed bool

	LastReadReplyAt *time.Time
	UnfollowedAt    *time.Time

	CreatedAt time.Time
	UpdatedAt time.Time
}

func (t *CablesThreadFollower) TableName() string {
	return "cables_engine.cables_thread_followers"
}

/* ---------------------------------- <***> --------------------------------- */
