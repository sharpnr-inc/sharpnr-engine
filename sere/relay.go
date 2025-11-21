package sere

import "sync"

type SSEClient chan string

type SSEHub struct {
	Clients map[SSEClient]bool
	Lock    sync.Mutex
}

func NewSSEHub() *SSEHub {
	return &SSEHub{
		Clients: make(map[SSEClient]bool),
	}
}

func (h *SSEHub) Add(c SSEClient) {
	h.Lock.Lock()
	h.Clients[c] = true
	h.Lock.Unlock()
}

func (h *SSEHub) Remove(c SSEClient) {
	h.Lock.Lock()
	delete(h.Clients, c)
	close(c)
	h.Lock.Unlock()
}

func (h *SSEHub) Broadcast(msg string) {
	h.Lock.Lock()
	defer h.Lock.Unlock()

	for client := range h.Clients {
		select {
		case client <- msg:
		default:
			close(client)
			delete(h.Clients, client)
		}
	}
}
