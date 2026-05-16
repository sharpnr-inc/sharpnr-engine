package webMiddleware

import (
	"net/http"
	"sharpnr-engine/cookieManager"
	jwt_manager "sharpnr-engine/pkg/jwtmanager"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid" // Or whichever UUID package you use
)

const (
	SignupUserIDContextKey = "signup_user_id"
)

// GetSignupUserID retrieves the UUID from the context with proper type assertion
func GetSignupUserID(ctx *gin.Context) (uuid.UUID, bool) {
	val, exists := ctx.Get(SignupUserIDContextKey)
	if !exists {
		return uuid.Nil, false
	}

	id, ok := val.(uuid.UUID)
	return id, ok
}

func SignupGuardMiddleware(
	jwtMgr jwt_manager.Manager,
	cookies cookieManager.Manager,
) gin.HandlerFunc {
	return func(ctx *gin.Context) {

		// Read signup cookie
		signupToken, err := cookies.Read(ctx, cookieManager.Signup)
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

		// Ensure claims.UserID is actually a uuid.UUID type here.
		// If claims.UserID is a string, parse it first: uuid.Parse(claims.UserID)
		ctx.Set(SignupUserIDContextKey, claims.UserID)
		ctx.Next()
	}
}
