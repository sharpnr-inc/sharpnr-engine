// Package cables_engine_models holds the persistence models for the chat /
// messaging module ("cables"). Messages are stored durably in Postgres (the
// source of truth and the searchable store via a tsvector column); ephemeral
// state such as presence/typing lives in Redis and is not modelled here.
package cables_engine_models

// ConversationType describes the shape of a conversation thread.
type ConversationType string

const (
	// ConversationDirect is a 1:1 thread between exactly two participants.
	ConversationDirect ConversationType = "DIRECT"
	// ConversationGroup is a named, multi-participant thread.
	ConversationGroup ConversationType = "GROUP"
)

// ParticipantRole controls what a participant may do inside a conversation.
type ParticipantRole string

const (
	ParticipantMember ParticipantRole = "MEMBER"
	ParticipantAdmin  ParticipantRole = "ADMIN"
	ParticipantOwner  ParticipantRole = "OWNER"
)

// MessageType distinguishes user content from system-generated events
// (e.g. "X added Y", "name changed") that are rendered inline in the thread.
type MessageType string

const (
	MessageText   MessageType = "TEXT"
	MessageImage  MessageType = "IMAGE"
	MessageFile   MessageType = "FILE"
	MessageSystem MessageType = "SYSTEM"
)
