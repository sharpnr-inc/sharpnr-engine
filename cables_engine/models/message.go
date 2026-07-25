package cables_engine_models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

// Message is a single chat message and the durable source of truth that the SSE
// engine streams a copy of. Sequence gives a per-conversation monotonic ordering
// so clients can resume via SSE Last-Event-ID and detect gaps.
type Message struct {
	ID             uuid.UUID `gorm:"type:uuid;primaryKey"`
	ConversationID uuid.UUID `gorm:"type:uuid;not null;index:idx_message_conv_seq,priority:1"`
	OrganizationID uuid.UUID `gorm:"type:uuid;not null;index"`

	// SenderID is null for SYSTEM messages.
	SenderID *uuid.UUID `gorm:"type:uuid"`

	Type MessageType `gorm:"size:20;not null;default:'TEXT'"`
	Body string      `gorm:"type:text"`

	// Sequence is assigned per conversation (1,2,3…) and is the stable ordering
	// key for streaming/replay. Indexed together with ConversationID.
	Sequence int64 `gorm:"not null;index:idx_message_conv_seq,priority:2"`

	// ReplyToMessageID threads a reply to an earlier message in the same convo.
	ReplyToMessageID *uuid.UUID `gorm:"type:uuid"`

	// SearchVector is a generated tsvector populated in the DB (trigger/generated
	// column defined in the migration). Read-only here ("->") so GORM never tries
	// to write it; a GIN index over it powers full-text search.
	SearchVector string `gorm:"->;type:tsvector" json:"-"`

	EditedAt  *time.Time
	CreatedAt time.Time
	UpdatedAt time.Time
	// Soft delete keeps the row so threads/replies stay consistent ("message deleted").
	DeletedAt gorm.DeletedAt `gorm:"index"`

	Attachments []MessageAttachment `gorm:"foreignKey:MessageID"`
	Reactions   []MessageReaction   `gorm:"foreignKey:MessageID"`
}

func (Message) TableName() string {
	return "cables.message"
}

// MessageAttachment is a file/image bound to a message. The binary itself lives
// in object storage; only metadata and the URL are kept here.
type MessageAttachment struct {
	ID        uuid.UUID `gorm:"type:uuid;primaryKey"`
	MessageID uuid.UUID `gorm:"type:uuid;not null;index"`

	URL      string `gorm:"size:1000;not null"`
	FileName string `gorm:"size:255"`
	MimeType string `gorm:"size:150"`
	SizeBytes int64
	// Width/Height let the client reserve layout space for images before load.
	Width  int
	Height int

	CreatedAt time.Time
}

func (MessageAttachment) TableName() string {
	return "cables.message_attachment"
}

// MessageReaction is a single user's emoji reaction to a message. The unique
// index keeps one reaction per (message, user, emoji).
type MessageReaction struct {
	ID        uuid.UUID `gorm:"type:uuid;primaryKey"`
	MessageID uuid.UUID `gorm:"type:uuid;not null;index:idx_reaction_unique,unique,priority:1"`
	UserID    uuid.UUID `gorm:"type:uuid;not null;index:idx_reaction_unique,unique,priority:2"`
	Emoji     string    `gorm:"size:32;not null;index:idx_reaction_unique,unique,priority:3"`

	CreatedAt time.Time
}

func (MessageReaction) TableName() string {
	return "cables.message_reaction"
}
