package cables_engine_models

import (
	cablesEngineConstants "sharpnr-engine/cables_engine/constants"
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

// so a channel account gonna have multiple channels
// a organization can as much channels as they want
// Channel is kinda Group
//
// A channel owns exactly one conversation (see CablesConversation) — the
// channel row holds the *identity* of the group (name, topic, who can see it)
// and the conversation row holds the *messages*.
//
// indexes:
//
//	unique (org_cables_account_id, slug) where deleted_at is null
//	index  (organization_id, visibility, is_archived)
//	index  (channel_category_id)
type CablesChannel struct {
	ID                 uuid.UUID
	OrganizationID     uuid.UUID
	OrgCablesAccountID uuid.UUID

	// the message container. Created in the same transaction as the channel.
	ConversationID uuid.UUID

	Name        string  // display name: "Design Team"
	Slug        string  // url/mention safe: "design-team", unique per cables account
	Topic       *string // one-liner shown in the header
	Description *string

	// Type lives here and NOT on the settings table on purpose: every channel
	// list query filters/renders by it, and we don't want to join settings for
	// a sidebar. Settings stays for things you only read on the settings screen.
	Type       cablesEngineConstants.ChannelType       // TEXT | AUDIO | VIDEO
	Visibility cablesEngineConstants.ChannelVisibility // PUBLIC | PRIVATE

	// optional discord-style grouping in the sidebar
	ChannelCategoryID *uuid.UUID
	SortOrder         int

	IsDefault bool // everybody joining the org is auto-added (your #general)

	// Archived != deleted: the channel goes read-only, history stays searchable.
	// Active is your existing on/off switch; keep it for "temporarily disabled".
	IsArchived bool
	ArchivedAt *time.Time
	ArchivedBy *uuid.UUID

	MemberCount int // denormalized, kept in sync by the participant service

	CreatedAt time.Time
	Active    bool
	UpdatedAt time.Time
	CreatedBy *uuid.UUID
	DeletedAt gorm.DeletedAt
}

func (t *CablesChannel) TableName() string {
	return "cables.cables_channels"
}

/* ---------------------------------- <***> --------------------------------- */

// Sidebar grouping ("ENGINEERING", "SOCIAL"). Purely cosmetic, no permissions.
type CablesChannelCategory struct {
	ID                 uuid.UUID
	OrganizationID     uuid.UUID
	OrgCablesAccountID uuid.UUID

	Name      string
	SortOrder int

	CreatedAt time.Time
	UpdatedAt time.Time
	CreatedBy *uuid.UUID
	DeletedAt gorm.DeletedAt
}

func (t *CablesChannelCategory) TableName() string {
	return "cables.cables_channel_categories"
}

/* ---------------------------------- <***> --------------------------------- */

// Now we will create the Configuration Tables for Channels
// each channels gonna have their setting tables
// Note: For now this channel setting tables is kinda empty but i have added some features
// we add the stuff as we move along
//
// Rule of thumb for what belongs here vs on CablesChannel: if the message list
// or the sidebar needs it, put it on the channel; if only the settings screen
// and the send/permission check need it, put it here.
//
// indexes: unique (cables_channel_id)
type CablesChannelSettings struct {
	ID              uuid.UUID
	OrganizationID  uuid.UUID
	CablesChannelID uuid.UUID

	ChannelHeadline   *string
	ChannelThemeColor *string

	// ---- posting / moderation ----
	PostingPolicy    cablesEngineConstants.ChannelPostingPolicy // EVERYONE | ADMINS_ONLY
	SlowModeSeconds  int                                        // 0 = off; min gap between two posts by same member
	ThreadsEnabled   bool
	AllowGuests      bool
	AllowReactions   bool
	AllowAttachments bool

	// ---- calls ----
	// A TEXT channel can still allow a quick huddle; a VIDEO channel is a
	// standing room people drop into. AllowGuestLinks is the switch that lets
	// someone mint a public join link for a call in this channel — off by
	// default, because that is the one setting that exposes an org channel to
	// the outside world.
	AllowCalls          bool
	AllowGuestLinks     bool
	MaxCallParticipants *int
	AutoRecordCalls     bool

	// ---- retention ----
	// nil = keep forever. Otherwise a cleanup job hard-deletes messages older
	// than this many days.
	RetentionDays *int

	// default for anyone who joins; the member can override on their participant row
	DefaultNotificationLevel cablesEngineConstants.NotificationLevel

	CreatedAt time.Time
	UpdatedAt time.Time
}

func (t *CablesChannelSettings) TableName() string {
	return "cables.cables_channel_settings"
}

/* ---------------------------------- <***> --------------------------------- */

// Invite to a PRIVATE channel (or to the cables account itself for guests).
// Public channels don't need this — members just join.
//
// indexes: index (cables_channel_id, status), unique (token_hash)
type CablesChannelInvite struct {
	ID              uuid.UUID
	OrganizationID  uuid.UUID
	CablesChannelID uuid.UUID

	InvitedMemberID *uuid.UUID // existing member being pulled in
	InvitedEmail    *string    // external guest, no member row yet
	InvitedBy       uuid.UUID

	TokenHash string // same pattern as employee_provisioning — never store the raw token
	Status    string // PENDING | ACCEPTED | REVOKED | EXPIRED

	ExpiresAt  time.Time
	AcceptedAt *time.Time

	CreatedAt time.Time
	UpdatedAt time.Time
}

func (t *CablesChannelInvite) TableName() string {
	return "cables.cables_channel_invites"
}

/* ---------------------------------- <***> --------------------------------- */
