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
//
// The tags are here because this is also the shape the history screen sends:
// the orchestrator binds the query string straight onto it rather than keeping
// a second, identical struct of its own.
type ListLogsQuery struct {
	Limit  int `json:"limit" form:"limit"` // defaults to 50, capped at 200
	Offset int `json:"offset" form:"offset"`

	// Optional filters — zero values mean "no filter".
	Module       string     `json:"module,omitempty" form:"module"`
	Action       string     `json:"action,omitempty" form:"action"`
	ResourceType string     `json:"resourceType,omitempty" form:"resourceType"`
	ResourceID   *uuid.UUID `json:"resourceId,omitempty" form:"resourceId"`
	ActorID      *uuid.UUID `json:"actorId,omitempty" form:"actorId"`
}
