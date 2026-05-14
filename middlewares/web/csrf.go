package webMiddleware

import (
	"log/slog"
	"net/http"
	"sharpnr-engine/cookieManager"

	"github.com/gin-gonic/gin"
)

const (
	csrfHeaderName = "X-CSRF-Token"
)

func CSRFMiddleware(cookies cookieManager.Manager) gin.HandlerFunc {
	return func(ctx *gin.Context) {

		// Only protect unsafe methods
		switch ctx.Request.Method {
		case http.MethodGet, http.MethodHead, http.MethodOptions:
			ctx.Next()
			return
		}

		// 1️⃣ Read CSRF cookie via manager
		csrfCookie, err := cookies.Read(ctx, cookieManager.CSRF)
		if err != nil || csrfCookie == "" {
			slog.Error("CSRF validation failed",
				"reason", "missing_cookie",
				"method", ctx.Request.Method,
				"path", ctx.Request.URL.Path,
				"ip", ctx.ClientIP(),
			)
			ctx.AbortWithStatusJSON(http.StatusForbidden, gin.H{
				"message": "CSRF cookie missing",
			})
			return
		}

		// 2️⃣ Read header token
		csrfHeader := ctx.GetHeader(csrfHeaderName)
		if csrfHeader == "" {
			slog.Error("CSRF validation failed",
				"reason", "missing_header",
				"method", ctx.Request.Method,
				"path", ctx.Request.URL.Path,
				"ip", ctx.ClientIP(),
			)
			ctx.AbortWithStatusJSON(http.StatusForbidden, gin.H{
				"message": "CSRF header missing",
			})
			return
		}

		if csrfHeader != csrfCookie {
			slog.Error("CSRF validation failed",
				"reason", "mismatch",
				"method", ctx.Request.Method,
				"path", ctx.Request.URL.Path,
				"ip", ctx.ClientIP(),
			)
			ctx.AbortWithStatusJSON(http.StatusForbidden, gin.H{
				"message": "CSRF validation failed",
			})
			return
		}

		ctx.Next()
	}
}
