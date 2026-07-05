package mobile_middleware

import (
	"errors"
	"log/slog"
	"net/http"
	"strings"
	"time"

	jwt_manager "sharpnr-engine/pkg/jwtmanager"
	sentinelContexts "sharpnr-engine/sentinel_contexts"

	"github.com/gin-gonic/gin"
)

const (
	UserIDContextKey  = "user_id"
	OrgUUIDContextKey = "OrgUUID"
)

// MobileMiddleware authenticates requests from the mobile app.
//
// Expects a Bearer token in the Authorization header:
//
//	Authorization: Bearer <access_token>
//
// On success: sets user_id and org_id in the Gin context and calls Next().
//
// On expired token: returns 401 with code "token_expired".
// The mobile client must catch this specific code and call the
// /refresh endpoint to obtain a new access token — do NOT retry
// with the same token.
//
// On any other failure: returns 401 with code "unauthorized".
func MobileMiddleware(jwtMgr jwt_manager.Manager) gin.HandlerFunc {
	return func(ctx *gin.Context) {

		// 1️⃣ Extract Bearer token from Authorization header
		accessToken, ok := extractBearer(ctx)
		if !ok {
			slog.Warn("mobile auth failed",
				"reason", "missing_or_malformed_header",
				"path", ctx.Request.URL.Path,
				"ip", ctx.ClientIP(),
			)
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"code":    "unauthorized",
				"message": "Authorization header missing or malformed",
			})
			return
		}

		// 2️⃣ Verify the access token
		claims, err := jwtMgr.VerifyAccess(accessToken)
		if err != nil {
			if errors.Is(err, jwt_manager.ErrExpiredToken) {
				// Return a distinct code so the mobile client knows to
				// trigger the refresh flow, not show a login screen.
				slog.Info("mobile access token expired",
					"path", ctx.Request.URL.Path,
					"ip", ctx.ClientIP(),
				)
				ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
					"code":    "token_expired",
					"message": "Access token expired, please refresh",
				})
				return
			}

			slog.Warn("mobile auth failed",
				"reason", "invalid_token",
				"error", err.Error(),
				"path", ctx.Request.URL.Path,
				"ip", ctx.ClientIP(),
			)
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"code":    "unauthorized",
				"message": "Invalid token",
			})
			return
		}

		// 3️⃣ Populate context — available to all downstream handlers
		ctx.Set(UserIDContextKey, claims.UserID)
		ctx.Set(OrgUUIDContextKey, claims.OrgID)

		// 4️⃣ Stamp the Sentinel authorization context onto the request so the
		// engine can expose it to policies (context.channel, context.ip, ...).
		ctx.Request = ctx.Request.WithContext(sentinelContexts.Inject(
			ctx.Request.Context(),
			sentinelContexts.AuthzContext{
				Request: sentinelContexts.RequestContext{
					Channel:     sentinelContexts.ChannelMobile,
					IPAddress:   ctx.ClientIP(),
					RequestedAt: time.Now(),
				},
			},
		))

		ctx.Next()
	}
}

// extractBearer pulls the token out of "Authorization: Bearer <token>".
// Returns the raw token string and true on success.
func extractBearer(ctx *gin.Context) (string, bool) {
	header := ctx.GetHeader("Authorization")
	if header == "" {
		return "", false
	}

	parts := strings.SplitN(header, " ", 2)
	if len(parts) != 2 || !strings.EqualFold(parts[0], "bearer") {
		return "", false
	}

	token := strings.TrimSpace(parts[1])
	if token == "" {
		return "", false
	}

	return token, true
}
