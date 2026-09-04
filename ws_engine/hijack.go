package ws_engine

import (
	"bufio"
	"errors"
	"net"
	"net/http"
)

// Frameworks that wrap http.ResponseWriter to record the status — gin's is the
// one this app runs — reject Hijack once they believe the header has been
// written:
//
//	func (w *responseWriter) Hijack() (...) {
//		if w.Written() { return nil, nil, errHijackAlreadyWritten }
//
// The websocket handshake trips exactly that. Accept writes 101, then calls the
// wrapper's WriteHeaderNow() to force the status out (its own compatibility
// shim for this family of wrappers), and that call is what marks the response
// written — so the Hijack immediately after it fails.
//
// The failure is quiet and confusing: Accept has already put a valid 101 on the
// wire, so the client completes its handshake and believes it has a socket,
// while the server never took over the connection. Every frame the client sends
// goes into a socket nobody reads, and the 500 Accept writes on its way out
// gets parsed by the client as a websocket frame — surfacing as
// "unexpected rsv bits set" rather than anything resembling the real cause.
//
// unwrapHijacker restores the intended sequence: the wrapper still gets to
// write the header and update its own bookkeeping (so it does not later write a
// second header onto a hijacked connection), while the hijack goes to the
// writer underneath it, which has no such guard.
type unwrapHijacker struct {
	http.ResponseWriter
}

// needsUnwrapHijack reports whether w is one of those wrappers. WriteHeaderNow
// is the marker: it is not part of any net/http interface, and it is precisely
// the method whose contract conflicts with hijacking.
func needsUnwrapHijack(w http.ResponseWriter) bool {
	_, ok := w.(interface{ WriteHeaderNow() })
	return ok
}

// WriteHeaderNow forwards the call Accept makes, so the wrapped writer emits
// the 101 and records that it did.
func (w unwrapHijacker) WriteHeaderNow() {
	if inner, ok := w.ResponseWriter.(interface{ WriteHeaderNow() }); ok {
		inner.WriteHeaderNow()
	}
}

// Hijack goes around the wrapper to the writer it wraps.
//
// Unwrap is the convention net/http's ResponseController uses for exactly this
// — reaching a capability the wrapper does not forward — so a wrapper that
// supports it is unwrapped, and one that does not is asked directly (which
// simply reproduces the error it would have returned anyway).
func (w unwrapHijacker) Hijack() (net.Conn, *bufio.ReadWriter, error) {
	writer := w.ResponseWriter

	for {
		unwrapper, ok := writer.(interface{ Unwrap() http.ResponseWriter })
		if !ok {
			break
		}

		writer = unwrapper.Unwrap()

		if hj, ok := writer.(http.Hijacker); ok {
			return hj.Hijack()
		}
	}

	if hj, ok := w.ResponseWriter.(http.Hijacker); ok {
		return hj.Hijack()
	}

	return nil, nil, errors.New("ws_engine: response writer does not support hijacking")
}
