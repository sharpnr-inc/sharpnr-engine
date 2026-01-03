package webMiddleware

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

const signupCookieName = "__Host-signup"

func SignupGuardMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {

		// If user is already authenticated, onboarding is irrelevant
		if _, err := ctx.Cookie("__Host-auth"); err == nil {
			ctx.AbortWithStatusJSON(http.StatusForbidden, gin.H{
				"message": "Already authenticated",
			})
			return
		}

		// Check signup cookie
		if _, err := ctx.Cookie(signupCookieName); err != nil {
			ctx.AbortWithStatusJSON(http.StatusForbidden, gin.H{
				"message": "Signup session expired or invalid",
			})
			return
		}

		ctx.Next()
	}
}
