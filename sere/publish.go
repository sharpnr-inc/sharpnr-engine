package sere

import "time"

// Publish announces a business fact
// It NEVER blocks the caller
func (s *SERE) Publish(name string, payload any) {
	s.events <- Event{
		Name:      name,
		Payload:   payload,
		CreatedAt: time.Now(),
	}
}
