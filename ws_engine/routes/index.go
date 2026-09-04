// Package ws_engine_routes mounts the app's single WebSocket endpoint.
//
// It is the glue between gin and the transport engine: the engine itself takes
// net/http types and knows nothing about the framework, the session cookie, or
// how this app names its context keys.
package ws_engine_routes

import (
	"errors"
	"fmt"
	"net/http"

	"sharpnr-engine/app"
	webMiddleware "sharpnr-engine/middlewares/web"
	"sharpnr-engine/shared/utils"
	ws_engine "sharpnr-engine/ws_engine"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

// RegisterWebSocketRoutes mounts /ws on the authenticated web group.
//
// One socket for the whole app, multiplexed by topic — a client does not open a
// connection per module. Which topics it may actually subscribe to is decided
// by the authorizer each module registers on the engine, so mounting this route
// grants nothing on its own.
//
// Web only, deliberately. The handshake is authenticated by the session cookie
// the browser attaches automatically; a mobile client using bearer tokens
// cannot set headers on a browser-style upgrade and needs its own scheme when
// that time comes.
func RegisterWebSocketRoutes(v1 *gin.RouterGroup, modules app.ModulesProvider) {
	engine := modules.WsEngine()

	v1.GET("/ws", func(ctx *gin.Context) {
		orgID, emplID, ok := resolveIdentity(ctx)
		if !ok {
			return
		}

		// Serve blocks for the life of the connection, which is what keeps gin
		// from finalizing a response on a socket that has already been
		// upgraded.
		engine.Serve(ctx.Writer, ctx.Request, ws_engine.Identity{
			UserID: emplID,
			OrgID:  orgID,
		})
	})
}

// resolveIdentity reads the identity the auth middleware stamped on the
// request. Same shape as the controllers' resolveRequestContexts — repeated
// rather than shared because that one belongs to a controller struct.
func resolveIdentity(ctx *gin.Context) (orgID, emplID uuid.UUID, ok bool) {
	emplIDRaw, exists := ctx.Get(webMiddleware.UserIDContextKey)
	if !exists {
		utils.Failure(ctx, http.StatusUnauthorized, errors.New("user not authenticated"))
		return uuid.Nil, uuid.Nil, false
	}

	emplID, err := uuid.Parse(fmt.Sprint(emplIDRaw))
	if err != nil {
		utils.Failure(ctx, http.StatusInternalServerError, errors.New("invalid employee uuid format"))
		return uuid.Nil, uuid.Nil, false
	}

	orgIDRaw, exists := ctx.Get(webMiddleware.OrgUUIDContextKey)
	if !exists {
		utils.Failure(ctx, http.StatusUnauthorized, errors.New("organization not found in context"))
		return uuid.Nil, uuid.Nil, false
	}

	orgID, err = uuid.Parse(fmt.Sprint(orgIDRaw))
	if err != nil {
		utils.Failure(ctx, http.StatusInternalServerError, errors.New("invalid org uuid format"))
		return uuid.Nil, uuid.Nil, false
	}

	return orgID, emplID, true
}
