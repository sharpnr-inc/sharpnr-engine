package webMiddleware

import (
	"net/http"
	"sharpnr-engine/cookieManager"

	"github.com/gin-gonic/gin"
)

const (
	csrfCookieName = "_dev_csrf"
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
		csrfCookie, err := cookies.ReadCSRF(ctx)
		if err != nil || csrfCookie == "" {
			ctx.AbortWithStatusJSON(http.StatusForbidden, gin.H{
				"message": "CSRF cookie missing",
			})
			return
		}

		// 2️⃣ Read header token
		csrfHeader := ctx.GetHeader(csrfHeaderName)
		if csrfHeader == "" || csrfHeader != csrfCookie {
			ctx.AbortWithStatusJSON(http.StatusForbidden, gin.H{
				"message": "CSRF validation failed",
			})
			return
		}

		ctx.Next()
	}
}
