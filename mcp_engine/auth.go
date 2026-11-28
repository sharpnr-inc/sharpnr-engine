package mcp_engine

import (
	"context"
	"errors"
	"fmt"
	"log/slog"
	"net/http"
	"time"

	jwt_manager "sharpnr-engine/pkg/jwtmanager"
	sentinelContexts "sharpnr-engine/sentinel_contexts"

	"github.com/google/uuid"
	"github.com/modelcontextprotocol/go-sdk/auth"
	"github.com/modelcontextprotocol/go-sdk/mcp"
)

// Principal is who a tool call acts for. It travels from the token verifier
// to the tool handlers inside auth.TokenInfo.Extra.
type Principal struct {
	EmployeeID uuid.UUID
	OrgID      uuid.UUID
	IPAddress  string
}

// principalKey is the auth.TokenInfo.Extra key the Principal is stored under.
const principalKey = "sharpnr.principal"

// ErrUnauthenticated is returned by a tool called without a verified token.
// RequireBearerToken rejects such requests before they reach a tool, so seeing
// this means the handler was mounted without the auth middleware.
var ErrUnauthenticated = errors.New("mcp: request has no verified principal")

// AccessTokenVerifier verifies the engine's own mobile access tokens.
//
// Temporary: authentication is moving to sharp-account. Its tokens do not yet
// carry an org or a per-API audience, so they cannot be mapped to a Sentinel
// principal. When they can, add a sharp-account verifier (JWKS, iss, aud =
// this server's URL) and pass it to New instead; the tools do not change.
//
// The error text reaches the client verbatim, so it is always the generic
// auth.ErrInvalidToken; the reason is logged here instead.
func AccessTokenVerifier(jwtMgr jwt_manager.Manager) auth.TokenVerifier {
	return func(ctx context.Context, token string, r *http.Request) (*auth.TokenInfo, error) {
		claims, err := jwtMgr.VerifyAccess(token)
		if err != nil {
			slog.Warn("mcp auth failed", "reason", "invalid_token", "error", err.Error())
			return nil, auth.ErrInvalidToken
		}

		if claims.UserID == uuid.Nil || claims.OrgID == uuid.Nil {
			slog.Warn("mcp auth failed", "reason", "token_missing_identity")
			return nil, auth.ErrInvalidToken
		}

		var expiry time.Time
		if claims.ExpiresAt != nil {
			expiry = claims.ExpiresAt.Time
		}

		return newTokenInfo(Principal{
			EmployeeID: claims.UserID,
			OrgID:      claims.OrgID,
			IPAddress:  ClientIPFromContext(r.Context()),
		}, expiry), nil
	}
}

// newTokenInfo packs a Principal the way principalFrom expects to find it.
//
// UserID binds a stateful MCP session to the org and employee that opened it,
// so a session id leaked to someone else is useless with their own token.
func newTokenInfo(p Principal, expiry time.Time) *auth.TokenInfo {
	return &auth.TokenInfo{
		UserID:     fmt.Sprintf("%s:%s", p.OrgID, p.EmployeeID),
		Expiration: expiry,
		Extra:      map[string]any{principalKey: p},
	}
}

// principalFrom returns the Principal of a tool call.
func principalFrom(req *mcp.CallToolRequest) (Principal, error) {
	if req == nil || req.Extra == nil || req.Extra.TokenInfo == nil {
		return Principal{}, ErrUnauthenticated
	}

	p, ok := req.Extra.TokenInfo.Extra[principalKey].(Principal)
	if !ok {
		return Principal{}, ErrUnauthenticated
	}

	return p, nil
}

// authorize resolves the caller and returns a context carrying the Sentinel
// facts for this request, with channel "mcp". Tool handlers call it first and
// pass the returned context on to services.
//
// The SDK runs tool handlers on the session's context, not the HTTP request's,
// so nothing the request carried reaches the handler except auth.TokenInfo.
// That is why the Sentinel context is stamped here rather than in middleware.
func authorize(ctx context.Context, req *mcp.CallToolRequest) (context.Context, Principal, error) {
	p, err := principalFrom(req)
	if err != nil {
		return ctx, Principal{}, err
	}

	ctx = sentinelContexts.Inject(ctx, sentinelContexts.AuthzContext{
		Request: sentinelContexts.RequestContext{
			Channel:     sentinelContexts.ChannelMCP,
			IPAddress:   p.IPAddress,
			RequestedAt: time.Now(),
		},
	})

	return ctx, p, nil
}

/* ---------------------------------- <***> --------------------------------- */

type clientIPKey struct{}

// WithClientIP records the caller's IP on a request context. The routes layer
// sets it from gin's ClientIP, which honours the trusted-proxy list; the
// verifier only sees net/http types and could not resolve it correctly itself.
func WithClientIP(ctx context.Context, ip string) context.Context {
	return context.WithValue(ctx, clientIPKey{}, ip)
}

// ClientIPFromContext returns the IP set by WithClientIP, or "".
func ClientIPFromContext(ctx context.Context) string {
	ip, _ := ctx.Value(clientIPKey{}).(string)
	return ip
}
