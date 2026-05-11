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
	UserIDContextKey  = "user_id"
	OrgUUIDContextKey = "OrgUUID"
)

func WebAuthMiddleware(
	jwtMgr jwt_manager.Manager,
	cookies cookieManager.Manager,
) gin.HandlerFunc {
	return func(ctx *gin.Context) {

		refreshToken, err := cookies.Read(ctx, cookieManager.Auth)
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

		// 3️⃣ Roll CSRF — or mint a new one if it was deleted
		csrfToken, err := cookies.Read(ctx, cookieManager.CSRF)
		if err != nil || csrfToken == "" {
			csrfToken = uuid.New().String()
			slog.Info("CSRF cookie missing, issuing new token",
				"path", ctx.Request.URL.Path,
				"user_id", claims.UserID,
			)
		}

		if err := cookies.Set(ctx, cookieManager.CSRF, csrfToken); err != nil {
			slog.Warn("failed to set CSRF cookie",
				"error", err.Error(),
				"path", ctx.Request.URL.Path,
			)
		}

		// 3️⃣ Populate legacy context keys
		ctx.Set(UserIDContextKey, claims.UserID)
		ctx.Set(OrgUUIDContextKey, orgUUID)

		ctx.Next()
	}
}
