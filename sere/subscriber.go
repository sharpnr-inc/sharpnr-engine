package sere

import (
	"context"
	"strings"
	"time"

	"github.com/google/uuid"
)

// MatchFn determines if a subscriber cares about an event.
type MatchFn func(Event) bool

// HandleFn processes an event. Keep it fast; offload heavy work yourself.
type HandleFn func(context.Context, Event)

// Subscriber represents a registered consumer.
type Subscriber struct {
	ID         uuid.UUID
	Name       string
	Match      MatchFn
	Handle     HandleFn
	BufferSize int           // channel buffer (default 256)
	Timeout    time.Duration // per-event handling deadline (0 = no deadline)

	ch   chan Event
	quit chan struct{}
}

// Helper: match by exact type.
func MatchType(t EventType) MatchFn { return func(e Event) bool { return e.Type == t } }

// Helper: match by org ID.
func MatchOrg(org uuid.UUID) MatchFn { return func(e Event) bool { return e.OrgID == org } }

// Helper: match by topic prefix (e.g., "org:<uuid>:kitchen").
func MatchTopicPrefix(prefix string) MatchFn {
	return func(e Event) bool { return strings.HasPrefix(string(e.Topic), prefix) }
}

// And composes multiple MatchFns (all must be true).
func And(fns ...MatchFn) MatchFn {
	return func(e Event) bool {
		for _, fn := range fns {
			if fn == nil || !fn(e) {
				return false
			}
		}
		return true
	}
}

// Or composes multiple MatchFns (any true).
func Or(fns ...MatchFn) MatchFn {
	return func(e Event) bool {
		for _, fn := range fns {
			if fn != nil && fn(e) {
				return true
			}
		}
		return false
	}
}

// newSubscriber wires internal channels but does not start the worker.
func newSubscriber(name string, match MatchFn, handle HandleFn, buf int, timeout time.Duration) *Subscriber {
	if buf <= 0 {
		buf = 256
	}
	return &Subscriber{
		ID:         uuid.New(),
		Name:       name,
		Match:      match,
		Handle:     handle,
		BufferSize: buf,
		Timeout:    timeout,
		ch:         make(chan Event, buf),
		quit:       make(chan struct{}),
	}
}

// run starts the delivery loop for this subscriber.
func (s *Subscriber) run() {
	for {
		select {
		case ev, ok := <-s.ch:
			if !ok {
				return
			}
			if s.Timeout <= 0 {
				s.Handle(context.Background(), ev)
				continue
			}
			ctx, cancel := context.WithTimeout(context.Background(), s.Timeout)
			s.Handle(ctx, ev)
			cancel()
		case <-s.quit:
			return
		}
	}
}

// stop signals the worker to finish.
func (s *Subscriber) stop() { close(s.quit) }
