package sere

import "sync"

type registeredHandler struct {
	mode    HandlerMode
	handler Handler
}
type SERE struct {
	events      chan Event
	subscribers map[string][]registeredHandler
	mu          sync.RWMutex

	started chan struct{} // 👈 readiness signal
}

func NewSERE() *SERE {
	return &SERE{
		events:      make(chan Event, 50),
		subscribers: make(map[string][]registeredHandler),

		started: make(chan struct{}),
	}
}
