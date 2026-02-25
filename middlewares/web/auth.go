package webMiddleware

import (
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
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"message": "Authentication required",
			})
			return
		}

		// 2️⃣ Verify refresh token
		claims, err := jwtMgr.VerifyWebCookie(refreshToken)
		if err != nil {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"message": "Invalid or expired session",
			})
			return
		}

		orgUUID, err := uuid.Parse(claims.OrgID)
		if err != nil {
			// token is already invalid if this fails
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"error": "invalid org id in token",
			})
			return
		}

		// 3️⃣ Populate legacy context keys
		ctx.Set(UserIDContextKey, claims.UserID)
		ctx.Set(OrgUUIDContextKey, orgUUID) // 👈 legacy-compatible

		ctx.Next()
	}
}
