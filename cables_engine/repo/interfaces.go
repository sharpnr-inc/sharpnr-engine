package cablesEngineRepo

import (
	"context"
	"time"

	cablesEngineConstants "sharpnr-engine/cables_engine/constants"
	cables_engine_models "sharpnr-engine/cables_engine/models"

	"github.com/google/uuid"
)

// CablesRepoInterface is the cables module's persistence port.
//
// Every method carries organizationID because the tenant predicate belongs on
// the query, not on the caller's memory: a channel id from another org must
// read as "not found" rather than resolve.
type CablesRepoInterface interface {
	// IntializeOrgCablesAccount mints the org's one cables account and returns
	// it. The unique index on organization_id is what makes a concurrent second
	// call fail rather than produce a duplicate.
	IntializeOrgCablesAccount(ctx context.Context, orgID uuid.UUID) (*cables_engine_models.OrgCablesAccount, error)

	// GetCablesAccountDetails returns the org's account, or
	// ErrCablesServiceNotInitialized when the org has never started cables.
	GetCablesAccountDetails(ctx context.Context, orgID uuid.UUID) (*cables_engine_models.OrgCablesAccount, error)

	// CablesAccountSwitch turns the org's cables service on or off. isActive is
	// an explicit target rather than a toggle so a retried request lands on the
	// state the caller asked for instead of flipping it back.
	CablesAccountSwitch(ctx context.Context, orgID, orgCablesAccountID uuid.UUID, isActive bool) error

	/* -------------------------------- channels ------------------------------- */

	CreateChannel(ctx context.Context, channel cables_engine_models.CablesChannel) error

	// ListChannels returns the account's channels in sidebar order.
	ListChannels(ctx context.Context, orgID, orgCablesAccountID uuid.UUID) ([]cables_engine_models.CablesChannel, error)

	GetChannelByID(ctx context.Context, orgID, channelID uuid.UUID) (*cables_engine_models.CablesChannel, error)

	// ChannelSlugExists reports whether a live channel in the account already
	// holds this slug. Soft-deleted rows do not count, matching the partial
	// unique index that enforces the same rule underneath.
	ChannelSlugExists(ctx context.Context, orgCablesAccountID uuid.UUID, slug string) (bool, error)

	/* --------------------------------- members -------------------------------- */

	// UpsertEmployeeMembers reconciles the org roster into cables_members. One
	// statement for the whole list; a member that already exists keeps its id
	// and presence and only has its display name refreshed.
	UpsertEmployeeMembers(ctx context.Context, members []cables_engine_models.CablesMember) error

	GetMemberByEmployee(ctx context.Context, orgCablesAccountID, employeeID uuid.UUID) (*cables_engine_models.CablesMember, error)

	// ListMembers returns the account's roster, excluding ephemeral guests.
	ListMembers(ctx context.Context, orgCablesAccountID uuid.UUID) ([]cables_engine_models.CablesMember, error)

	SetMemberPresence(ctx context.Context, memberID uuid.UUID, presence cablesEngineConstants.PresenceStatus) error

	/* ------------------------------ conversations ----------------------------- */

	CreateConversation(ctx context.Context, conversation *cables_engine_models.CablesConversation) error
	GetConversation(ctx context.Context, orgID, conversationID uuid.UUID) (*cables_engine_models.CablesConversation, error)

	// GetDirectConversation finds an existing DM by its direct_key, which is
	// what makes "open a DM with X" idempotent.
	GetDirectConversation(ctx context.Context, orgCablesAccountID uuid.UUID, directKey string) (*cables_engine_models.CablesConversation, error)

	// ListConversationsForMember is the sidebar: every conversation this member
	// is in, joined to their own read state, in display order.
	ListConversationsForMember(ctx context.Context, orgID, cablesMemberID uuid.UUID) ([]ConversationWithMembership, error)

	// AddConversationMember is an upsert: rejoining clears removed_at on the
	// existing row and keeps the read state.
	AddConversationMember(ctx context.Context, member *cables_engine_models.CablesConversationMember) error

	// GetConversationMember is the membership check behind every read, send and
	// websocket subscribe. Not a member reads as ErrNotAConversationMember.
	GetConversationMember(ctx context.Context, conversationID, cablesMemberID uuid.UUID) (*cables_engine_models.CablesConversationMember, error)

	ListConversationMembers(ctx context.Context, conversationID uuid.UUID) ([]cables_engine_models.CablesConversationMember, error)

	MarkConversationRead(ctx context.Context, conversationID, cablesMemberID uuid.UUID, seq int64) error

	// RevealConversation un-hides a conversation for everyone in it, so a DM
	// reaches the other person's sidebar on the first message rather than when
	// it was opened.
	RevealConversation(ctx context.Context, conversationID uuid.UUID) error

	/* -------------------------------- messages -------------------------------- */

	// NextConversationSeq allocates the next per-conversation sequence with one
	// atomic statement. Call it inside the send transaction.
	NextConversationSeq(ctx context.Context, conversationID uuid.UUID) (int64, error)

	CreateMessage(ctx context.Context, message *cables_engine_models.CablesMessage) error

	// GetMessageByClientID resolves a retried send to the row it already wrote.
	GetMessageByClientID(ctx context.Context, conversationID uuid.UUID, clientMessageID string) (*cables_engine_models.CablesMessage, error)

	// ListMessages returns one page of history newest-first. beforeSeq is a
	// cursor, not an offset; 0 means from the newest.
	ListMessages(ctx context.Context, conversationID uuid.UUID, beforeSeq int64, limit int) ([]cables_engine_models.CablesMessage, error)

	// TouchConversationAfterSend refreshes the denormalised sidebar columns.
	TouchConversationAfterSend(ctx context.Context, conversationID, messageID uuid.UUID, brief string, sentAt time.Time) error
}
