// sere/handlers.go

package sere

// Handler reacts to an event
// Returning error allows retry/DLQ in future versions
type Handler func(event Event) error
