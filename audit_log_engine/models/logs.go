// Audit log entries — one row per tenant (organization) event.
// Every engine injects the audit log service and records what happened,
// so the organization has a full history of who did what and when.
package auditEngineModels

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/datatypes"
)

// Who performed the action.
const (
	ActorTypeEmployee = "employee"
	ActorTypeSystem   = "system"
)

type AuditLog struct {
	ID             uuid.UUID
	OrganizationID uuid.UUID // tenant scope — every event belongs to an org

	ActorID        *uuid.UUID // employee who performed the action; nil for system events
	ActorType      string
	ActorIPAddress string

	Module string // engine/module that emitted the event: "hrms", "customers", "payments", ...
	Action string // dot-notation event name: "customer.created", "employee.updated", ...

	ResourceType string     // entity the event is about: "customer", "employee", ...
	ResourceID   *uuid.UUID // id of that entity, when applicable

	Description string         // human-readable summary shown in the org's history
	Metadata    datatypes.JSON // free-form event payload (changed fields, old/new values, ...)

	CreatedAt time.Time
}

func (AuditLog) TableName() string {
	return "audit_logs.logs"
}
