package ws_engine

import (
	"context"
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"sync"
	"testing"
	"time"

	"github.com/coder/websocket"
	"github.com/google/uuid"
)

/* -------------------------------------------------------------------------- */
/*                                  Harness                                   */
/* -------------------------------------------------------------------------- */

var testIdentity = Identity{UserID: uuid.New(), OrgID: uuid.New()}

// allowAll / denyAll stand in for a domain module's authorizer.
func allowAll() TopicAuthorizer {
	return AuthorizerFunc(func(context.Context, Identity, string) error { return nil })
}

func denyAll() TopicAuthorizer {
	return AuthorizerFunc(func(context.Context, Identity, string) error {
		return errors.New("nope")
	})
}

// newTestServer wires the engine behind an httptest server, standing in for the
// authenticated route.
func newTestServer(t *testing.T, e *Engine) *httptest.Server {
	t.Helper()

	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		e.Serve(w, r, testIdentity)
	}))
	t.Cleanup(srv.Close)
	return srv
}

func dial(t *testing.T, srv *httptest.Server) *websocket.Conn {
	t.Helper()

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	conn, _, err := websocket.Dial(ctx, "ws"+srv.URL[len("http"):], nil)
	if err != nil {
		t.Fatalf("dial: %v", err)
	}
	t.Cleanup(func() { _ = conn.CloseNow() })
	return conn
}

func writeFrame(t *testing.T, conn *websocket.Conn, frame clientFrame) {
	t.Helper()

	data, err := json.Marshal(frame)
	if err != nil {
		t.Fatalf("marshal frame: %v", err)
	}

	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()

	if err := conn.Write(ctx, websocket.MessageText, data); err != nil {
		t.Fatalf("write frame: %v", err)
	}
}

func readFrame(t *testing.T, conn *websocket.Conn) serverFrame {
	t.Helper()

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	_, data, err := conn.Read(ctx)
	if err != nil {
		t.Fatalf("read frame: %v", err)
	}

	var frame serverFrame
	if err := json.Unmarshal(data, &frame); err != nil {
		t.Fatalf("unmarshal frame: %v", err)
	}
	return frame
}

/* -------------------------------------------------------------------------- */
/*                                   Topics                                   */
/* -------------------------------------------------------------------------- */

func TestSplitTopic(t *testing.T) {
	tests := []struct {
		name    string
		topic   string
		want    string
		wantErr bool
	}{
		{name: "namespaced", topic: "conv:" + uuid.NewString(), want: "conv"},
		{name: "rest may contain separators", topic: "call:a:b", want: "call"},
		{name: "empty", topic: "", wantErr: true},
		{name: "no separator", topic: "conv", wantErr: true},
		{name: "empty namespace", topic: ":abc", wantErr: true},
		{name: "empty remainder", topic: "conv:", wantErr: true},
		{name: "over length", topic: "conv:" + string(make([]byte, maxTopicLen)), wantErr: true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := splitTopic(tt.topic)
			if tt.wantErr {
				if err == nil {
					t.Fatalf("expected error for %q", tt.topic)
				}
				return
			}
			if err != nil {
				t.Fatalf("unexpected error: %v", err)
			}
			if got != tt.want {
				t.Fatalf("namespace = %q, want %q", got, tt.want)
			}
		})
	}
}

// An unclaimed namespace must deny. This is the fail-closed guarantee: a module
// that forgets to register its authorizer gets silence, not a public topic.
func TestUnregisteredNamespaceIsDenied(t *testing.T) {
	e := New(Config{})

	err := e.authorize(context.Background(), testIdentity, "conv:"+uuid.NewString())
	if !errors.Is(err, ErrNoAuthorizer) {
		t.Fatalf("expected ErrNoAuthorizer, got %v", err)
	}
}

func TestRegisterAuthorizerRejectsDuplicate(t *testing.T) {
	e := New(Config{})
	e.RegisterAuthorizer("conv", allowAll())

	defer func() {
		if recover() == nil {
			t.Fatal("expected a panic when a namespace is registered twice")
		}
	}()
	e.RegisterAuthorizer("conv", allowAll())
}

/* -------------------------------------------------------------------------- */
/*                            Subscribe & delivery                            */
/* -------------------------------------------------------------------------- */

func TestSubscribeThenReceivePublishedEvent(t *testing.T) {
	e := New(Config{})
	e.RegisterAuthorizer("conv", allowAll())
	defer e.Close()

	srv := newTestServer(t, e)
	conn := dial(t, srv)

	topic := "conv:" + uuid.NewString()
	writeFrame(t, conn, clientFrame{Type: frameSubscribe, Topics: []string{topic}})

	if ack := readFrame(t, conn); ack.Type != frameSubscribed {
		t.Fatalf("expected %q ack, got %q", frameSubscribed, ack.Type)
	}

	// Publish until it lands: the ack proves the server processed the
	// subscribe, so a single publish would already be enough, but the retry
	// keeps the test from depending on goroutine scheduling.
	done := make(chan serverFrame, 1)
	go func() { done <- readFrame(t, conn) }()

	deadline := time.After(3 * time.Second)
	for {
		e.Publish(Event{Topic: topic, Name: "message.created", Data: map[string]string{"body": "hi"}})

		select {
		case got := <-done:
			if got.Type != frameEvent || got.Name != "message.created" || got.Topic != topic {
				t.Fatalf("unexpected frame: %+v", got)
			}
			var payload map[string]string
			if err := json.Unmarshal(got.Data, &payload); err != nil {
				t.Fatalf("payload: %v", err)
			}
			if payload["body"] != "hi" {
				t.Fatalf("payload = %v", payload)
			}
			return
		case <-deadline:
			t.Fatal("event never arrived")
		case <-time.After(20 * time.Millisecond):
		}
	}
}

func TestDeniedSubscribeGetsErrorAndNoEvents(t *testing.T) {
	e := New(Config{})
	e.RegisterAuthorizer("conv", denyAll())
	defer e.Close()

	srv := newTestServer(t, e)
	conn := dial(t, srv)

	topic := "conv:" + uuid.NewString()
	writeFrame(t, conn, clientFrame{Type: frameSubscribe, Topics: []string{topic}})

	got := readFrame(t, conn)
	if got.Type != frameError || got.Code != codeForbidden {
		t.Fatalf("expected forbidden error, got %+v", got)
	}

	// The connection survives a denial — a client asking for a channel it just
	// left is normal.
	writeFrame(t, conn, clientFrame{Type: framePing})
	if pong := readFrame(t, conn); pong.Type != framePong {
		t.Fatalf("expected pong after denial, got %+v", pong)
	}
}

// A refused topic must not sink the rest of the batch.
func TestPartialSubscribeGrantsWhatItCan(t *testing.T) {
	allowed := "conv:" + uuid.NewString()

	e := New(Config{})
	e.RegisterAuthorizer("conv", AuthorizerFunc(func(_ context.Context, _ Identity, topic string) error {
		if topic == allowed {
			return nil
		}
		return errors.New("nope")
	}))
	defer e.Close()

	srv := newTestServer(t, e)
	conn := dial(t, srv)

	writeFrame(t, conn, clientFrame{
		Type:   frameSubscribe,
		Topics: []string{"conv:" + uuid.NewString(), allowed},
	})

	if errFrame := readFrame(t, conn); errFrame.Type != frameError {
		t.Fatalf("expected the denial first, got %+v", errFrame)
	}

	ack := readFrame(t, conn)
	if ack.Type != frameSubscribed || len(ack.Topics) != 1 || ack.Topics[0] != allowed {
		t.Fatalf("expected an ack for the allowed topic only, got %+v", ack)
	}
}

func TestUnsubscribeStopsDelivery(t *testing.T) {
	e := New(Config{})
	e.RegisterAuthorizer("conv", allowAll())
	defer e.Close()

	srv := newTestServer(t, e)
	conn := dial(t, srv)

	topic := "conv:" + uuid.NewString()
	writeFrame(t, conn, clientFrame{Type: frameSubscribe, Topics: []string{topic}})
	readFrame(t, conn) // ack

	writeFrame(t, conn, clientFrame{Type: frameUnsubscribe, Topics: []string{topic}})

	// Round-trip a ping so the unsubscribe is definitely processed — frames on
	// one connection are handled in order.
	writeFrame(t, conn, clientFrame{Type: framePing})
	if pong := readFrame(t, conn); pong.Type != framePong {
		t.Fatalf("expected pong, got %+v", pong)
	}

	e.Publish(Event{Topic: topic, Name: "message.created"})

	// Nothing should arrive; a short read timeout is the assertion.
	ctx, cancel := context.WithTimeout(context.Background(), 300*time.Millisecond)
	defer cancel()
	if _, _, err := conn.Read(ctx); err == nil {
		t.Fatal("received an event after unsubscribing")
	}
}

func TestMaxTopicsPerConnection(t *testing.T) {
	e := New(Config{MaxTopicsPerConn: 2})
	e.RegisterAuthorizer("conv", allowAll())
	defer e.Close()

	srv := newTestServer(t, e)
	conn := dial(t, srv)

	writeFrame(t, conn, clientFrame{Type: frameSubscribe, Topics: []string{
		"conv:" + uuid.NewString(),
		"conv:" + uuid.NewString(),
		"conv:" + uuid.NewString(),
	}})

	got := readFrame(t, conn)
	if got.Type != frameError || got.Code != codeTooMany {
		t.Fatalf("expected the limit error, got %+v", got)
	}
}

func TestMalformedFrameIsRejectedWithoutClosing(t *testing.T) {
	e := New(Config{})
	defer e.Close()

	srv := newTestServer(t, e)
	conn := dial(t, srv)

	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()
	if err := conn.Write(ctx, websocket.MessageText, []byte("{not json")); err != nil {
		t.Fatalf("write: %v", err)
	}

	if got := readFrame(t, conn); got.Type != frameError || got.Code != codeBadFrame {
		t.Fatalf("expected bad_frame, got %+v", got)
	}
}

/* -------------------------------------------------------------------------- */
/*                                   Origin                                   */
/* -------------------------------------------------------------------------- */

func TestOriginAllowList(t *testing.T) {
	e := New(Config{AllowedOrigins: []string{"https://cables.sharpnr.com"}})

	tests := []struct {
		name    string
		origin  string
		wantErr bool
	}{
		{name: "allowed", origin: "https://cables.sharpnr.com"},
		{name: "case insensitive", origin: "https://CABLES.sharpnr.com"},
		{name: "absent is a non-browser client", origin: ""},
		{name: "other origin", origin: "https://evil.example", wantErr: true},
		{name: "scheme matters", origin: "http://cables.sharpnr.com", wantErr: true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := httptest.NewRequest(http.MethodGet, "/ws", nil)
			if tt.origin != "" {
				r.Header.Set("Origin", tt.origin)
			}

			err := e.checkOrigin(r)
			if tt.wantErr != (err != nil) {
				t.Fatalf("origin %q: err = %v, wantErr = %v", tt.origin, err, tt.wantErr)
			}
		})
	}
}

// An engine with no configured origins refuses every browser handshake rather
// than defaulting open.
func TestEmptyOriginListRefusesBrowsers(t *testing.T) {
	e := New(Config{})

	r := httptest.NewRequest(http.MethodGet, "/ws", nil)
	r.Header.Set("Origin", "https://cables.sharpnr.com")

	if err := e.checkOrigin(r); !errors.Is(err, ErrForbiddenOrigin) {
		t.Fatalf("expected ErrForbiddenOrigin, got %v", err)
	}
}

/* -------------------------------------------------------------------------- */
/*                          Lifecycle & slow consumers                        */
/* -------------------------------------------------------------------------- */

// A client that never reads must be dropped rather than buffered forever.
func TestSlowConsumerIsDisconnected(t *testing.T) {
	e := New(Config{SendBuffer: 1})
	e.RegisterAuthorizer("conv", allowAll())
	defer e.Close()

	srv := newTestServer(t, e)
	conn := dial(t, srv)

	topic := "conv:" + uuid.NewString()
	writeFrame(t, conn, clientFrame{Type: frameSubscribe, Topics: []string{topic}})
	readFrame(t, conn) // ack, the last thing this client will read

	// Outrun the write pump. The buffer is 1, so this overflows quickly.
	for i := 0; i < 5000; i++ {
		e.Publish(Event{Topic: topic, Name: "message.created", Data: i})
	}

	// The connection is closed from the server side; the client sees it on its
	// next read, once it has drained whatever did make it through.
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	for {
		if _, _, err := conn.Read(ctx); err != nil {
			var closeErr websocket.CloseError
			if errors.As(err, &closeErr) && closeErr.Code != statusSlowConsumer {
				t.Fatalf("closed with %v, want %v", closeErr.Code, statusSlowConsumer)
			}
			return
		}
	}
}

func TestCloseDisconnectsEveryone(t *testing.T) {
	e := New(Config{})
	e.RegisterAuthorizer("conv", allowAll())

	srv := newTestServer(t, e)
	conn := dial(t, srv)

	topic := "conv:" + uuid.NewString()
	writeFrame(t, conn, clientFrame{Type: frameSubscribe, Topics: []string{topic}})
	readFrame(t, conn) // ack

	e.Close()
	e.Close() // idempotent

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()
	if _, _, err := conn.Read(ctx); err == nil {
		t.Fatal("expected the socket to be closed by shutdown")
	}
}

// Publishing to a topic nobody holds, and publishing after shutdown, are no-ops
// rather than panics — a domain service should never have to ask whether the
// socket layer is up before committing a write.
func TestPublishIsSafeWithNoSubscribersAndAfterClose(t *testing.T) {
	e := New(Config{})
	e.Publish(Event{Topic: "conv:" + uuid.NewString(), Name: "message.created"})

	e.Close()
	e.Publish(Event{Topic: "conv:" + uuid.NewString(), Name: "message.created"})
}

func TestPublishRejectsIncompleteEvents(t *testing.T) {
	e := New(Config{})
	defer e.Close()

	// Neither should panic or fan anything out.
	e.Publish(Event{Name: "message.created"})
	e.Publish(Event{Topic: "conv:x"})
}

/* -------------------------------------------------------------------------- */
/*                                   Races                                    */
/* -------------------------------------------------------------------------- */

// The teardown paths are the ones worth racing: subscribe, publish, and a
// disconnect all touching the same maps. Run with -race.
func TestConcurrentSubscribePublishDisconnect(t *testing.T) {
	e := New(Config{SendBuffer: 4})
	e.RegisterAuthorizer("conv", allowAll())
	defer e.Close()

	srv := newTestServer(t, e)

	topic := "conv:" + uuid.NewString()

	var wg sync.WaitGroup
	for i := 0; i < 8; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()

			conn := dial(t, srv)
			writeFrame(t, conn, clientFrame{Type: frameSubscribe, Topics: []string{topic}})

			// Drop the socket while events are in flight.
			time.Sleep(20 * time.Millisecond)
			_ = conn.CloseNow()
		}()
	}

	stop := make(chan struct{})
	var pub sync.WaitGroup
	pub.Add(1)
	go func() {
		defer pub.Done()
		for {
			select {
			case <-stop:
				return
			default:
				e.Publish(Event{Topic: topic, Name: "message.created", Data: "x"})
			}
		}
	}()

	wg.Wait()
	close(stop)
	pub.Wait()
}
