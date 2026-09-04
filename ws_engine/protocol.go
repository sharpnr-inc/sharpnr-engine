package ws_engine

import "encoding/json"

/* -------------------------------------------------------------------------- */
/*                                   Events                                   */
/* -------------------------------------------------------------------------- */

// Event is one thing that happened, addressed to a topic.
//
// Name is the domain's event name ("message.created", "call.ringing") and Data
// is whatever that domain's client contract says — the engine never looks
// inside it. Publishing marshals Data once and fans the same bytes out to every
// subscriber, so a payload must not be mutated after it is handed over.
type Event struct {
	Topic string `json:"topic"`
	Name  string `json:"event"`
	Data  any    `json:"data,omitempty"`
}

/* -------------------------------------------------------------------------- */
/*                              Frames on the wire                            */
/* -------------------------------------------------------------------------- */

// Frame types. Kept as short strings because every event carries one and the
// client parses it on every message.
const (
	// client -> server
	frameSubscribe   = "sub"
	frameUnsubscribe = "unsub"
	framePing        = "ping"

	// server -> client
	frameEvent      = "event"
	frameSubscribed = "subbed"
	frameError      = "err"
	framePong       = "pong"
)

// clientFrame is everything a client is allowed to send. The read limit on the
// connection is what keeps this cheap to parse: control frames are tiny, and a
// client has no business sending anything large over the socket — messages are
// posted over HTTP so they get the normal authorization, validation and audit
// path.
type clientFrame struct {
	Type   string   `json:"t"`
	Topics []string `json:"topics,omitempty"`
}

// serverFrame is the single envelope for everything going out. Event payloads
// are pre-marshalled into Data so fan-out to N subscribers costs one encode.
type serverFrame struct {
	Type   string          `json:"t"`
	Topic  string          `json:"topic,omitempty"`
	Name   string          `json:"event,omitempty"`
	Data   json.RawMessage `json:"data,omitempty"`
	Topics []string        `json:"topics,omitempty"`
	Code   string          `json:"code,omitempty"`
	Reason string          `json:"reason,omitempty"`
}

// Error codes sent in a frameError. The client is expected to branch on these,
// not on the human-readable reason.
const (
	codeBadFrame  = "bad_frame"
	codeForbidden = "forbidden"
	codeTooMany   = "too_many_topics"
)
