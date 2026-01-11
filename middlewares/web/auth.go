package webMiddleware

import (
	"net/http"

	jwt_manager "sharpnr-engine/pkg/jwtmanager"

	"github.com/gin-gonic/gin"
)

const (
	AuthCookie = "__Host-Auth"

	// Legacy-compatible context keys
	UserIDContextKey    = "user_id"
	OrgUUIDContextKey   = "OrgUUID"
	SessionIDContextKey = "session_id"
)

func WebAuthMiddleware(jwtMgr jwt_manager.Manager) gin.HandlerFunc {
	return func(ctx *gin.Context) {

		// 1️⃣ Read refresh token from cookie
		refreshToken, err := ctx.Cookie(AuthCookie)
		if err != nil {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"message": "Authentication required",
			})
			return
		}

		// 2️⃣ Verify refresh token
		claims, err := jwtMgr.VerifyRefresh(refreshToken)
		if err != nil {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"message": "Invalid or expired session",
			})
			return
		}

		// 3️⃣ Populate legacy context keys
		ctx.Set(UserIDContextKey, claims.UserID)
		ctx.Set(OrgUUIDContextKey, claims.OrgID) // 👈 legacy-compatible
		ctx.Set(SessionIDContextKey, claims.SessionID)

		// 4️⃣ Continue request
		ctx.Next()
	}
}
