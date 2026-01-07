// sere/events.go

package sere

import (
	"context"
	"time"
)

type EventName string

// Event represents a business fact that already happened
type Event struct {
	Ctx       context.Context
	Name      EventName
	Payload   any       // domain-specific data
	CreatedAt time.Time // when event was created
}
