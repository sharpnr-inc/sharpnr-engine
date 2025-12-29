package mobile_auth_controller

import (
	"net/http"
	"sharpnr-engine/core/interfaces"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/redis/go-redis/v9"
)

type MobileAuthController struct {
	service interfaces.AuthService
	redis   *redis.Client
}

func NewMobileAuthController(service interfaces.AuthService) *MobileAuthController {
	return &MobileAuthController{service: service}
}

func (s *MobileAuthController) Login(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"statusCode": 200,
		"message":    "Mobile login dummy success",
		"error": gin.H{
			"isError": false,
			"message": "",
		},
		"data": gin.H{
			"access_token":  "dummy-access-token",
			"refresh_token": "dummy-refresh-token",
			"expires_in":    900,
			"user": gin.H{
				"id":    "user-123",
				"name":  "Test User",
				"email": "test@sharpnr.dev",
			},
		},
		"timestamp": time.Now(),
	})
}
