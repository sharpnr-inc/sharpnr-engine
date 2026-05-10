package webMiddleware

import (
	"log/slog"
	"net/http"

	"sharpnr-engine/cookieManager"
	jwt_manager "sharpnr-engine/pkg/jwtmanager"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

const (
	UserIDContextKey    = "user_id"
	OrgUUIDContextKey   = "OrgUUID"
	SessionIDContextKey = "session_id"
)

func WebAuthMiddleware(
	jwtMgr jwt_manager.Manager,
	cookies cookieManager.Manager,
) gin.HandlerFunc {
	return func(ctx *gin.Context) {

		refreshToken, err := cookies.ReadAuth(ctx)
		if err != nil {
			slog.Warn("authentication failed",
				"reason", "missing_auth_cookie",
				"path", ctx.Request.URL.Path,
				"ip", ctx.ClientIP(),
			)
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"message": "Authentication required",
			})
			return
		}

		// 2️⃣ Verify refresh token
		claims, err := jwtMgr.VerifyWebCookie(refreshToken)
		if err != nil {
			slog.Warn("authentication failed",
				"reason", "invalid_token",
				"error", err.Error(),
				"path", ctx.Request.URL.Path,
				"ip", ctx.ClientIP(),
			)
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"message": "Invalid or expired session",
			})
			return
		}

		orgUUID, err := uuid.Parse(claims.OrgID)
		if err != nil {
			slog.Error("token payload corruption",
				"reason", "invalid_org_uuid",
				"user_id", claims.UserID,
				"org_id_raw", claims.OrgID,
				"error", err.Error(),
			)
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"error": "invalid org id in token",
			})
			return
		}

		// 3️⃣ Populate legacy context keys
		ctx.Set(UserIDContextKey, claims.UserID)
		ctx.Set(OrgUUIDContextKey, orgUUID)

		ctx.Next()
	}
}
