package sere

import "fmt"

// Start boots the event dispatcher
func (s *SERE) Start() {
	fmt.Println("🚀 SERE starting...")

	go func() {
		for event := range s.events {

			s.mu.RLock()
			handlers := s.subscribers[event.Name]
			s.mu.RUnlock()

			if len(handlers) == 0 {
				fmt.Println("⚠️ No handlers for event:", event.Name)
				continue
			}

			for _, handler := range handlers {
				go func(h Handler) {
					if err := h(event); err != nil {
						// v1: log only
						fmt.Println("❌ Event handler error:", err)
					}
				}(handler)
			}
		}
	}()
}
