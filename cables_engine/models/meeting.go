package cables_engine_models

import (
	cablesEngineConstants "sharpnr-engine/cables_engine/constants"
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

// CablesMeeting is a *named room* — instant ("start a meeting now") or
// scheduled ("standup, tomorrow 10am"), and it is the only thing that can hand
// out public join links.
//
// The three ways a call starts, and which of these tables they touch:
//
//  1. video/audio CHANNEL  -> cables_calls only. Channel members walk in and
//     out, discord style. No meeting row needed.
//  2. DM / group DM        -> cables_calls only. It rings the participants.
//  3. MEETING              -> cables_meetings + cables_meeting_links (+ calls).
//     This is the one outsiders can join.
//
// Every meeting owns a conversation (type MEETING). That is not decoration: it
// gives the meeting its chat, its file drops and its participant list for free,
// on the same tables as the rest of cables, instead of a parallel "meeting_chat".
//
// IsPersistent = the "my standing room" case: same link forever, many calls
// over its lifetime. Otherwise the room dies with the call.
//
// indexes:
//
//	unique (org_cables_account_id, room_code)
//	index  (organization_id, status, scheduled_start_at)
//	index  (host_member_id, scheduled_start_at DESC)
type CablesMeeting struct {
	ID                 uuid.UUID
	OrganizationID     uuid.UUID
	OrgCablesAccountID uuid.UUID

	// chat/files/participants container for this meeting (type MEETING)
	ConversationID uuid.UUID

	// If the meeting lives inside a channel ("the #design weekly"), keep the
	// backlink so the recording and the summary post land in that channel.
	CablesChannelID *uuid.UUID

	Title       string
	Description *string

	// RoomCode is the human-shareable part of the url — "cbl-xkqa-2m4".
	// Unique per cables account. Note this is NOT the credential: the token on
	// cables_meeting_links is. A guessed room code alone gets you nothing.
	RoomCode string

	Type   cablesEngineConstants.CallType // AUDIO | VIDEO
	Status cablesEngineConstants.MeetingStatus

	HostMemberID uuid.UUID // the owner; co-hosts are per-call

	// ---- scheduling (all null for an instant meeting) ----
	ScheduledStartAt *time.Time
	ScheduledEndAt   *time.Time
	TimeZone         *string // IANA — needed to render recurring meetings correctly
	RecurrenceRule   *string // RFC 5545 RRULE; null = one-off

	IsPersistent bool // reusable room vs one-shot

	// ---- who gets in ----
	JoinPolicy cablesEngineConstants.MeetingJoinPolicy

	// Only for ANYONE_WITH_CODE. Hashed, never stored raw — same rule as
	// employee_provisioning.
	PasscodeHash *string

	// Lobby/waiting room. Strongly recommended on for any meeting where
	// JoinPolicy allows outsiders, so a stranger with the link can't just
	// appear in the room.
	LobbyEnabled         bool
	GuestsMustBeAdmitted bool // even invited outsiders knock first

	// what an outsider is allowed to do once inside
	AllowGuestChat        bool
	AllowGuestScreenShare bool
	AllowGuestCamera      bool
	AllowGuestMic         bool

	// hard cap; also a cheap abuse brake on a public link
	MaxParticipants *int

	// ---- recording ----
	AutoRecord               bool
	RecordingConsentRequired bool // show the "this call is recorded" gate on join

	// The public link stops working after this. Null for a persistent room.
	ExpiresAt *time.Time

	CreatedBy uuid.UUID
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt
}

func (t *CablesMeeting) TableName() string {
	return "cables_engine.cables_meetings"
}

/* ---------------------------------- <***> --------------------------------- */

// A join link. A meeting can have several — that's the point: one link for
// staff, one public link for clients, revoke the public one without breaking
// the internal one.
//
// The raw token is shown once, at creation, and only its hash is stored. The
// url looks like:
//
//	https://<app>/cables/join/<room_code>?t=<raw_token>
//
// Redeeming a token does NOT put you in the call. It gets you a guest session
// (below), which then still faces the lobby if the meeting says so.
//
// indexes:
//
//	unique (token_hash)
//	index  (cables_meeting_id, is_revoked)
type CablesMeetingLink struct {
	ID              uuid.UUID
	OrganizationID  uuid.UUID
	CablesMeetingID uuid.UUID

	TokenHash string
	Label     *string // "client link", "internal" — so the host knows what they're revoking

	// The role this link hands out. A webinar link gives VIEWER; the co-host
	// link gives COHOST. Never let a link grant HOST.
	GrantsRole cablesEngineConstants.CallParticipantRole

	// True = no sharpnr account required. This is the flag that turns a link
	// into a public one, and the one to check twice in the join handler.
	AllowAnonymous bool

	// Force a name prompt before joining, so the tile isn't "Guest".
	RequireDisplayName bool
	RequireEmail       bool // soft identification for client-facing links

	MaxUses  *int // null = unlimited
	UseCount int

	ExpiresAt *time.Time
	IsRevoked bool
	RevokedAt *time.Time
	RevokedBy *uuid.UUID

	CreatedBy uuid.UUID
	CreatedAt time.Time
	UpdatedAt time.Time
}

func (t *CablesMeetingLink) TableName() string {
	return "cables_engine.cables_meeting_links"
}

/* ---------------------------------- <***> --------------------------------- */

// The identity of somebody who joined off a public link with no sharpnr account.
//
// Design note — this is the important one. An anonymous joiner still needs to
// appear in the participant list, send in-call chat, and be kickable. Rather
// than making cables_member_id nullable across half the schema, joining as a
// guest mints an *ephemeral* cables_members row (MemberType = GUEST,
// IsEphemeral = true) and this session row holds the credential + the audit
// trail for it. Everything downstream — participants, messages, reactions —
// keeps working on member_id and needs no special guest branch.
//
// The session token is scoped to ONE meeting and expires. It is not a login.
//
// indexes:
//
//	unique (token_hash)
//	index  (cables_meeting_id, created_at DESC)
//	index  (cables_member_id)
type CablesGuestSession struct {
	ID              uuid.UUID
	OrganizationID  uuid.UUID
	CablesMeetingID uuid.UUID
	MeetingLinkID   uuid.UUID // which link they came through

	// the ephemeral GUEST member created for them
	CablesMemberID uuid.UUID

	TokenHash string

	DisplayName string  // what they typed at the door
	Email       *string // only when the link demanded it; unverified, treat as a label

	// abuse / audit — a public link is an open door, so keep enough to ban
	IPAddress *string
	UserAgent *string
	Country   *string

	ExpiresAt time.Time
	RevokedAt *time.Time // host kicked them; token dies immediately

	CreatedAt  time.Time
	LastSeenAt *time.Time
}

func (t *CablesGuestSession) TableName() string {
	return "cables_engine.cables_guest_sessions"
}

/* ---------------------------------- <***> --------------------------------- */

// Invitee list for a scheduled meeting. Internal people get a member id;
// outsiders are just an email we mail a link to.
//
// If org_calender_engine later owns scheduling, this table is what gets
// replaced — nothing else here depends on it.
//
// indexes:
//
//	unique (cables_meeting_id, cables_member_id) where cables_member_id is not null
//	unique (cables_meeting_id, email)            where email is not null
type CablesMeetingInvitee struct {
	ID              uuid.UUID
	OrganizationID  uuid.UUID
	CablesMeetingID uuid.UUID

	CablesMemberID *uuid.UUID // internal invitee
	Email          *string    // external invitee
	DisplayName    *string

	IsRequired bool
	Status     cablesEngineConstants.MeetingInviteeStatus

	// the personal link mailed to this invitee, so revoking one person's access
	// doesn't kill everyone else's
	MeetingLinkID *uuid.UUID

	InvitedBy   uuid.UUID
	RespondedAt *time.Time

	CreatedAt time.Time
	UpdatedAt time.Time
}

func (t *CablesMeetingInvitee) TableName() string {
	return "cables_engine.cables_meeting_invitees"
}

/* ---------------------------------- <***> --------------------------------- */
