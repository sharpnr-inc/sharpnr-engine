package sere

// Subscribe registers a handler for an event name
// func (s *SERE) Subscribe(eventName string, handler Handler) {
// 	s.mu.Lock()
// 	defer s.mu.Unlock()

// 	s.subscribers[eventName] = append(
// 		s.subscribers[eventName],
// 		handler,
// 	)
// }

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
