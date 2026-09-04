package cablesEngineDtos

import (
	"time"

	"github.com/google/uuid"
)

/* -------------------------------------------------------------------------- */
/*                                  Bootstrap                                 */
/* -------------------------------------------------------------------------- */

// BootstrapResponseDto is everything the cables shell needs to render, in one
// round trip.
//
// It is one endpoint rather than four because the sidebar cannot draw until it
// has conversations AND the roster (every author name and DM title resolves
// against the roster locally), and the app's route guard already spends three
// sequential requests before this one runs. Messages are deliberately NOT here:
// they load per conversation, so this payload stays the same size whether the
// org has a day of history or a year.
type BootstrapResponseDto struct {
	Me            MemberResponseDto        `json:"me"`
	Account       CablesServiceResponseDto `json:"account"`
	Conversations []ConversationSummaryDto `json:"conversations"`
	Members       []MemberResponseDto      `json:"members"`
}

/* -------------------------------------------------------------------------- */
/*                                   Members                                  */
/* -------------------------------------------------------------------------- */

// MemberResponseDto is a cables identity — an employee, a bot, or a guest.
// EmployeeID is null for the latter two, which is exactly why messages address
// members rather than employees.
type MemberResponseDto struct {
	ID          uuid.UUID  `json:"id"`
	EmployeeID  *uuid.UUID `json:"employeeId,omitempty"`
	Type        string     `json:"type"`
	DisplayName string     `json:"displayName"`
	AvatarURL   *string    `json:"avatarUrl,omitempty"`
	Presence    string     `json:"presence"`
}

/* -------------------------------------------------------------------------- */
/*                                Conversations                               */
/* -------------------------------------------------------------------------- */

// ConversationSummaryDto is one row of the sidebar.
//
// UnreadCount is computed server-side from (conversation.last_seq -
// member.last_read_seq) so the client never counts messages. A DIRECT
// conversation carries OtherMemberID instead of a name — the client renders the
// other person from the roster it already holds.
type ConversationSummaryDto struct {
	ID   uuid.UUID `json:"id"`
	Type string    `json:"type"` // CHANNEL | DIRECT | GROUP | MEETING

	// Channel identity, present only for Type == CHANNEL.
	ChannelID   *uuid.UUID `json:"channelId,omitempty"`
	Name        *string    `json:"name,omitempty"`
	Slug        *string    `json:"slug,omitempty"`
	ChannelType *string    `json:"channelType,omitempty"`

	// The other party, present only for Type == DIRECT.
	OtherMemberID *uuid.UUID `json:"otherMemberId,omitempty"`

	LastSeq          int64      `json:"lastSeq"`
	LastMessageAt    *time.Time `json:"lastMessageAt,omitempty"`
	LastMessageBrief *string    `json:"lastMessageBrief,omitempty"`

	UnreadCount    int64 `json:"unreadCount"`
	UnreadMentions int   `json:"unreadMentions"`
	IsMuted        bool  `json:"isMuted"`
	IsPinned       bool  `json:"isPinned"`

	MemberCount int `json:"memberCount"`
}

// OpenDirectConversationRequestDto opens (or reopens) a DM with one other
// member. Idempotent — the same pair always resolves to the same conversation.
type OpenDirectConversationRequestDto struct {
	MemberID uuid.UUID `json:"memberId" binding:"required"`
}

/* -------------------------------------------------------------------------- */
/*                                  Messages                                  */
/* -------------------------------------------------------------------------- */

type MessageResponseDto struct {
	ID             uuid.UUID  `json:"id"`
	ConversationID uuid.UUID  `json:"conversationId"`
	Seq            int64      `json:"seq"`
	SenderMemberID *uuid.UUID `json:"senderMemberId,omitempty"`

	Type    string  `json:"type"`
	SubType *string `json:"subType,omitempty"`
	Body    string  `json:"body"`

	// Echoed back so the sender can reconcile its optimistic bubble with the
	// stored row instead of rendering the message twice.
	ClientMessageID *string `json:"clientMessageId,omitempty"`

	IsEdited  bool       `json:"isEdited"`
	IsDeleted bool       `json:"isDeleted"`
	CreatedAt time.Time  `json:"createdAt"`
	EditedAt  *time.Time `json:"editedAt,omitempty"`
}

// MessagePageResponseDto is one page of history, newest first.
//
// NextCursor is the seq to pass as ?before= for the previous page; it is null
// once the top of the conversation is reached, which is how the client knows to
// stop asking.
type MessagePageResponseDto struct {
	Messages   []MessageResponseDto `json:"messages"`
	HasMore    bool                 `json:"hasMore"`
	NextCursor *int64               `json:"nextCursor,omitempty"`
}

// SendMessageRequestDto posts one message.
//
// ClientMessageID is generated by the client before the request. It is what
// makes a retry after a dropped connection safe: the second POST resolves to
// the row the first one wrote rather than double-posting.
type SendMessageRequestDto struct {
	Body            string `json:"body" binding:"required,max=8000"`
	ClientMessageID string `json:"clientMessageId"`
}

// MarkReadRequestDto moves this member's read cursor. Seq never travels
// backwards server-side, so two tabs reporting out of order is harmless.
type MarkReadRequestDto struct {
	Seq int64 `json:"seq"`
}
