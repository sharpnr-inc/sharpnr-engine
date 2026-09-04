package ws_engine

import (
	"context"
	"errors"
	"strings"

	"github.com/google/uuid"
)

// Identity is who is on the other end of the socket, resolved by the HTTP
// middleware during the handshake. The engine carries it, compares nothing
// against it, and hands it to the authorizer.
type Identity struct {
	UserID uuid.UUID
	OrgID  uuid.UUID
}

// TopicAuthorizer decides whether one identity may subscribe to one topic.
//
// It is the only place a domain's rules enter the engine. Cables answers it
// from channel membership, kitchen from property access. Returning any error
// denies the subscription; the connection stays open, because a client asking
// for a channel it just left is normal, not hostile.
//
// Called on every subscribe, so it must be cheap — cache in the implementation
// if the underlying read is not.
type TopicAuthorizer interface {
	CanSubscribe(ctx context.Context, id Identity, topic string) error
}

// AuthorizerFunc adapts a function to TopicAuthorizer.
type AuthorizerFunc func(ctx context.Context, id Identity, topic string) error

func (f AuthorizerFunc) CanSubscribe(ctx context.Context, id Identity, topic string) error {
	return f(ctx, id, topic)
}

var (
	// ErrNoAuthorizer is a deny, not a misconfiguration the caller can ignore:
	// a namespace nobody claimed is refused rather than allowed. Registering an
	// authorizer is how a module opts its topics into existence.
	ErrNoAuthorizer = errors.New("ws_engine: no authorizer registered for topic namespace")

	ErrInvalidTopic = errors.New("ws_engine: malformed topic")

	// ErrTooManyTopics caps a single connection's subscription set. One socket
	// asking for fifty thousand topics is either a bug or an attempt to make
	// the fan-out map expensive.
	ErrTooManyTopics = errors.New("ws_engine: subscription limit reached for this connection")
)

// A topic is "<namespace>:<rest>", e.g. "conv:9f3c...". The namespace is what
// an authorizer registers against; the rest is opaque to the engine.
const topicSeparator = ":"

// maxTopicLen is a parsing guard, not a policy — a uuid-shaped topic is ~40
// bytes and nothing legitimate approaches this.
const maxTopicLen = 128

// splitTopic returns the namespace of a well-formed topic.
func splitTopic(topic string) (namespace string, err error) {
	if topic == "" || len(topic) > maxTopicLen {
		return "", ErrInvalidTopic
	}

	ns, rest, found := strings.Cut(topic, topicSeparator)
	if !found || ns == "" || rest == "" {
		return "", ErrInvalidTopic
	}

	return ns, nil
}

// RegisterAuthorizer claims a topic namespace for a module.
//
// Bootstrap-time only: the engine is built in main before the modules that own
// topics exist, so registration cannot happen in the constructor. It is
// mutex-guarded anyway, because getting this wrong would be a data race in
// something that never crashes visibly.
//
// Registering a namespace twice panics rather than silently replacing the
// previous owner — two modules fighting over "conv:" is a wiring bug worth
// failing at startup.
func (e *Engine) RegisterAuthorizer(namespace string, a TopicAuthorizer) {
	if namespace == "" || strings.Contains(namespace, topicSeparator) {
		panic("ws_engine: namespace must be non-empty and contain no " + topicSeparator)
	}
	if a == nil {
		panic("ws_engine: nil authorizer for namespace " + namespace)
	}

	e.authMu.Lock()
	defer e.authMu.Unlock()

	if _, exists := e.authorizers[namespace]; exists {
		panic("ws_engine: namespace already registered: " + namespace)
	}
	e.authorizers[namespace] = a
}

// authorize resolves the topic's namespace to its owner and asks it.
func (e *Engine) authorize(ctx context.Context, id Identity, topic string) error {
	namespace, err := splitTopic(topic)
	if err != nil {
		return err
	}

	e.authMu.RLock()
	authorizer, ok := e.authorizers[namespace]
	e.authMu.RUnlock()

	if !ok {
		return ErrNoAuthorizer
	}

	return authorizer.CanSubscribe(ctx, id, topic)
}
