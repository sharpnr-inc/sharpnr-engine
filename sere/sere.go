// sere/sere.go

package sere

import (
	"context"
	"sync"
)

type TxExecutor func(ctx context.Context, fn func(ctx context.Context) error) error

type registeredHandler struct {
	mode    HandlerMode
	handler Handler
}
type SERE struct {
	events              chan Event
	subscribers         map[string][]registeredHandler
	mu                  sync.RWMutex
	workflowSubscribers map[string][]WorkflowHandler
	txExecutor          TxExecutor
	started             chan struct{} // 👈 readiness signal
}

func NewSERE() *SERE {
	return &SERE{
		events:              make(chan Event, 50),
		subscribers:         make(map[string][]registeredHandler),
		workflowSubscribers: make(map[string][]WorkflowHandler),
		started:             make(chan struct{}),
	}
}
