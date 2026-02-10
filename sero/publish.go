package sero

import (
	"context"
	"fmt"
	"log"
)

func (e *SEROEngine) Publish(ctx context.Context, eventName SeroEngineEventType, payload interface{}) error {
	e.mu.RLock()
	syncs := e.syncHandlers[eventName]
	asyncs := e.asyncHandlers[eventName]
	e.mu.RUnlock()

	for _, handler := range syncs {
		if err := handler(ctx, payload); err != nil {
			return fmt.Errorf("sync handler failed for the %s: %w", eventName, err)
		}
	}

	for _, handler := range asyncs {
		go func(h Handler) {
			// We use a background context for async to avoid cancellation
			// if the main request context finishes.
			// Async tasks usually shouldn't share the same DB transaction.
			err := h(context.Background(), payload)
			if err != nil {
				log.Printf("SERO Async Error [%s]: %v", eventName, err)
			}

		}(handler)
	}

	return nil
}
