package sere

import (
	"context"
	"log"
	"time"
)

// Start boots the event dispatcher
// Start boots the event dispatcher
func (s *SERE) Start() {
	log.Println("🚀 SERE starting...")

	go func() {

		// 🔥 Protect dispatcher from crashing
		defer func() {
			if r := recover(); r != nil {
				log.Println("🔥 SERE dispatcher panic recovered:", r)
			}
		}()

		// ✅ Signal readiness
		close(s.started)
		log.Println("✅ SERE started")

		for event := range s.events {

			s.mu.RLock()
			handlers := s.subscribers[event.Name]
			s.mu.RUnlock()

			if len(handlers) == 0 {
				log.Println("⚠️ No handlers for event:", event.Name)
				continue
			}

			for _, rh := range handlers {

				// 🔒 SYNC HANDLER
				if rh.mode == Sync {
					if err := rh.handler(event); err != nil {
						log.Println(
							"❌ Sync handler failed:",
							"event =", event.Name,
							"err =", err,
						)
						continue // or break, policy choice
					}
					continue
				}

				// 🚀 ASYNC HANDLER
				go func(h Handler, evt Event) {
					defer func() {
						if r := recover(); r != nil {
							log.Println("🔥 Async handler panic recovered:", r)
						}
					}()

					ctx, cancel := context.WithTimeout(evt.Ctx, 10*time.Second)
					defer cancel()

					evt.Ctx = ctx

					if err := h(evt); err != nil {
						log.Println(
							"❌ Async handler error:",
							"event =", evt.Name,
							"err =", err,
						)
					}
				}(rh.handler, event)
			}

		}

		log.Println("🛑 SERE stopped")
	}()
}

func (s *SERE) WaitUntilStarted() {
	<-s.started
}
