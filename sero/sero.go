package sero

import (
	"context"
	"encoding/json"
	"fmt"
	dbtx "sharpnr-engine/core/dbTx"
	"sync"
)

type Handler func(ctx context.Context, payload interface{}) error

type SeroEngineEventType string

type SEROEngine struct {
	sero_repo     OutboxRepository
	txManager     dbtx.TxManager
	syncHandlers  map[SeroEngineEventType][]Handler
	asyncHandlers map[SeroEngineEventType][]Handler
	mu            sync.RWMutex
}

func NewSEROEngine(repo OutboxRepository) *SEROEngine {
	return &SEROEngine{
		sero_repo:     repo,
		syncHandlers:  make(map[SeroEngineEventType][]Handler),
		asyncHandlers: make(map[SeroEngineEventType][]Handler),
	}
}

func (e *SEROEngine) Subscribe(eventName SeroEngineEventType, handler Handler, isAsync bool) {
	e.mu.Lock()
	defer e.mu.Unlock()

	if isAsync {
		e.asyncHandlers[eventName] = append(e.asyncHandlers[eventName], handler)
	} else {
		e.syncHandlers[eventName] = append(e.syncHandlers[eventName], handler)
	}
}

func (e *SEROEngine) Publish(ctx context.Context, evt SeroEngineEventType, payload interface{}) error {
	// 1. Get the current transaction
	tx, ok := dbtx.ExtractTx(ctx)
	if !ok {
		return fmt.Errorf("sero: no active transaction found in context")
	}

	// 2. Persist to Outbox (Atomic with business logic)
	jsonPayload, _ := json.Marshal(payload)
	outboxEntry := &SeroOutbox{
		EventName: evt,
		Payload:   jsonPayload,
		Status:    "PENDING",
	}
	if err := e.sero_repo.Save(ctx, tx, outboxEntry); err != nil {
		return err
	}

	// 3. Execute Sync Handlers immediately (Blocking)
	e.mu.RLock()
	handlers := e.syncHandlers[evt]
	e.mu.RUnlock()

	for _, h := range handlers {
		if err := h(ctx, payload); err != nil {
			return err // Fail the whole transaction if sync handler fails
		}
	}
	return nil
}
