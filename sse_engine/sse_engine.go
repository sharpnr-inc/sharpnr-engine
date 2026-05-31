package sse_engine

import (
	"sync"
)

type Event struct {
	Topic   string      // e.g., "kitchen_123" or "bar_456"
	Payload interface{} // The actual KOT data
}

type StreamManager struct {
	// Map of Topic -> Map of Client Channels
	// We use a map of maps to quickly find who is listening to "kitchen_1"
	Subscribers map[string]map[chan interface{}]bool
	Mu          sync.RWMutex
	MessageChan chan Event
}

func NewStreamManager() *StreamManager {
	sm := &StreamManager{
		Subscribers: make(map[string]map[chan interface{}]bool),
		MessageChan: make(chan Event, 1000), // Buffer for incoming events
	}
	go sm.listen()
	return sm
}

// listen runs in the background and fans out messages to subscribers
func (sm *StreamManager) listen() {
	for event := range sm.MessageChan {
		sm.Mu.RLock()
		if clients, ok := sm.Subscribers[event.Topic]; ok {
			for clientChan := range clients {
				// Non-blocking send to avoid one slow client hanging the engine
				select {
				case clientChan <- event.Payload:
				default:
				}
			}
		}
		sm.Mu.RUnlock()
	}
}

func (sm *StreamManager) Subscribe(topic string) chan interface{} {
	sm.Mu.Lock()
	defer sm.Mu.Unlock()

	clientChan := make(chan interface{}, 10)
	if _, ok := sm.Subscribers[topic]; !ok {
		sm.Subscribers[topic] = make(map[chan interface{}]bool)
	}
	sm.Subscribers[topic][clientChan] = true
	return clientChan
}

func (sm *StreamManager) Unsubscribe(topic string, clientChan chan interface{}) {
	sm.Mu.Lock()
	defer sm.Mu.Unlock()
	if clients, ok := sm.Subscribers[topic]; ok {
		delete(clients, clientChan)
		close(clientChan)
		if len(clients) == 0 {
			delete(sm.Subscribers, topic)
		}
	}
}
