package sere

import "sync"

type SERE struct {
	events      chan Event
	subscribers map[string][]Handler
	mu          sync.RWMutex

	started chan struct{} // 👈 readiness signal
}

func NewSERE() *SERE {
	return &SERE{
		events:      make(chan Event, 50),
		subscribers: make(map[string][]Handler),

		started: make(chan struct{}),
	}
}
