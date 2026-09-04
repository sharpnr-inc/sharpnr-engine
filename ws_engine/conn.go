package ws_engine

import (
	"context"
	"encoding/json"
	"errors"
	"log/slog"
	"sync"
	"time"

	"github.com/coder/websocket"
)

// Close codes above 4000 are application-defined. StatusSlowConsumer is the one
// a client must handle: it means the server gave up delivering and the client
// should reconnect and refetch by seq rather than assume it is in sync.
const (
	statusSlowConsumer = websocket.StatusCode(4000)
	statusGoingAway    = websocket.StatusGoingAway
)

// connection is one socket. Exactly one read pump and one write pump run per
// connection, which is also what the underlying library requires: it permits
// one concurrent reader and one concurrent writer, no more.
type connection struct {
	engine *Engine
	ws     *websocket.Conn
	id     Identity

	// send is the outbound queue. Bounded on purpose — see closeSlow.
	send chan []byte

	// topics is this connection's subscription set, guarded by the hub's mutex
	// (not by closeOnce) because it is always touched together with the hub's
	// own maps.
	topics map[string]struct{}

	closeOnce sync.Once
	done      chan struct{}
}

// enqueue offers a frame to the connection without blocking. False means the
// buffer is full and the caller should treat this connection as a slow consumer.
func (c *connection) enqueue(frame []byte) bool {
	select {
	case <-c.done:
		return false
	case c.send <- frame:
		return true
	default:
		return false
	}
}

// closeSlow drops a connection that cannot keep up.
//
// The alternative — growing the buffer — trades a dropped socket for unbounded
// memory held by whoever is furthest behind, and the client can recover on its
// own here: it reconnects, resubscribes, and pulls the gap by seq. So being
// dropped is a normal, recoverable event, not an error path.
func (c *connection) closeSlow() {
	c.shutdown(statusSlowConsumer, "client too slow, resync required")
}

// shutdown is idempotent, safe from either pump, and — importantly — does not
// block the caller.
//
// Unregistering is synchronous, so the moment shutdown returns the connection
// receives no further fan-out. The closing handshake is not: Close waits for
// the peer to acknowledge, which an unresponsive peer never does, and this is
// called from the publisher's goroutine by way of closeSlow. Blocking there
// would let one dead socket stall delivery for everyone else on the topic —
// the exact failure closeSlow exists to prevent.
func (c *connection) shutdown(code websocket.StatusCode, reason string) {
	c.closeOnce.Do(func() {
		close(c.done)
		c.engine.hub.remove(c)

		go func() {
			// Best effort: if the peer never replies, CloseNow releases the
			// socket regardless.
			if err := c.ws.Close(code, reason); err != nil {
				_ = c.ws.CloseNow()
			}
		}()
	})
}

/* -------------------------------------------------------------------------- */
/*                                  Read pump                                 */
/* -------------------------------------------------------------------------- */

// readPump owns the socket's reader for the connection's whole life. It runs on
// the caller's goroutine (the HTTP handler), so Serve blocks until the client
// goes away — which is what keeps the handler from returning and letting the
// framework finalize a response on a hijacked connection.
func (c *connection) readPump(ctx context.Context) {
	defer c.shutdown(websocket.StatusNormalClosure, "")

	// A client frame is a control frame — subscribe, unsubscribe, ping. Content
	// is posted over HTTP, so anything large here is a mistake or an attack.
	c.ws.SetReadLimit(maxClientFrameBytes)

	for {
		_, data, err := c.ws.Read(ctx)
		if err != nil {
			// Normal closure, a dropped network, or our own shutdown racing the
			// read. None of these are worth logging at error level — sockets
			// end all the time.
			return
		}

		var frame clientFrame
		if err := json.Unmarshal(data, &frame); err != nil {
			c.sendError("", codeBadFrame, "frame is not valid json")
			continue
		}

		switch frame.Type {
		case frameSubscribe:
			c.handleSubscribe(ctx, frame.Topics)

		case frameUnsubscribe:
			for _, topic := range frame.Topics {
				c.engine.hub.unsubscribe(c, topic)
			}

		case framePing:
			// Application-level ping, separate from the protocol ping in the
			// write pump: some proxies and mobile stacks hide the latter from
			// the client, and a client that wants to prove liveness itself
			// should be able to.
			c.trySend(serverFrame{Type: framePong})

		default:
			c.sendError("", codeBadFrame, "unknown frame type")
		}
	}
}

// handleSubscribe authorizes each topic independently, so one refused topic
// does not sink the rest of the batch.
func (c *connection) handleSubscribe(ctx context.Context, topics []string) {
	granted := make([]string, 0, len(topics))

	for _, topic := range topics {
		if c.engine.hub.topicCount(c) >= c.engine.cfg.MaxTopicsPerConn {
			c.sendError(topic, codeTooMany, ErrTooManyTopics.Error())
			break
		}

		if err := c.engine.authorize(ctx, c.id, topic); err != nil {
			// Deliberately not echoing the authorizer's error text: it can
			// carry domain detail ("channel not found in org"), and a socket
			// probing topic names should learn nothing from the difference
			// between "does not exist" and "not yours".
			if !errors.Is(err, ErrInvalidTopic) {
				slog.Debug("ws subscribe denied",
					"topic", topic,
					"user_id", c.id.UserID,
					"org_id", c.id.OrgID,
					"error", err,
				)
			}
			c.sendError(topic, codeForbidden, "not allowed to subscribe to this topic")
			continue
		}

		c.engine.hub.subscribe(c, topic)
		granted = append(granted, topic)
	}

	if len(granted) > 0 {
		c.trySend(serverFrame{Type: frameSubscribed, Topics: granted})
	}
}

/* -------------------------------------------------------------------------- */
/*                                 Write pump                                 */
/* -------------------------------------------------------------------------- */

// writePump owns the socket's writer, including the heartbeat: Ping uses the
// writer too, so it cannot live on its own goroutine.
//
// The heartbeat is what detects the connection that is gone without having said
// so — a laptop lid closing, a NAT dropping the flow. Without it those sockets
// sit in the hub forever, holding subscriptions and receiving fan-out nobody
// reads.
func (c *connection) writePump(ctx context.Context) {
	ticker := time.NewTicker(c.engine.cfg.PingInterval)
	defer ticker.Stop()

	for {
		select {
		case <-c.done:
			return

		case <-ctx.Done():
			c.shutdown(statusGoingAway, "server shutting down")
			return

		case frame := <-c.send:
			writeCtx, cancel := context.WithTimeout(ctx, writeWait)
			err := c.ws.Write(writeCtx, websocket.MessageText, frame)
			cancel()
			if err != nil {
				c.shutdown(websocket.StatusInternalError, "write failed")
				return
			}

		case <-ticker.C:
			// Ping blocks until the pong comes back, so this doubles as the
			// liveness check: a timeout means the peer is unreachable even
			// though TCP has not noticed yet.
			pingCtx, cancel := context.WithTimeout(ctx, c.engine.cfg.PongTimeout)
			err := c.ws.Ping(pingCtx)
			cancel()
			if err != nil {
				c.shutdown(websocket.StatusPolicyViolation, "heartbeat timeout")
				return
			}
		}
	}
}

/* -------------------------------------------------------------------------- */
/*                                  Helpers                                   */
/* -------------------------------------------------------------------------- */

// trySend encodes and queues a server frame, dropping it if the queue is full.
// Used for the small control replies, where the connection being backed up
// already means the client has bigger problems than a missing ack.
func (c *connection) trySend(frame serverFrame) {
	encoded, err := json.Marshal(frame)
	if err != nil {
		return
	}
	c.enqueue(encoded)
}

func (c *connection) sendError(topic, code, reason string) {
	c.trySend(serverFrame{
		Type:   frameError,
		Topic:  topic,
		Code:   code,
		Reason: reason,
	})
}
