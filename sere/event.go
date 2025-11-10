package sere

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/google/uuid"
)

// Event is the unit dispatched by SERE from publishers to subscribers.
// Keep Payload JSON-serializable and reasonably small.
type Event struct {
	ID        uuid.UUID `json:"id"`
	Type      EventType `json:"type"`
	OrgID     uuid.UUID `json:"org_id"`
	Module    string    `json:"module"` // e.g., "restaurant", "kitchen"
	Topic     Topic     `json:"topic"`
	Payload   any       `json:"payload"`
	Version   int       `json:"version"` // schema version for payload shape
	EmittedAt time.Time `json:"emitted_at"`
	Meta      Metadata  `json:"meta"`
}

// NewEvent constructs a validated Event with defaults filled.
func NewEvent(t EventType, orgID uuid.UUID, topic Topic, module string, payload any, opts ...func(*Event)) (Event, error) {
	ev := Event{
		ID:        uuid.New(),
		Type:      t,
		OrgID:     orgID,
		Module:    module,
		Topic:     topic,
		Payload:   payload,
		Version:   1,
		EmittedAt: time.Now().UTC(),
	}
	for _, o := range opts {
		o(&ev)
	}
	if err := ev.Validate(); err != nil {
		return Event{}, err
	}
	return ev, nil
}

// Validate basic required fields.
func (e Event) Validate() error {
	if e.Type == "" || e.OrgID == uuid.Nil || e.Topic == "" {
		return ErrInvalidEvent
	}
	return nil
}

// JSON renders the entire event to bytes.
func (e Event) JSON() []byte {
	b, _ := json.Marshal(e)
	return b
}

// WithMeta allows setting Metadata in NewEvent.
func WithMeta(md Metadata) func(*Event) { return func(e *Event) { e.Meta = md } }

// WithVersion sets a custom schema version.
func WithVersion(v int) func(*Event) { return func(e *Event) { e.Version = v } }

// String is a concise debug representation.
func (e Event) String() string {
	return fmt.Sprintf("%s %s %s %s@%s", e.ID.String(), e.Type, e.OrgID, e.Topic, e.EmittedAt.Format(time.RFC3339))
}
