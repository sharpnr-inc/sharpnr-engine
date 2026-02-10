package sero

import (
	"context"
	dbtx "sharpnr-engine/core/dbTx"
	"sync"
)

type Handler func(ctx context.Context, payload interface{}) error

type SeroEngineEventType string

type SEROEngine struct {
	txManager     dbtx.TxManager
	syncHandlers  map[SeroEngineEventType][]Handler
	asyncHandlers map[SeroEngineEventType][]Handler
	mu            sync.RWMutex
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
