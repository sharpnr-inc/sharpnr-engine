package sero

import (
	"context"
	"encoding/json"
	"time"
)

func (e *SEROEngine) StartRelay(ctx context.Context) {
	ticker := time.NewTicker(2 * time.Second)
	defer ticker.Stop()

	for {
		select {
		case <-ctx.Done():
			return
		case <-ticker.C:
			e.processBatch(ctx)
		}
	}
}

func (e *SEROEngine) processBatch(ctx context.Context) {
	events, err := e.sero_repo.FetchPending(ctx, 10)
	if err != nil || len(events) == 0 {
		return
	}

	for _, evt := range events {
		if err := e.executeAsync(evt); err != nil {
			e.sero_repo.UpdateStatus(ctx, evt.ID, "FAILED", err.Error())
		} else {
			e.sero_repo.UpdateStatus(ctx, evt.ID, "PROCESSED", "")
		}
	}
}

func (e *SEROEngine) executeAsync(evt SeroOutbox) error {
	e.mu.RLock()
	handlers := e.asyncHandlers[evt.EventName]
	e.mu.RUnlock()

	var payload interface{}
	json.Unmarshal(evt.Payload, &payload)

	for _, h := range handlers {
		// Async handlers get a fresh background context
		if err := h(context.Background(), payload); err != nil {
			return err
		}
	}
	return nil
}
