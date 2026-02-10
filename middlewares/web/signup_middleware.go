package webMiddleware

import (
	"net/http"
	jwt_manager "sharpnr-engine/pkg/jwtmanager"

	"github.com/gin-gonic/gin"
)

const (
	SignupCookieName       = "_signup"
	SignupUserIDContextKey = "signup_user_id"
)

func SignupGuardMiddleware(jwtMgr jwt_manager.Manager) gin.HandlerFunc {
	return func(ctx *gin.Context) {

		// If user is already authenticated, onboarding is irrelevant
		if _, err := ctx.Cookie("__Host-Auth"); err == nil {
			ctx.AbortWithStatusJSON(http.StatusForbidden, gin.H{
				"message": "Already authenticated",
			})
			return
		}

		// Check signup cookie
		signupToken, err := ctx.Cookie(SignupCookieName)
		if err != nil || signupToken == "" {
			ctx.AbortWithStatusJSON(http.StatusForbidden, gin.H{
				"message": "Signup session expired or invalid",
			})
			return
		}

		claims, err := jwtMgr.VerifySignup(signupToken)
		if err != nil {
			ctx.AbortWithStatusJSON(http.StatusForbidden, gin.H{
				"message": "Invalid or expired signup token",
			})
			return
		}

		ctx.Set(SignupCookieName, claims.UserID)

		ctx.Next()
	}
}
