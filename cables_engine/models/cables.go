// this cables main models
package cables_engine_models

import (
	cablesEngineConstants "sharpnr-engine/cables_engine/constants"
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

// A organization only gonna have one Cables Account
// the reason for this Account tables because one day we might need to Individual Cables Account

type OrgCablesAccount struct {
	ID             uuid.UUID
	OrganizationID uuid.UUID
	CreatedAt      time.Time
	IsActive       bool // this is switch which will or can turn-off the Cables service in the organization
}

func (t *OrgCablesAccount) TableName() string {
	return "cables.org_cables_accounts"
}

/* ---------------------------------- <***> --------------------------------- */

// CablesMember is the identity *inside* cables.
//
// Why not just use employee_id everywhere?
//   - a bot / integration posts messages but is not an employee
//   - an external guest can be in one channel without being an employee
//   - the chat display name, avatar and timezone are chat-owned, we don't
//     want to keep joining employee.employee on every message render
//
// So: one row per person-in-a-cables-account. Everything else in cables
// (participants, messages, reactions) points at member_id, never employee_id.
//
// This is also what makes anonymous meeting guests cheap: someone joining a
// public link gets an ephemeral GUEST row here (see CablesGuestSession), so the
// participant list, in-call chat and moderation all keep working on member_id
// with no nullable-member branch anywhere.
//
// indexes:
//
//	unique (org_cables_account_id, employee_id) where employee_id is not null
//	unique (org_cables_account_id, handle)      where handle is not null
//	index  (organization_id)
//	index  (guest_expires_at) where is_ephemeral -- the guest cleanup sweeper
type CablesMember struct {
	ID                 uuid.UUID
	OrganizationID     uuid.UUID
	OrgCablesAccountID uuid.UUID

	EmployeeID *uuid.UUID                       // employee.employee.id — null for BOT / external GUEST
	MemberType cablesEngineConstants.MemberType // EMPLOYEE | BOT | GUEST

	// chat identity — denormalized on purpose so the message list renders
	// without touching the employee tables
	DisplayName string
	Handle      *string // @handle, unique per cables account; null for meeting guests
	AvatarURL   *string
	Title       *string // "Backend Engineer" — shown in the profile popover
	TimeZone    *string // IANA, e.g. "Asia/Kolkata"

	// ---- ephemeral guests ----
	// True for a row minted at a public meeting link. These are NOT org members:
	// they must never show up in the member directory, in @mentions, or in the
	// DM picker — every such query filters is_ephemeral = false.
	IsEphemeral    bool
	GuestExpiresAt *time.Time // sweeper deletes the row (and its messages) after this

	IsActive bool // deactivated member: history stays, cannot post
	JoinedAt time.Time

	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt
}

func (t *CablesMember) TableName() string {
	return "cables.cables_members"
}

/* ---------------------------------- <***> --------------------------------- */

// Presence is written by the websocket layer (connect / disconnect / heartbeat),
// not by normal CRUD. One row per member, updated in place.
//
// Note: keep the *live* fan-out in redis. This table is the durable fallback so
// a restart doesn't show the whole org as offline, and so "last seen" survives.
type CablesMemberPresence struct {
	CablesMemberID uuid.UUID // PK, also FK to cables_members
	OrganizationID uuid.UUID

	Status            cablesEngineConstants.PresenceStatus
	ActiveConnections int // number of open sockets; 0 -> OFFLINE

	// the manual "🌴 on leave until friday" thing
	CustomStatusEmoji     *string
	CustomStatusText      *string
	CustomStatusExpiresAt *time.Time

	LastSeenAt time.Time
	UpdatedAt  time.Time
}

func (t *CablesMemberPresence) TableName() string {
	return "cables.cables_member_presence"
}

/* ---------------------------------- <***> --------------------------------- */

// Account-wide preferences for one member. Per-conversation overrides live on
// CablesParticipant (mute this one channel, etc.) and win over these.
type CablesMemberSettings struct {
	CablesMemberID uuid.UUID // PK
	OrganizationID uuid.UUID

	DefaultNotificationLevel cablesEngineConstants.NotificationLevel

	// do-not-disturb schedule, minutes from midnight in the member's timezone
	// (nil,nil = no schedule). Kept as minutes so it's timezone-shift proof.
	DNDStartMinute *int
	DNDEndMinute   *int
	DNDUntil       *time.Time // one-off "pause notifications for 2 hours"

	EmailOnMention   bool // fall back to email if they were offline
	PushOnMention    bool
	ShowReadReceipts bool

	CreatedAt time.Time
	UpdatedAt time.Time
}

func (t *CablesMemberSettings) TableName() string {
	return "cables.cables_member_settings"
}

/* ---------------------------------- <***> --------------------------------- */
