package sere

import (
	"context"
	"time"
)

// Event represents a business fact that already happened
type Event struct {
	Ctx       context.Context
	Name      string    // e.g. "OrganizationSignupCompleted"
	Payload   any       // domain-specific data
	CreatedAt time.Time // when event was created
}
