package auditEngineDtos

import (
	"github.com/google/uuid"
)

// RecordEventInput is what every module passes to the audit log service
// when something happens inside a tenant.
type RecordEventInput struct {
	OrganizationID uuid.UUID

	ActorID        *uuid.UUID // employee who did it; nil for system-triggered events
	ActorType      string     // models.ActorTypeEmployee | models.ActorTypeSystem
	ActorIPAddress string     // client IP the request came from, when known

	Module string // emitting engine: "hrms", "customers", "payments", ...
	Action string // dot-notation event name: "customer.created"

	ResourceType string
	ResourceID   *uuid.UUID

	Description string
	Metadata    map[string]interface{} // optional extra payload, stored as JSONB
}

// ListLogsQuery is the pagination/filter input for reading an org's history.
type ListLogsQuery struct {
	Limit  int // defaults to 50, capped at 200
	Offset int

	// Optional filters — zero values mean "no filter".
	Module       string
	Action       string
	ResourceType string
	ResourceID   *uuid.UUID
	ActorID      *uuid.UUID
}
