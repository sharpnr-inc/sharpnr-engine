package cables_engine_models

import (
	"time"

	"github.com/google/uuid"
)

// ConversationParticipant is a user's membership in a conversation. It also
// carries the per-user read state that drives unread counts and read receipts.
type ConversationParticipant struct {
	ID             uuid.UUID `gorm:"type:uuid;primaryKey"`
	ConversationID uuid.UUID `gorm:"type:uuid;not null;index:idx_participant_conv_user,unique"`
	UserID         uuid.UUID `gorm:"type:uuid;not null;index:idx_participant_conv_user,unique"`

	OrganizationID uuid.UUID `gorm:"type:uuid;not null;index"`

	Role ParticipantRole `gorm:"size:20;not null;default:'MEMBER'"`

	// LastRead* track how far this user has read. Unread = messages newer than
	// LastReadAt. Updated when the client acks delivery over the stream / REST.
	LastReadMessageID *uuid.UUID `gorm:"type:uuid"`
	LastReadAt        *time.Time

	// MutedUntil suppresses push/notification fan-out without leaving the thread.
	MutedUntil *time.Time

	JoinedAt time.Time
	// LeftAt marks a soft departure; the row is kept so history stays attributable.
	LeftAt *time.Time

	CreatedAt time.Time
	UpdatedAt time.Time
}

func (ConversationParticipant) TableName() string {
	return "cables.conversation_participant"
}
