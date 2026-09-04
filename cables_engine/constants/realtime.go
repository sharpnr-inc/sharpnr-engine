package cablesEngineConstants

/* -------------------------------------------------------------------------- */
/*                              Realtime topics                               */
/* -------------------------------------------------------------------------- */

// Topic namespaces cables claims on the websocket engine.
//
// The engine denies any namespace nobody registered, so these constants and the
// RegisterAuthorizer calls at bootstrap are the only thing that brings a cables
// topic into existence.
const (
	// TopicNamespaceConversation carries everything that happens inside one
	// conversation — messages, edits, typing, call badges. Subscribing requires
	// membership.
	TopicNamespaceConversation = "conv"

	// TopicNamespaceMember is a member's private channel: DM invites, mention
	// pings, anything addressed to them rather than to a room. Subscribing
	// requires being that member.
	TopicNamespaceMember = "member"
)

// Event names carried on those topics. The client branches on these, so they
// are part of the API contract as much as any JSON field.
const (
	EventMessageCreated = "message.created"
	EventMessageUpdated = "message.updated"
	EventMessageDeleted = "message.deleted"

	// EventConversationOpened reaches a member's own topic when someone starts
	// a DM with them, so the sidebar grows a row without a refetch.
	EventConversationOpened = "conversation.opened"

	EventPresenceChanged = "presence.changed"
)

/* -------------------------------------------------------------------------- */
/*                                  Messages                                  */
/* -------------------------------------------------------------------------- */

const (
	// MaxMessageBodyLength matches the binding tag on SendMessageRequestDto.
	// Re-checked in the service because the service is also reachable from the
	// orchestrator, which does not run that tag.
	MaxMessageBodyLength = 8000

	// DefaultMessagePageSize is one screen of history plus room to scroll.
	DefaultMessagePageSize = 50
	MaxMessagePageSize     = 100

	// MaxMessageBriefLength bounds the sidebar preview stored on the
	// conversation row.
	MaxMessageBriefLength = 140
)
