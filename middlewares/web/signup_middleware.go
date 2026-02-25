package webMiddleware

import (
	"net/http"
	"sharpnr-engine/cookieManager"
	jwt_manager "sharpnr-engine/pkg/jwtmanager"

	"github.com/gin-gonic/gin"
)

const (
	SignupUserIDContextKey = "signup_user_id"
)

func SignupGuardMiddleware(
	jwtMgr jwt_manager.Manager,
	cookies cookieManager.Manager,
) gin.HandlerFunc {
	return func(ctx *gin.Context) {

		// 2️⃣ Read signup cookie
		signupToken, err := cookies.ReadSignup(ctx)
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

		ctx.Set(SignupUserIDContextKey, claims.UserID)

		ctx.Next()
	}
}
