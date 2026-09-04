package cables_engine_models

import (
	cablesEngineConstants "sharpnr-engine/cables_engine/constants"
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

// CablesConversationMember = "this member is in this conversation".
//
// Keyed on the CONVERSATION, not on the channel. A channel owns exactly one
// conversation (1:1, enforced by the unique index on cables_channels), so for a
// channel this is the same row set either way — but keying it here means a DM
// and a group DM use the same table, the same membership checks, the same
// unread math and the same websocket fan-out as a channel, instead of a second
// parallel stack.
//
// It carries the read state as well as the membership because every client
// sync ("give me my sidebar with unread badges") is then one indexed read for
// one member with no joins.
//
// indexes:
//
//	unique (conversation_id, cables_member_id)
//	index  (cables_member_id, is_hidden, left_at)   -- the sidebar query
//	index  (conversation_id) where left_at is null  -- member list / fan-out
type CablesConversationMember struct {
	ID             uuid.UUID
	OrganizationID uuid.UUID
	ConversationID uuid.UUID
	CablesMemberID uuid.UUID

	Role cablesEngineConstants.ParticipantRole // OWNER | ADMIN | MEMBER | GUEST

	// ---- read state ----
	// LastReadSeq is compared against conversation.last_seq, so unread is one
	// subtraction rather than a COUNT(*) over messages. A sequence rather than
	// a timestamp, so it stays exact when two messages land in the same
	// millisecond.
	LastReadSeq       int64
	LastReadMessageID *uuid.UUID
	LastReadAt        *time.Time

	// UnreadMentionCount cannot be derived from seq math — it depends on
	// whether THIS member was mentioned — so it is a real counter, incremented
	// when a mention row is written and zeroed on read.
	UnreadMentionCount int

	// ---- per-conversation preferences ----
	NotificationLevel cablesEngineConstants.NotificationLevel
	IsMuted           bool
	MutedUntil        *time.Time
	IsPinned          bool // favourited to the top of the sidebar

	// IsHidden keeps an opened-but-unused DM out of the other person's sidebar
	// until someone actually says something. Cleared for everyone on the first
	// send, and settable again by "close conversation".
	IsHidden bool

	// ---- lifecycle ----
	// Leaving is soft so old posts and system messages ("Ali left") still
	// resolve the member. Rejoining clears RemovedAt on the same row rather
	// than inserting a second one, which is what the unconditional unique index
	// above enforces.
	AddedAt   time.Time
	AddedBy   *uuid.UUID // null when they joined a public channel themselves
	RemovedAt *time.Time
	RemovedBy *uuid.UUID

	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt
}

func (CablesConversationMember) TableName() string {
	return "cables_engine.cables_conversation_members"
}

/* ---------------------------------- <***> --------------------------------- */

// Who gets notified about replies on one thread.
//
// A thread is not its own conversation (see CablesMessage.ParentMessageID) —
// it's just replies pointing at a root message. But notification targeting for
// a thread is different from the channel's: only people who posted in the
// thread or explicitly followed it should be pinged. Hence this small table.
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

func (CablesThreadFollower) TableName() string {
	return "cables_engine.cables_thread_followers"
}

/* ---------------------------------- <***> --------------------------------- */
