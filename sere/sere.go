package sere

import "sync"

type SERE struct {
	events      chan Event
	subscribers map[string][]Handler
	mu          sync.RWMutex
}

func NewSERE() *SERE {
	return &SERE{
		events:      make(chan Event, 50),
		subscribers: make(map[string][]Handler),
	}
}
