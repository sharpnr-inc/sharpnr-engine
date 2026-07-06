package sse_engine

import (
	"fmt"
	"sync"
	"testing"
	"time"

	"github.com/go-openapi/testify/v2/assert"
	"github.com/go-openapi/testify/v2/require"
)

// receive reads one event or fails the test after a timeout.
func receive(t *testing.T, sub *Subscription) Event {
	t.Helper()
	select {
	case ev, ok := <-sub.C():
		require.True(t, ok, "subscription channel closed unexpectedly")
		return ev
	case <-time.After(2 * time.Second):
		t.Fatal("timed out waiting for event")
		return Event{}
	}
}

func TestPublishReachesSubscribers(t *testing.T) {
	sm := NewStreamManager()
	defer sm.Close()

	subA := sm.Subscribe("kitchen_1")
	defer subA.Close()
	subB := sm.Subscribe("kitchen_1")
	defer subB.Close()
	other := sm.Subscribe("kitchen_2")
	defer other.Close()

	sm.Publish("kitchen_1", "hello")

	for _, sub := range []*Subscription{subA, subB} {
		ev := receive(t, sub)
		assert.Equal(t, "kitchen_1", ev.Topic)
		assert.Equal(t, "hello", ev.Payload)
	}

	// The other topic must stay quiet.
	select {
	case ev := <-other.C():
		t.Fatalf("kitchen_2 received stray event: %+v", ev)
	case <-time.After(50 * time.Millisecond):
	}
}

// A subscriber that never drains its channel must end up with the newest
// event, not the oldest — the latest-value slot replaces stale snapshots.
func TestSlowSubscriberGetsLatestSnapshot(t *testing.T) {
	sm := NewStreamManager()
	defer sm.Close()

	sub := sm.Subscribe("kitchen_1")
	defer sub.Close()

	const last = 50
	for i := 1; i <= last; i++ {
		sm.Publish("kitchen_1", i)
	}

	// Fan-out is async: keep reading until the final snapshot arrives.
	deadline := time.After(2 * time.Second)
	for {
		select {
		case ev := <-sub.C():
			if ev.Payload == last {
				return
			}
		case <-deadline:
			t.Fatalf("never received final snapshot %d", last)
		}
	}
}

func TestUnsubscribeIsIdempotent(t *testing.T) {
	sm := NewStreamManager()
	defer sm.Close()

	// A second subscriber keeps the topic alive so the double Close hits
	// the guarded path rather than the missing-topic path.
	stays := sm.Subscribe("kitchen_1")
	defer stays.Close()

	sub := sm.Subscribe("kitchen_1")
	sub.Close()
	sub.Close() // must not panic on double close

	_, ok := <-sub.C()
	assert.False(t, ok, "channel should be closed after Close")

	// The surviving subscriber still receives events.
	sm.Publish("kitchen_1", "still alive")
	assert.Equal(t, "still alive", receive(t, stays).Payload)
}

func TestManagerClose(t *testing.T) {
	sm := NewStreamManager()

	sub := sm.Subscribe("kitchen_1")
	sm.Close()
	sm.Close() // idempotent

	_, ok := <-sub.C()
	assert.False(t, ok, "subscriber channel should close when manager closes")

	// After Close, Publish/Subscribe are safe no-ops.
	sm.Publish("kitchen_1", "ignored")
	late := sm.Subscribe("kitchen_1")
	_, ok = <-late.C()
	assert.False(t, ok, "post-Close subscription should be born closed")
	late.Close()
	sub.Close()
}

// Exercises publish, subscribe, unsubscribe and manager close concurrently;
// meaningful under `go test -race`.
func TestConcurrentChurn(t *testing.T) {
	sm := NewStreamManager()

	var wg sync.WaitGroup
	for g := range 8 {
		wg.Add(1)
		go func(g int) {
			defer wg.Done()
			topic := fmt.Sprintf("kitchen_%d", g%3)
			for i := range 200 {
				sub := sm.Subscribe(topic)
				sm.Publish(topic, i)
				select {
				case <-sub.C():
				default:
				}
				sub.Close()
			}
		}(g)
	}
	wg.Wait()
	sm.Close()
}
