// Package sse_engine is an in-process pub/sub hub backing the app's
// Server-Sent Event streams.
//
// Delivery semantics: each subscriber holds only the LATEST event for its
// topic. If a client consumes slower than events are published, older
// undelivered events are replaced by newer ones. Payloads must therefore be
// full snapshots of state (e.g. the complete KOT list), never deltas — a
// replaced event must be safe to skip.
package sse_engine

import (
	"sync"
)

type Event struct {
	Topic   string // e.g. "kitchen_<uuid>"
	Payload any    // full-state snapshot for the topic
}

type StreamManager struct {
	mu          sync.RWMutex
	subscribers map[string]map[chan Event]struct{}
	events      chan Event
	done        chan struct{}
	closeOnce   sync.Once
}

func NewStreamManager() *StreamManager {
	sm := &StreamManager{
		subscribers: make(map[string]map[chan Event]struct{}),
		events:      make(chan Event, 1000),
		done:        make(chan struct{}),
	}
	go sm.listen()
	return sm
}

// Publish enqueues an event for fan-out to the topic's subscribers.
// After Close it is a no-op.
func (sm *StreamManager) Publish(topic string, payload any) {
	select {
	case <-sm.done:
	case sm.events <- Event{Topic: topic, Payload: payload}:
	}
}

// Subscription is one client's registration on a topic. Read events from C;
// Close is idempotent and safe to defer.
type Subscription struct {
	Topic string
	ch    chan Event
	sm    *StreamManager
}

// C is closed when either the subscription or the whole manager is closed,
// so range/comma-ok reads terminate on their own.
func (s *Subscription) C() <-chan Event { return s.ch }

func (s *Subscription) Close() {
	s.sm.unsubscribe(s.Topic, s.ch)
}

func (sm *StreamManager) Subscribe(topic string) *Subscription {
	sm.mu.Lock()
	defer sm.mu.Unlock()

	// Capacity 1: the channel is a latest-value slot, not a queue.
	ch := make(chan Event, 1)
	sub := &Subscription{Topic: topic, ch: ch, sm: sm}

	select {
	case <-sm.done:
		// Manager already closed: return an already-closed channel so the
		// caller's read loop exits immediately.
		close(ch)
		return sub
	default:
	}

	if sm.subscribers[topic] == nil {
		sm.subscribers[topic] = make(map[chan Event]struct{})
	}
	sm.subscribers[topic][ch] = struct{}{}
	return sub
}

func (sm *StreamManager) unsubscribe(topic string, ch chan Event) {
	sm.mu.Lock()
	defer sm.mu.Unlock()

	clients, ok := sm.subscribers[topic]
	if !ok {
		return
	}
	// Only close a channel still present in the map — a second Close on the
	// same subscription must not close the channel twice.
	if _, exists := clients[ch]; !exists {
		return
	}
	delete(clients, ch)
	close(ch)
	if len(clients) == 0 {
		delete(sm.subscribers, topic)
	}
}

// Close stops fan-out and closes every subscriber channel, disconnecting all
// clients. Publish becomes a no-op. Safe to call more than once.
func (sm *StreamManager) Close() {
	sm.closeOnce.Do(func() {
		close(sm.done)

		sm.mu.Lock()
		defer sm.mu.Unlock()
		for _, clients := range sm.subscribers {
			for ch := range clients {
				close(ch)
			}
		}
		sm.subscribers = make(map[string]map[chan Event]struct{})
	})
}

// listen runs in the background and fans out published events.
func (sm *StreamManager) listen() {
	for {
		select {
		case <-sm.done:
			return
		case ev := <-sm.events:
			sm.fanOut(ev)
		}
	}
}

func (sm *StreamManager) fanOut(ev Event) {
	// Holding the read lock during sends is what makes the non-blocking
	// sends safe: unsubscribe/Close need the write lock, so a channel can
	// never be closed mid-send.
	sm.mu.RLock()
	defer sm.mu.RUnlock()

	for ch := range sm.subscribers[ev.Topic] {
		select {
		case ch <- ev:
		default:
			// Subscriber hasn't consumed the previous event yet: drop the
			// stale one and install this newer snapshot in its place.
			select {
			case <-ch:
			default:
			}
			select {
			case ch <- ev:
			default:
			}
		}
	}
}
