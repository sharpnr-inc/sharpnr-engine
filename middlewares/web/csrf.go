package webMiddleware

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

const (
	csrfCookieName = "XSRF-TOKEN"
	csrfHeaderName = "X-CSRF-Token"
)

func CSRFMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {

		// 🔐 Ensure user is authenticated first
		if _, exists := ctx.Get(SessionIDContextKey); !exists {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"message": "Authentication required",
			})
			return
		}

		// Only protect unsafe methods
		switch ctx.Request.Method {
		case http.MethodGet, http.MethodHead, http.MethodOptions:
			ctx.Next()
			return
		}

		csrfCookie, err := ctx.Cookie(csrfCookieName)
		if err != nil || csrfCookie == "" {
			ctx.AbortWithStatusJSON(http.StatusForbidden, gin.H{
				"message": "CSRF cookie missing",
			})
			return
		}

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
