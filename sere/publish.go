package sere

import (
	"context"
	"time"
)

// Publish announces a business fact
// It NEVER blocks the caller
func (s *SERE) Publish(
	ctx context.Context,
	name EventName,
	payload any,
) {
	s.events <- Event{
		Ctx:       ctx,
		Name:      name,
		Payload:   payload,
		CreatedAt: time.Now(),
	}
}
