package cables_engine_models

import (
	"time"

	"github.com/google/uuid"
)

// Conversation is a chat thread. It groups participants and messages and acts
// as the SSE topic boundary (one stream per conversation: "chat_conv_<id>").
type Conversation struct {
	ID             uuid.UUID `gorm:"type:uuid;primaryKey"`
	OrganizationID uuid.UUID `gorm:"type:uuid;not null;index"`

	Type ConversationType `gorm:"size:20;not null;default:'DIRECT'"`

	// Name / Description / Avatar apply to GROUP conversations; for DIRECT
	// threads the client renders the other participant instead.
	Name        string `gorm:"size:150"`
	Description  string `gorm:"size:500"`
	AvatarURL    string `gorm:"size:500"`

	CreatedBy uuid.UUID `gorm:"type:uuid;not null"`

	// DirectKey is a deterministic, sorted hash of the two participant IDs for
	// DIRECT threads. A unique index on it prevents duplicate 1:1 conversations
	// between the same pair. Empty for GROUP conversations.
	DirectKey string `gorm:"size:100;index"`

	// Denormalised pointers to the latest message, used to sort and preview the
	// conversation list cheaply without aggregating the messages table.
	LastMessageID   *uuid.UUID `gorm:"type:uuid"`
	LastMessageText string     `gorm:"size:500"`
	LastMessageAt   *time.Time

	CreatedAt time.Time
	UpdatedAt time.Time

	Participants []ConversationParticipant `gorm:"foreignKey:ConversationID"`
}

func (Conversation) TableName() string {
	return "cables.conversation"
}
