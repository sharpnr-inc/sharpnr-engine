package ws_engine

import (
	"context"
	"encoding/json"
	"errors"
	"log/slog"
	"net/http"
	"strings"
	"sync"
	"time"

	"github.com/coder/websocket"
)

// Defaults. Every one of them is a policy choice, so Config can override.
const (
	// defaultSendBuffer is how far behind a client may fall before it is
	// dropped. Deep enough to ride out a GC pause or a slow frame; shallow
	// enough that a thousand stalled sockets cost megabytes, not gigabytes.
	defaultSendBuffer = 64

	// defaultPingInterval is well under the 60s idle timeout most proxies and
	// load balancers apply, so an open-but-quiet socket is not reaped by
	// something between us and the client.
	defaultPingInterval = 25 * time.Second
	defaultPongTimeout  = 10 * time.Second

	defaultMaxTopicsPerConn = 200

	// maxClientFrameBytes caps an inbound frame. Subscribing to 200 uuid topics
	// at once still fits in a few KB.
	maxClientFrameBytes = 16 * 1024

	// writeWait bounds a single frame write so one unresponsive TCP peer cannot
	// pin a write pump indefinitely.
	writeWait = 10 * time.Second
)

// Backplane carries events between processes.
//
// The in-memory hub only reaches connections held by THIS replica, so a
// multi-replica deployment needs one of these or half the audience silently
// misses every event. No implementation exists yet — Redis pub/sub is the
// intended one, since Redis is already a dependency. The interface is here so
// that adding it later is a wiring change rather than a rewrite of every
// Publish call site.
type Backplane interface {
	// Publish hands the event to the other replicas. The local fan-out has
	// already happened by the time this is called.
	Publish(ctx context.Context, ev Event) error

	// Subscribe delivers events published by OTHER replicas. Implementations
	// must not loop an event back to the replica that published it.
	Subscribe(ctx context.Context, deliver func(Event)) error
}

type Config struct {
	// AllowedOrigins is the exact list of browser origins permitted to open a
	// socket, e.g. "https://cables.sharpnr.com". Required: an empty list
	// refuses every browser handshake rather than accepting any, because a
	// permissive default here is a session-hijacking hole (see the package doc).
	AllowedOrigins []string

	SendBuffer       int
	PingInterval     time.Duration
	PongTimeout      time.Duration
	MaxTopicsPerConn int

	// Backplane is optional; nil means single-process fan-out.
	Backplane Backplane
}

func (c *Config) applyDefaults() {
	if c.SendBuffer <= 0 {
		c.SendBuffer = defaultSendBuffer
	}
	if c.PingInterval <= 0 {
		c.PingInterval = defaultPingInterval
	}
	if c.PongTimeout <= 0 {
		c.PongTimeout = defaultPongTimeout
	}
	if c.MaxTopicsPerConn <= 0 {
		c.MaxTopicsPerConn = defaultMaxTopicsPerConn
	}
}

// Engine is the app's single WebSocket hub. Build one at startup, register the
// topic namespaces each module owns, and inject it wherever events are
// published.
type Engine struct {
	cfg Config
	hub *hub

	authMu      sync.RWMutex
	authorizers map[string]TopicAuthorizer

	// allowedOrigins is the lowercased AllowedOrigins, for comparison.
	allowedOrigins map[string]struct{}

	closeOnce sync.Once
}

func New(cfg Config) *Engine {
	cfg.applyDefaults()

	origins := make(map[string]struct{}, len(cfg.AllowedOrigins))
	for _, o := range cfg.AllowedOrigins {
		origins[strings.ToLower(strings.TrimSpace(o))] = struct{}{}
	}

	return &Engine{
		cfg:            cfg,
		hub:            newHub(),
		authorizers:    make(map[string]TopicAuthorizer),
		allowedOrigins: origins,
	}
}

/* -------------------------------------------------------------------------- */
/*                                 Publishing                                 */
/* -------------------------------------------------------------------------- */

// Publish fans an event out to everyone subscribed to its topic.
//
// It never blocks on a slow client and never returns a delivery error, because
// there is no delivery guarantee to report on (see the package doc). A payload
// that fails to encode is dropped with a log line: that is a programming error
// in the caller's payload type, and failing the caller's transaction over a
// notification would be worse than losing the notification.
//
// Call it AFTER the write it announces has committed. Publishing from inside a
// transaction races the reader: a client can receive "message.created", fetch
// it over HTTP, and find nothing there yet.
func (e *Engine) Publish(ev Event) {
	if ev.Topic == "" || ev.Name == "" {
		slog.Error("ws publish rejected", "reason", "topic and event name are required")
		return
	}

	data, err := json.Marshal(ev.Data)
	if err != nil {
		slog.Error("ws publish dropped", "topic", ev.Topic, "event", ev.Name, "error", err)
		return
	}

	// One encode, N sends: every subscriber on the topic gets the same bytes.
	frame, err := json.Marshal(serverFrame{
		Type:  frameEvent,
		Topic: ev.Topic,
		Name:  ev.Name,
		Data:  data,
	})
	if err != nil {
		slog.Error("ws publish dropped", "topic", ev.Topic, "event", ev.Name, "error", err)
		return
	}

	e.hub.deliverTopic(ev.Topic, frame)

	if e.cfg.Backplane != nil {
		if err := e.cfg.Backplane.Publish(context.Background(), ev); err != nil {
			slog.Error("ws backplane publish failed", "topic", ev.Topic, "event", ev.Name, "error", err)
		}
	}
}

// deliverLocal fans out an event that arrived from another replica. It skips
// the backplane, so events cannot echo between processes.
func (e *Engine) deliverLocal(ev Event) {
	data, err := json.Marshal(ev.Data)
	if err != nil {
		return
	}
	frame, err := json.Marshal(serverFrame{
		Type:  frameEvent,
		Topic: ev.Topic,
		Name:  ev.Name,
		Data:  data,
	})
	if err != nil {
		return
	}
	e.hub.deliverTopic(ev.Topic, frame)
}

// StartBackplane begins consuming events published by other replicas. No-op
// when no backplane is configured. Blocks until ctx is cancelled, so run it in
// its own goroutine at startup.
func (e *Engine) StartBackplane(ctx context.Context) error {
	if e.cfg.Backplane == nil {
		return nil
	}
	return e.cfg.Backplane.Subscribe(ctx, e.deliverLocal)
}

/* -------------------------------------------------------------------------- */
/*                              Accepting sockets                             */
/* -------------------------------------------------------------------------- */

// ErrForbiddenOrigin is returned when a handshake's Origin is not in the
// configured allow-list.
var ErrForbiddenOrigin = errors.New("ws_engine: origin not allowed")

// Serve upgrades an authenticated HTTP request into a socket and runs it until
// the client disconnects or the engine shuts down.
//
// Identity must already be resolved — this is mounted behind the normal auth
// middleware, and the engine does no authentication of its own. It BLOCKS for
// the life of the connection.
func (e *Engine) Serve(w http.ResponseWriter, r *http.Request, id Identity) {
	if err := e.checkOrigin(r); err != nil {
		http.Error(w, "forbidden origin", http.StatusForbidden)
		return
	}

	ws, err := websocket.Accept(w, r, &websocket.AcceptOptions{
		// The library's own origin check compares hosts and ignores the scheme.
		// checkOrigin above already compared the full origin against the app's
		// configured list, which is stricter, so this disables the weaker
		// duplicate rather than skipping verification.
		InsecureSkipVerify: true,
		CompressionMode:    websocket.CompressionDisabled,
	})
	if err != nil {
		// Accept has already written the error response.
		return
	}

	conn := &connection{
		engine: e,
		ws:     ws,
		id:     id,
		send:   make(chan []byte, e.cfg.SendBuffer),
		done:   make(chan struct{}),
	}

	if !e.hub.add(conn) {
		_ = ws.Close(statusGoingAway, "server shutting down")
		return
	}

	// The pumps share one context so either ending tears down the other: the
	// write pump's blocking Ping and the read pump's blocking Read both need a
	// cancellation to return from.
	ctx, cancel := context.WithCancel(r.Context())
	defer cancel()

	go conn.writePump(ctx)
	conn.readPump(ctx) // blocks until the socket ends
}

// checkOrigin enforces the allow-list. A handshake with no Origin header is a
// non-browser client (a Go test, a native app, curl) — those are not subject to
// the ambient-cookie problem this check exists for, but they still had to pass
// the auth middleware to get here.
func (e *Engine) checkOrigin(r *http.Request) error {
	origin := r.Header.Get("Origin")
	if origin == "" {
		return nil
	}

	if _, ok := e.allowedOrigins[strings.ToLower(origin)]; !ok {
		slog.Warn("ws handshake rejected",
			"reason", "forbidden_origin",
			"origin", origin,
			"ip", r.RemoteAddr,
		)
		return ErrForbiddenOrigin
	}
	return nil
}

/* -------------------------------------------------------------------------- */
/*                                  Shutdown                                  */
/* -------------------------------------------------------------------------- */

// Close disconnects every socket and refuses new ones. Safe to call twice, and
// safe to defer next to the other engines in main.
func (e *Engine) Close() {
	e.closeOnce.Do(func() {
		for _, c := range e.hub.closeAll() {
			c.shutdown(statusGoingAway, "server shutting down")
		}
	})
}
