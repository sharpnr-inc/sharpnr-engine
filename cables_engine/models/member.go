package cables_engine_models

import (
	cablesEngineConstants "sharpnr-engine/cables_engine/constants"
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

// CablesMember is an identity inside cables — the thing that authors messages,
// joins conversations and appears in a call roster.
//
// It is NOT the employee row, and that separation is the point. A bot posting
// build results and a guest who joined a meeting off a public link both need to
// appear in a member list and be referenced by a message, and neither has an
// employee id. Making cables_messages.sender_member_id point here instead means
// no nullable-employee branch anywhere downstream.
//
// It is also where identity facts that belong to cables rather than to HR live:
// the display name people see in chat, and presence, which is driven by
// websocket connections rather than by any CRUD.
//
// Rows for employees are created by syncing the org roster (see
// SyncMembersFromEmployees) rather than by the employee module writing here —
// cables owns its own table, and the sync is the one cross-module read.
//
// indexes:
//
//	index  (org_cables_account_id)
//	unique (org_cables_account_id, employee_id) where employee_id is not null
type CablesMember struct {
	ID                 uuid.UUID
	OrganizationID     uuid.UUID
	OrgCablesAccountID uuid.UUID

	// Null for BOT and GUEST members. Set, and unique per account, for the
	// EMPLOYEE members that make up the org roster.
	EmployeeID *uuid.UUID

	Type cablesEngineConstants.MemberType // EMPLOYEE | BOT | GUEST

	// DisplayName is denormalized from the employee at sync time so rendering a
	// page of messages never joins to the employee module. A rename in HR
	// reaches cables on the next sync, which is the right latency for a
	// display name.
	DisplayName string
	AvatarURL   *string

	// Presence is written by the websocket layer, not by a request handler.
	Presence     cablesEngineConstants.PresenceStatus
	LastSeenAt   *time.Time
	CustomStatus *string

	// IsEphemeral marks a guest member minted for one meeting; the retention
	// sweeper is allowed to delete these.
	IsEphemeral bool

	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt
}

func (CablesMember) TableName() string {
	return "cables_engine.cables_members"
}

/* ---------------------------------- <***> --------------------------------- */
