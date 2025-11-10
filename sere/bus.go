package sere

import (
	"sync"
	"time"

	"github.com/google/uuid"
)

// Engine is the core in-memory bus. It is concurrency-safe.
type Engine struct {
	mu          sync.RWMutex
	closed      bool
	subscribers map[uuid.UUID]*Subscriber
}

// NewEngine creates a fresh SERE bus.
func NewEngine() *Engine {
	return &Engine{subscribers: make(map[uuid.UUID]*Subscriber)}
}

// Subscribe registers a subscriber and starts its worker.
func (e *Engine) Subscribe(name string, match MatchFn, handle HandleFn, buffer int, timeout time.Duration) (uuid.UUID, error) {
	e.mu.Lock()
	defer e.mu.Unlock()
	if e.closed {
		return uuid.Nil, ErrClosed
	}
	sub := newSubscriber(name, match, handle, buffer, timeout)
	e.subscribers[sub.ID] = sub
	go sub.run()
	return sub.ID, nil
}

// Unsubscribe removes a subscriber and closes its channel.
func (e *Engine) Unsubscribe(id uuid.UUID) {
	e.mu.Lock()
	defer e.mu.Unlock()
	if s, ok := e.subscribers[id]; ok {
		delete(e.subscribers, id)
		close(s.ch)
		s.stop()
	}
}

// Publish fan-outs the event to all matching subscribers in a non-blocking way.
func (e *Engine) Publish(ev Event) error {
	if err := ev.Validate(); err != nil {
		return err
	}
	e.mu.RLock()
	defer e.mu.RUnlock()
	if e.closed {
		return ErrClosed
	}
	for _, s := range e.subscribers {
		if s.Match != nil && !s.Match(ev) {
			continue
		}
		select {
		case s.ch <- ev:
		default:
			// channel full → drop event to protect bus
		}
	}
	return nil
}

// Close stops all subscribers and prevents further publishes/subscribes.
func (e *Engine) Close() {
	e.mu.Lock()
	defer e.mu.Unlock()
	if e.closed {
		return
	}
	e.closed = true
	for id, s := range e.subscribers {
		delete(e.subscribers, id)
		close(s.ch)
		s.stop()
	}
}
