package cablesEngineConstants

/* ---------------------------------- <***> --------------------------------- */
// Channels

type ChannelType string

const (
	ChannelTypeVideo ChannelType = "VIDEO"
	ChannelTypeText  ChannelType = "TEXT"
	ChannelTypeAudio ChannelType = "AUDIO"
)

// Who can see the channel and how someone gets in.
// PUBLIC  -> any member of the org's cables account can find it and join
// PRIVATE -> invite only, does not show up in the channel browser
type ChannelVisibility string

const (
	ChannelVisibilityPublic  ChannelVisibility = "PUBLIC"
	ChannelVisibilityPrivate ChannelVisibility = "PRIVATE"
)

// Who is allowed to post in the channel.
type ChannelPostingPolicy string

const (
	PostingPolicyEveryone   ChannelPostingPolicy = "EVERYONE"    // normal channel
	PostingPolicyAdminsOnly ChannelPostingPolicy = "ADMINS_ONLY" // announcements
)

/* ---------------------------------- <***> --------------------------------- */
// Conversations

// A conversation is the thing messages actually live in.
// CHANNEL -> backed by a cables_channels row (named, discoverable)
// DIRECT  -> 1:1 DM, exactly two participants
// GROUP   -> ad-hoc DM with 3+ people, no channel row
// MEETING -> backed by a cables_meetings row; exists so in-call chat, files and
//
//	guests reuse the exact same message pipeline as everything else
type ConversationType string

const (
	ConversationTypeChannel ConversationType = "CHANNEL"
	ConversationTypeDirect  ConversationType = "DIRECT"
	ConversationTypeGroup   ConversationType = "GROUP"
	ConversationTypeMeeting ConversationType = "MEETING"
)

/* ---------------------------------- <***> --------------------------------- */
// Members & participants

// What kind of identity is behind a cables member.
type MemberType string

const (
	MemberTypeEmployee MemberType = "EMPLOYEE" // linked to employee.employee
	MemberTypeBot      MemberType = "BOT"      // integrations posting into channels
	MemberTypeGuest    MemberType = "GUEST"    // external/limited access
)

// Role of a member *inside a single conversation*.
// This is conversation-local and has nothing to do with sentinel roles.
type ParticipantRole string

const (
	ParticipantRoleOwner  ParticipantRole = "OWNER"  // creator, can delete the channel
	ParticipantRoleAdmin  ParticipantRole = "ADMIN"  // manage members/settings
	ParticipantRoleMember ParticipantRole = "MEMBER" // normal
	ParticipantRoleGuest  ParticipantRole = "GUEST"  // read/post only
)

// Per-conversation notification preference.
type NotificationLevel string

const (
	NotificationLevelAll      NotificationLevel = "ALL"      // every message
	NotificationLevelMentions NotificationLevel = "MENTIONS" // only @me / @here / @everyone
	NotificationLevelNone     NotificationLevel = "NONE"     // muted
)

// Live presence of a member. Driven by websocket connections, not by CRUD.
type PresenceStatus string

const (
	PresenceStatusOnline  PresenceStatus = "ONLINE"
	PresenceStatusAway    PresenceStatus = "AWAY"
	PresenceStatusDND     PresenceStatus = "DND"
	PresenceStatusOffline PresenceStatus = "OFFLINE"
)

/* ---------------------------------- <***> --------------------------------- */
// Messages

type MessageType string

const (
	MessageTypeText   MessageType = "TEXT"   // normal user message (may still carry attachments)
	MessageTypeFile   MessageType = "FILE"   // attachment-only post
	MessageTypeSystem MessageType = "SYSTEM" // "X joined the channel", "topic changed", ...
	MessageTypeCall   MessageType = "CALL"   // call started/ended card
)

// Only meaningful when MessageType == SYSTEM. Tells the client what to render.
type MessageSubType string

const (
	MessageSubTypeMemberJoined    MessageSubType = "MEMBER_JOINED"
	MessageSubTypeMemberLeft      MessageSubType = "MEMBER_LEFT"
	MessageSubTypeMemberAdded     MessageSubType = "MEMBER_ADDED"
	MessageSubTypeMemberRemoved   MessageSubType = "MEMBER_REMOVED"
	MessageSubTypeChannelRenamed  MessageSubType = "CHANNEL_RENAMED"
	MessageSubTypeTopicChanged    MessageSubType = "TOPIC_CHANGED"
	MessageSubTypeChannelArchived MessageSubType = "CHANNEL_ARCHIVED"
	MessageSubTypeMessagePinned   MessageSubType = "MESSAGE_PINNED"

	// call cards rendered inline in the message list
	MessageSubTypeCallStarted    MessageSubType = "CALL_STARTED"
	MessageSubTypeCallEnded      MessageSubType = "CALL_ENDED"
	MessageSubTypeCallMissed     MessageSubType = "CALL_MISSED"
	MessageSubTypeRecordingReady MessageSubType = "RECORDING_READY"
	MessageSubTypeGuestJoined    MessageSubType = "GUEST_JOINED"
)

type MentionType string

const (
	MentionTypeMember   MentionType = "MEMBER"   // @someone
	MentionTypeEveryone MentionType = "EVERYONE" // @everyone in the conversation
	MentionTypeHere     MentionType = "HERE"     // @here, only currently online
	MentionTypeChannel  MentionType = "CHANNEL"  // #channel link
)

/* ---------------------------------- <***> --------------------------------- */
// Calls & conferencing

type CallType string

const (
	CallTypeAudio CallType = "AUDIO"
	CallTypeVideo CallType = "VIDEO"
)

type CallStatus string

const (
	CallStatusRinging CallStatus = "RINGING" // DM call, waiting for pickup
	CallStatusOngoing CallStatus = "ONGOING"
	CallStatusEnded   CallStatus = "ENDED"
	CallStatusMissed  CallStatus = "MISSED"
)

// Where the call came from. Changes who is allowed in and how it's announced.
// CHANNEL -> someone hit "start call" in a voice/video channel; the channel's
//
//	members can walk in and out freely, discord style
//
// DIRECT  -> rang a DM or group DM; it rings the other participants
// MEETING -> a named meeting room (scheduled or instant), the only kind that
//
//	can hand out public join links
type CallOrigin string

const (
	CallOriginChannel CallOrigin = "CHANNEL"
	CallOriginDirect  CallOrigin = "DIRECT"
	CallOriginMeeting CallOrigin = "MEETING"
)

// Role inside a single call. Separate from ParticipantRole — being a channel
// admin does not make you the host of a call happening in it.
type CallParticipantRole string

const (
	CallParticipantRoleHost        CallParticipantRole = "HOST"        // started it; can end for everyone
	CallParticipantRoleCoHost      CallParticipantRole = "COHOST"      // moderation, admit from lobby
	CallParticipantRoleParticipant CallParticipantRole = "PARTICIPANT" // normal, can speak/share
	CallParticipantRoleViewer      CallParticipantRole = "VIEWER"      // watch only, webinar style
)

// Lifecycle of one person in one call. LOBBY is the waiting room — the guest
// has a valid link but a host has not let them in yet.
type CallParticipantStatus string

const (
	CallParticipantStatusLobby    CallParticipantStatus = "LOBBY"
	CallParticipantStatusJoined   CallParticipantStatus = "JOINED"
	CallParticipantStatusLeft     CallParticipantStatus = "LEFT"
	CallParticipantStatusRemoved  CallParticipantStatus = "REMOVED"  // kicked by host
	CallParticipantStatusDenied   CallParticipantStatus = "DENIED"   // refused at the lobby
	CallParticipantStatusRejected CallParticipantStatus = "REJECTED" // declined a ringing DM call
)

/* ---------------------------------- <***> --------------------------------- */
// Meetings & public join links

type MeetingStatus string

const (
	MeetingStatusScheduled MeetingStatus = "SCHEDULED"
	MeetingStatusLive      MeetingStatus = "LIVE"
	MeetingStatusEnded     MeetingStatus = "ENDED"
	MeetingStatusCancelled MeetingStatus = "CANCELLED"
)

// The security question for every join attempt: who is this link good for?
//
// ORG_ONLY          -> must be a signed-in member of this org's cables account
// INVITED_ONLY      -> must be on the invitee list (or a member with a link)
// ANYONE_WITH_LINK  -> fully public, no sharpnr account needed
// ANYONE_WITH_CODE  -> public link but the passcode must also be entered
type MeetingJoinPolicy string

const (
	MeetingJoinPolicyOrgOnly        MeetingJoinPolicy = "ORG_ONLY"
	MeetingJoinPolicyInvitedOnly    MeetingJoinPolicy = "INVITED_ONLY"
	MeetingJoinPolicyAnyoneWithLink MeetingJoinPolicy = "ANYONE_WITH_LINK"
	MeetingJoinPolicyAnyoneWithCode MeetingJoinPolicy = "ANYONE_WITH_CODE"
)

// RSVP on a scheduled meeting.
type MeetingInviteeStatus string

const (
	MeetingInviteeStatusPending   MeetingInviteeStatus = "PENDING"
	MeetingInviteeStatusAccepted  MeetingInviteeStatus = "ACCEPTED"
	MeetingInviteeStatusDeclined  MeetingInviteeStatus = "DECLINED"
	MeetingInviteeStatusTentative MeetingInviteeStatus = "TENTATIVE"
)

type RecordingStatus string

const (
	RecordingStatusRecording  RecordingStatus = "RECORDING"
	RecordingStatusProcessing RecordingStatus = "PROCESSING" // provider is muxing/uploading
	RecordingStatusReady      RecordingStatus = "READY"
	RecordingStatusFailed     RecordingStatus = "FAILED"
	RecordingStatusDeleted    RecordingStatus = "DELETED"
)

// What the host chose to capture.
type RecordingKind string

const (
	RecordingKindAudioVideo RecordingKind = "AUDIO_VIDEO"
	RecordingKindAudioOnly  RecordingKind = "AUDIO_ONLY"
	RecordingKindTranscript RecordingKind = "TRANSCRIPT"
)

// Moderation/audit trail inside a call. Kept as an enum because the UI renders
// some of these live ("Ali raised their hand") and compliance needs the rest.
type CallEventType string

const (
	CallEventTypeJoined             CallEventType = "JOINED"
	CallEventTypeLeft               CallEventType = "LEFT"
	CallEventTypeKnocked            CallEventType = "KNOCKED" // guest hit the lobby
	CallEventTypeAdmitted           CallEventType = "ADMITTED"
	CallEventTypeDenied             CallEventType = "DENIED"
	CallEventTypeRemoved            CallEventType = "REMOVED"
	CallEventTypeMutedByHost        CallEventType = "MUTED_BY_HOST"
	CallEventTypeHandRaised         CallEventType = "HAND_RAISED"
	CallEventTypeHandLowered        CallEventType = "HAND_LOWERED"
	CallEventTypeScreenShareStarted CallEventType = "SCREEN_SHARE_STARTED"
	CallEventTypeScreenShareStopped CallEventType = "SCREEN_SHARE_STOPPED"
	CallEventTypeRecordingStarted   CallEventType = "RECORDING_STARTED"
	CallEventTypeRecordingStopped   CallEventType = "RECORDING_STOPPED"
	CallEventTypeRoleChanged        CallEventType = "ROLE_CHANGED"
)

/* ---------------------------------- <***> --------------------------------- */
