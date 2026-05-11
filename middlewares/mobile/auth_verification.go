package mobile_middleware

import (
	jwt_manager "sharpnr-engine/pkg/jwtmanager"

	"github.com/gin-gonic/gin"
)

func MobileMiddleware(
	jwtMgr jwt_manager.Manager,
) gin.HandlerFunc {
	return func(ctx *gin.Context) {

		ctx.Next()
	}
}
