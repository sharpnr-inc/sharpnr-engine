package kitchenservices

import (
	"context"
	"log/slog"
	"sync"

	kitchen_contracts "sharpnr-engine/business_ops/hospitality/kitchen/contracts"

	"github.com/google/uuid"
)

// broadcastKots fetches the kitchen's current KOT list and publishes it to
// SSE subscribers. Callers run it via `go s.broadcastKots(...)` to keep API
// responses fast.
//
// Fetch+publish is serialized per kitchen: without the lock, two concurrent
// updates could publish in reverse order (older fetch published last) and
// leave every subscriber on a stale snapshot.
func (s *kitchenService) broadcastKots(kitchenID uuid.UUID) {
	lock, _ := s.kotBroadcastLocks.LoadOrStore(kitchenID, &sync.Mutex{})
	mu := lock.(*sync.Mutex)
	mu.Lock()
	defer mu.Unlock()

	// Fresh context: the originating request may already be cancelled.
	kots, err := s.GetKots(context.Background(), kitchenID)
	if err != nil {
		slog.Error("sse: KOT broadcast fetch failed", "kitchen_id", kitchenID, "error", err)
		return
	}

	s.sseEngine.Publish(kitchen_contracts.KotsTopic(kitchenID), kots)
}
