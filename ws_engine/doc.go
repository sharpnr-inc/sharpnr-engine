// Package ws_engine is the app's WebSocket transport: it owns connections and
// topic fan-out, and nothing else.
//
// # What lives here and what does not
//
// The engine knows about connections, topics and bytes. It does not know what
// a channel, a message or a call is. A domain module (cables first, kitchen
// later) owns three things the engine asks it for:
//
//   - the topic names it publishes on ("conv:<uuid>", "call:<uuid>")
//   - a TopicAuthorizer answering "may this identity subscribe to this topic?"
//   - the event payloads, published from the same place the write is committed
//
// That split is why this is a top-level engine rather than a file inside
// cables_engine: a user holds ONE socket for the whole app, and the kitchen
// board must not import the cables module to push a KOT update. It is the same
// arrangement sse_engine already has with kitchen.
//
// # Delivery: best-effort on purpose
//
// This is the one semantic worth understanding before using it. The socket
// makes no delivery guarantee at all. A connection that cannot keep up is
// closed with StatusSlowConsumer rather than buffered without bound.
//
// Correctness comes from the caller's own sequence numbers, not from this
// engine: cables_messages.seq is per-conversation and monotonic, so a client
// that sees a gap — or that reconnects after being dropped — refetches the
// delta over plain HTTP with the cursor it already has. Delivering a message
// twice is therefore harmless and losing one is recoverable, which is what
// lets the engine stay this small.
//
// Do not add an ack/retry layer here. If a payload cannot be dropped safely,
// it does not belong on the socket.
//
// # Single process, for now
//
// The default fan-out is in-memory, so it reaches only the connections held by
// THIS process. A second replica behind a load balancer would split the
// audience in half without erroring. Config.Backplane is the seam that fixes
// that (Redis pub/sub is the obvious implementation, and Redis is already in
// the stack); no implementation ships yet, deliberately. Publish is a method on
// Engine rather than direct hub access precisely so adding one changes no
// caller.
//
// # Security
//
// WebSocket handshakes are NOT governed by CORS, and the browser attaches the
// session cookie to them like any other request. Any origin can therefore open
// a socket as a logged-in user unless the server checks Origin itself, which
// checkOrigin does against the app's configured web origins. That check is the
// only thing standing between this endpoint and cross-site socket hijacking —
// treat it as load-bearing.
package ws_engine
