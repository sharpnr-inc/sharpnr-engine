package cablesEngineDtos

import (
	"time"

	"github.com/google/uuid"
)

// CreateChannelRequestDto is the payload for creating a channel.
//
// The slug is deliberately not accepted from the caller: it is derived from the
// name, so a channel's @mention key cannot be squatted independently of the
// name people actually see.
type CreateChannelRequestDto struct {
	Name        string `json:"name" binding:"required,max=150"`
	Description string `json:"description"`
	Topic       string `json:"topic"`

	// Type defaults to TEXT and Visibility to PUBLIC when omitted — see the
	// Default* constants. PRIVATE is never implicit.
	Type       string `json:"type" binding:"omitempty,oneof=TEXT AUDIO VIDEO"`
	Visibility string `json:"visibility" binding:"omitempty,oneof=PUBLIC PRIVATE"`
}

// ChannelResponseDto is the wire shape of one channel.
//
// ConversationID is carried even though conversations do not exist yet: it is
// the handle the client will use to read and post messages, and shipping it
// from the start means the read path does not change shape when they land.
type ChannelResponseDto struct {
	ID             uuid.UUID `json:"id"`
	ConversationID uuid.UUID `json:"conversationId"`

	Name        string  `json:"name"`
	Slug        string  `json:"slug"`
	Topic       *string `json:"topic,omitempty"`
	Description *string `json:"description,omitempty"`

	Type       string `json:"type"`
	Visibility string `json:"visibility"`

	IsDefault   bool `json:"isDefault"`
	IsArchived  bool `json:"isArchived"`
	MemberCount int  `json:"memberCount"`

	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}
