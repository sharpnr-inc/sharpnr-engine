package cables_engine_models

import (
	cablesEngineConstants "sharpnr-engine/cables_engine/constants"
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type CablesCall struct {
	ID             uuid.UUID
	OrganizationID uuid.UUID

	// Always set. For a channel call it's the channel's conversation, for a DM
	// call the DM, for a meeting the meeting's own conversation. This is why
	// in-call chat needs no new tables.
	ConversationID uuid.UUID

	Origin cablesEngineConstants.CallOrigin // CHANNEL | DIRECT | MEETING
	Type   cablesEngineConstants.CallType   // AUDIO | VIDEO
	Status cablesEngineConstants.CallStatus

	// set only when Origin == MEETING
	CablesMeetingID *uuid.UUID
	// convenience backlink for channel calls; also lets a meeting call report
	// itself into the channel it belongs to
	CablesChannelID *uuid.UUID

	// ---- media provider handles ----
	Provider       *string // "livekit" — so a migration can tell old rows apart
	ProviderRoomID *string // room to rejoin on reconnect
	ProviderRegion *string

	StartedBy uuid.UUID // cables_members.id (a guest can never start one)
	StartedAt time.Time
	EndedAt   *time.Time
	EndedBy   *uuid.UUID
	EndReason *string // "host_ended" | "everyone_left" | "timeout" | "provider_error"

	DurationSeconds *int // filled on end; billing/reporting reads this, not the diff

	// ---- live counters, updated as people come and go ----
	ParticipantCount     int // currently in the call
	PeakParticipantCount int
	TotalJoinCount       int // includes rejoins; the number that matters for abuse
	GuestCount           int // anonymous joiners currently in — surfaced in the UI

	// ---- recording ----
	IsRecording    bool
	HasRecording   bool
	RecordingCount int

	// The "call started / call ended · 12 min" card in the message list.
	// Same message row is updated on end rather than posting a second one.
	MessageID *uuid.UUID

	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt
}

func (t *CablesCall) TableName() string {
	return "cables_engine.cables_calls"
}

/* ---------------------------------- <***> --------------------------------- */

// One row per join attempt — NOT one row per person. Somebody who drops on the
// train and rejoins twice has three rows, which is exactly what you want for
// "why did the recording cut out" and for honest talk-time numbers.
//
// Guests are here too, on the ephemeral GUEST member row minted by
// CablesGuestSession, so this table never needs a nullable member id or a
// parallel guest table.
//
// indexes:
//
//	index (call_id, status)
//	index (call_id) where status in ('JOINED','LOBBY')  -- the live roster
//	index (cables_member_id, joined_at DESC)            -- "my call history"
type CablesCallParticipant struct {
	ID             uuid.UUID
	OrganizationID uuid.UUID
	CallID         uuid.UUID
	CablesMemberID uuid.UUID

	Role   cablesEngineConstants.CallParticipantRole
	Status cablesEngineConstants.CallParticipantStatus

	// ---- how they got in ----
	// Set when they came off a public/shared link. Null = internal member who
	// joined from the app. This single field answers "was this person external",
	// which the UI needs on every tile.
	GuestSessionID *uuid.UUID
	MeetingLinkID  *uuid.UUID

	// lobby handling
	KnockedAt  *time.Time
	AdmittedAt *time.Time
	AdmittedBy *uuid.UUID

	JoinedAt        *time.Time // null while still in the lobby
	LeftAt          *time.Time
	DurationSeconds *int

	RemovedBy     *uuid.UUID
	RemovalReason *string

	// ---- live media state, updated by provider webhooks ----
	IsMuted         bool
	IsVideoOn       bool
	IsScreenSharing bool
	IsHandRaised    bool
	HandRaisedAt    *time.Time

	// ---- client info, useful when someone reports "it was choppy" ----
	DeviceType        *string // web | ios | android | desktop
	Platform          *string
	AppVersion        *string
	IPAddress         *string
	ConnectionQuality *string // last reported: good | poor | lost

	CreatedAt time.Time
	UpdatedAt time.Time
}

func (t *CablesCallParticipant) TableName() string {
	return "cables_engine.cables_call_participants"
}

/* ---------------------------------- <***> --------------------------------- */

// Append-only moderation and lifecycle trail for a call.
//
// Worth its own table rather than folding into audit_logs.logs: guests have no
// employee id, these rows are written at socket rate (hand raised, mute, share),
// and they get pruned with the call. Anything an org auditor should see —
// "a public link was created", "the meeting was recorded" — should ALSO go to
// audit_logs.logs through the normal audit service.
//
// indexes: index (call_id, created_at)
type CablesCallEvent struct {
	ID             uuid.UUID
	OrganizationID uuid.UUID
	CallID         uuid.UUID

	EventType cablesEngineConstants.CallEventType

	// who it happened to, and who did it (host actions)
	CablesMemberID *uuid.UUID
	ActorMemberID  *uuid.UUID

	Metadata *string // small JSON blob: old/new role, deny reason, ...

	CreatedAt time.Time
}

func (t *CablesCallEvent) TableName() string {
	return "cables_engine.cables_call_events"
}

/* ---------------------------------- <***> --------------------------------- */

// A recording produced by a call. Separate rows for the video file and the
// transcript, because they finish at different times and are deleted under
// different retention rules.
//
// Consent: a recording of a call that outsiders joined is sensitive. Keep
// ConsentCapturedAt honest — the join gate writes it — and let retention
// actually delete things.
//
// indexes:
//
//	index (call_id)
//	index (organization_id, created_at DESC)
//	index (expires_at) where status = 'READY'   -- the retention sweeper
type CablesCallRecording struct {
	ID             uuid.UUID
	OrganizationID uuid.UUID
	CallID         uuid.UUID
	ConversationID uuid.UUID // so the "Files"/recordings tab is one indexed read

	Kind   cablesEngineConstants.RecordingKind
	Status cablesEngineConstants.RecordingStatus

	// object storage — never a signed url, those are minted per request
	StorageKey      *string
	ThumbnailKey    *string
	TranscriptKey   *string
	MimeType        *string
	SizeBytes       *int64
	DurationSeconds *int

	ProviderEgressID *string // provider's job id, for polling/webhook reconciliation
	FailureReason    *string

	StartedBy uuid.UUID
	StartedAt time.Time
	EndedAt   *time.Time

	ConsentCapturedAt *time.Time

	// null = keep until someone deletes it; otherwise the sweeper hard-deletes
	ExpiresAt *time.Time

	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt
}

func (t *CablesCallRecording) TableName() string {
	return "cables_engine.cables_call_recordings"
}

/* ---------------------------------- <***> --------------------------------- */
