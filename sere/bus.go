package sharpnreventbus

import (
	"log"
)

type EventBus struct {
	ch chan SharpEvent
}

// 🔧 Create a new event bus with buffered channel
func NewEventBus(bufferSize int) *EventBus {
	return &EventBus{
		ch: make(chan SharpEvent, bufferSize),
	}
}

// 🚀 Publish an event to the bus
func (b *EventBus) Publish(event SharpEvent) {
	select {
	case b.ch <- event:
		log.Printf("[EventBus] Published event: %s", event.Name)
	default:
		log.Printf("[EventBus] WARNING: Event channel full, dropped: %s", event.Name)
	}
}

// 🎧 Subscribe returns a read-only channel for consumers
func (b *EventBus) Subscribe() <-chan SharpEvent {
	return b.ch
}
