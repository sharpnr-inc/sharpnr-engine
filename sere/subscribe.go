package sere

// Subscribe registers a handler for an event name
func (s *SERE) Subscribe(eventName string, handler Handler) {
	s.mu.Lock()
	defer s.mu.Unlock()

	s.subscribers[eventName] = append(
		s.subscribers[eventName],
		handler,
	)
}
