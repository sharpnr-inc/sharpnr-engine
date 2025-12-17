package sere

import "time"

// Event represents a business fact that already happened
type Event struct {
	Name      string    // e.g. "OrganizationSignupCompleted"
	Payload   any       // domain-specific data
	CreatedAt time.Time // when event was created
}
