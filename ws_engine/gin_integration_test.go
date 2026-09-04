package ws_engine

import (
	"context"
	"encoding/json"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/coder/websocket"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

// The engine is mounted behind gin in the real app, not on a bare net/http
// handler like the rest of these tests. Gin wraps the ResponseWriter, and a
// websocket upgrade hijacks it — so this exercises the combination the app
// actually runs, which is where an upgrade can succeed while no frame ever
// reaches the reader.
func TestServeBehindGin(t *testing.T) {
	gin.SetMode(gin.TestMode)

	e := New(Config{})
	e.RegisterAuthorizer("conv", allowAll())
	defer e.Close()

	router := gin.New()
	router.GET("/ws", func(c *gin.Context) {
		e.Serve(c.Writer, c.Request, Identity{UserID: uuid.New(), OrgID: uuid.New()})
	})

	srv := httptest.NewServer(router)
	defer srv.Close()

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	conn, _, err := websocket.Dial(ctx, "ws"+srv.URL[len("http"):]+"/ws", nil)
	if err != nil {
		t.Fatalf("dial: %v", err)
	}
	defer conn.CloseNow()

	topic := "conv:" + uuid.NewString()
	sub, _ := json.Marshal(clientFrame{Type: frameSubscribe, Topics: []string{topic}})
	if err := conn.Write(ctx, websocket.MessageText, sub); err != nil {
		t.Fatalf("write subscribe: %v", err)
	}

	readCtx, readCancel := context.WithTimeout(ctx, 5*time.Second)
	defer readCancel()

	_, data, err := conn.Read(readCtx)
	if err != nil {
		t.Fatalf("no reply to subscribe behind gin: %v", err)
	}

	var ack serverFrame
	if err := json.Unmarshal(data, &ack); err != nil {
		t.Fatalf("unmarshal: %v", err)
	}
	if ack.Type != frameSubscribed {
		t.Fatalf("expected %q, got %+v", frameSubscribed, ack)
	}

	// And an event published after the ack must reach it.
	e.Publish(Event{Topic: topic, Name: "message.created", Data: map[string]string{"body": "hi"}})

	eventCtx, eventCancel := context.WithTimeout(ctx, 5*time.Second)
	defer eventCancel()

	_, data, err = conn.Read(eventCtx)
	if err != nil {
		t.Fatalf("no event behind gin: %v", err)
	}

	var got serverFrame
	if err := json.Unmarshal(data, &got); err != nil {
		t.Fatalf("unmarshal event: %v", err)
	}
	if got.Type != frameEvent || got.Name != "message.created" {
		t.Fatalf("unexpected event frame: %+v", got)
	}
}
