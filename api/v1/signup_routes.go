package v1

import (
	"sharpnr-engine/cookieManager"
	"sharpnr-engine/sere"

	"github.com/gin-gonic/gin"
	"github.com/redis/go-redis/v9"
	"gorm.io/gorm"
)

func SignUpRoutes(v1 *gin.RouterGroup, db *gorm.DB, redis *redis.Client, sere *sere.SERE, cookies cookieManager.Manager) {

	signup := v1.Group("/signup")

	signup.POST("/initiate")
	signup.POST("/verify-email")
	signup.POST("/On-boarding")
	signup.GET("/resend-otp")

}
