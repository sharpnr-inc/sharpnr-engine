// Package mcp_engine_routes mounts the MCP endpoint on the gin router.
//
// The same split as ws_engine: the engine takes net/http types, this package
// knows about gin, the modules provider and how the app resolves a client IP.
package mcp_engine_routes

import (
	"sharpnr-engine/app"
	mcp_engine "sharpnr-engine/mcp_engine"
	"sharpnr-engine/meta"
	jwt_manager "sharpnr-engine/pkg/jwtmanager"

	"github.com/gin-gonic/gin"
)

// RegisterMCPRoutes mounts the MCP server at /mcp.
//
// It sits on the root router, outside /api/v1, because it is not part of the
// web or mobile API: MCP clients speak JSON-RPC over Streamable HTTP and
// authenticate with a bearer token, not the session cookie and CSRF pair.
// Any() because the transport uses POST, and in stateful mode GET and DELETE,
// on the one path.
func RegisterMCPRoutes(router *gin.Engine, jwtMgr jwt_manager.Manager, modules app.ModulesProvider) {
	handler := mcp_engine.NewHandler(
		mcp_engine.Config{
			Version:  meta.Version,
			Verifier: mcp_engine.AccessTokenVerifier(jwtMgr),
		},
		mcp_engine.Deps{
			Employees: modules.SharpnrOrchestrator(),
		},
	)

	router.Any("/mcp", func(ctx *gin.Context) {
		// gin resolves the client IP against the trusted-proxy list; hand it
		// to the engine, which only sees the *http.Request.
		ctx.Request = ctx.Request.WithContext(
			mcp_engine.WithClientIP(ctx.Request.Context(), ctx.ClientIP()),
		)
		handler.ServeHTTP(ctx.Writer, ctx.Request)
	})
}
