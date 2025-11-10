// ============================================================================
// PATH: sere/types.go
// ============================================================================
package sere

import (
	"errors"

	"github.com/google/uuid"
)

// EventType is a semantic identifier for events (namespaced by domain).
// Example: "hospitality.restaurant.order_created".
type EventType string

// Topic is a routing key for subscribers.
// Convention: org:<OrgUUID>[:<module>[:<resource>[:<id>...]]
type Topic string

// Metadata carries optional tracing/diagnostic fields that ride with the Event.
type Metadata struct {
	TraceID   uuid.UUID         `json:"trace_id"`
	ActorID   uuid.UUID         `json:"actor_id"` // who caused the event (user/service)
	ActorName string            `json:"actor_name"`
	SourceIP  string            `json:"source_ip"`
	Tags      map[string]string `json:"tags,omitempty"`
}

// Common errors returned by the engine or validators.
var (
	ErrInvalidEvent   = errors.New("sere: invalid event (missing type, org, or topic)")
	ErrClosed         = errors.New("sere: engine is closed")
	ErrSubscriberGone = errors.New("sere: subscriber removed")
)
