package auditEngineDtos

import (
	"encoding/json"
	"time"

	"github.com/google/uuid"
)

// The read side of the audit log: what a stored event looks like once it is on
// its way to a screen. They live here, beside RecordEventInput and
// ListLogsQuery, so the whole audit vocabulary — write input, read query, read
// output — is one package.

// AuditLogActor is who did it. Name is filled in by the caller that has an
// employee directory to resolve it against; it stays empty for system events
// and for an actor who is no longer on the roster.
type AuditLogActor struct {
	ID        *uuid.UUID `json:"id,omitempty"`
	Name      string     `json:"name,omitempty"`
	Type      string     `json:"type"`
	IPAddress string     `json:"ipAddress,omitempty"`
}

// AuditLogEntry is one event in the organization's history.
//
// Metadata is passed through as raw JSON: it is the emitting module's own
// payload — changed fields, old/new values — and re-shaping it would mean this
// package had to know every module's event schema.
type AuditLogEntry struct {
	ID    uuid.UUID     `json:"id"`
	Actor AuditLogActor `json:"actor"`

	Module string `json:"module"`
	Action string `json:"action"`

	ResourceType string     `json:"resourceType,omitempty"`
	ResourceID   *uuid.UUID `json:"resourceId,omitempty"`

	Description string          `json:"description,omitempty"`
	Metadata    json.RawMessage `json:"metadata,omitempty"`

	CreatedAt time.Time `json:"createdAt"`
}

// PaginationMeta describes the page that was served. HasMore saves the frontend
// from doing the Offset+Limit vs Total arithmetic itself on every render.
type PaginationMeta struct {
	Total   int64 `json:"total"`
	Limit   int   `json:"limit"`
	Offset  int   `json:"offset"`
	HasMore bool  `json:"hasMore"`
}

// AuditLogPageResponseDto is one page of the organization's history.
type AuditLogPageResponseDto struct {
	Logs       []AuditLogEntry `json:"logs"`
	Pagination PaginationMeta  `json:"pagination"`
}
