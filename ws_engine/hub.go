package ws_engine

import "sync"

// hub is the routing table: which live connections want which topics.
//
// Two maps rather than one because both directions are hot. Fan-out needs
// topic -> connections on every published event; a disconnect needs the
// connection's own topic set to unregister itself without scanning the whole
// table. The per-connection set lives on the connection (conn.topics) under the
// same lock, so the two never disagree.
type hub struct {
	mu     sync.RWMutex
	topics map[string]map[*connection]struct{}
	conns  map[*connection]struct{}
	closed bool
}

func newHub() *hub {
	return &hub{
		topics: make(map[string]map[*connection]struct{}),
		conns:  make(map[*connection]struct{}),
	}
}

// add registers a new connection. Returns false once the hub is closed, so a
// socket accepted during shutdown is rejected instead of being orphaned.
func (h *hub) add(c *connection) bool {
	h.mu.Lock()
	defer h.mu.Unlock()

	if h.closed {
		return false
	}
	h.conns[c] = struct{}{}
	return true
}

// remove drops the connection and every subscription it held. Idempotent: the
// read and write pumps both call it on the way out.
func (h *hub) remove(c *connection) {
	h.mu.Lock()
	defer h.mu.Unlock()

	delete(h.conns, c)

	for topic := range c.topics {
		h.detachLocked(topic, c)
	}
	c.topics = nil
}

func (h *hub) subscribe(c *connection, topic string) {
	h.mu.Lock()
	defer h.mu.Unlock()

	if h.closed {
		return
	}

	// A subscribe racing this connection's own teardown must not re-add it to
	// the topic map that remove has just cleaned, or the entry outlives the
	// socket and fan-out keeps finding a dead connection. shutdown closes done
	// before it calls remove, so testing it under this lock is decisive: either
	// we see it open and remove will run after us, or we see it closed and skip.
	select {
	case <-c.done:
		return
	default:
	}

	if h.topics[topic] == nil {
		h.topics[topic] = make(map[*connection]struct{})
	}
	h.topics[topic][c] = struct{}{}

	if c.topics == nil {
		c.topics = make(map[string]struct{})
	}
	c.topics[topic] = struct{}{}
}

func (h *hub) unsubscribe(c *connection, topic string) {
	h.mu.Lock()
	defer h.mu.Unlock()

	h.detachLocked(topic, c)
	delete(c.topics, topic)
}

// detachLocked removes one connection from one topic and drops the topic entry
// when it empties, so an org that churns through channels does not leave the
// map full of empty sets. Caller holds the write lock.
func (h *hub) detachLocked(topic string, c *connection) {
	subs, ok := h.topics[topic]
	if !ok {
		return
	}

	delete(subs, c)
	if len(subs) == 0 {
		delete(h.topics, topic)
	}
}

// topicCount reports how many topics the connection currently holds.
func (h *hub) topicCount(c *connection) int {
	h.mu.RLock()
	defer h.mu.RUnlock()
	return len(c.topics)
}

// deliverTopic fans one encoded frame out to a topic's subscribers.
//
// The send is non-blocking by design: enqueue returns false when a connection's
// buffer is full, and that connection is then closed as a slow consumer rather
// than allowed to stall the fan-out for everyone else on the topic. The closing
// happens outside the lock — closeSlow tears down the pumps, which call back
// into remove, which wants the same mutex.
func (h *hub) deliverTopic(topic string, frame []byte) {
	h.mu.RLock()
	subs := h.topics[topic]

	var slow []*connection
	for c := range subs {
		if !c.enqueue(frame) {
			slow = append(slow, c)
		}
	}
	h.mu.RUnlock()

	for _, c := range slow {
		c.closeSlow()
	}
}

// closeAll disconnects everyone and refuses new connections. Used on shutdown.
func (h *hub) closeAll() []*connection {
	h.mu.Lock()
	defer h.mu.Unlock()

	if h.closed {
		return nil
	}
	h.closed = true

	open := make([]*connection, 0, len(h.conns))
	for c := range h.conns {
		open = append(open, c)
	}

	h.topics = make(map[string]map[*connection]struct{})
	h.conns = make(map[*connection]struct{})

	return open
}
