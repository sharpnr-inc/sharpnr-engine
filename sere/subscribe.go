// sere/subscribe.go
package sere

func (s *SERE) SubscribeAsync(eventName string, handler Handler) {
	s.mu.Lock()
	defer s.mu.Unlock()

	s.subscribers[eventName] = append(
		s.subscribers[eventName],
		registeredHandler{
			mode:    Async,
			handler: handler,
		},
	)
}

func (s *SERE) SubscribeSync(eventName string, handler Handler) {
	s.mu.Lock()
	defer s.mu.Unlock()

	s.subscribers[eventName] = append(
		s.subscribers[eventName],
		registeredHandler{
			mode:    Sync,
			handler: handler,
		},
	)
}
