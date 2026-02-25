package sharpnrIdentity_routes

import (
	"github.com/gin-gonic/gin"
	"github.com/redis/go-redis/v9"
	"gorm.io/gorm"
)

func RegisterSharpnrIdentityMobileRoutes(
	mobileV1 *gin.RouterGroup,
	db *gorm.DB,
	redis *redis.Client) {

	idMobileRoutes := mobileV1.Group("/identity/mobile")
	{
		idMobileRoutes.POST("/login")
		idMobileRoutes.POST("/refresh")
		idMobileRoutes.GET("/myprofile")
		idMobileRoutes.POST("/renderer")
		idMobileRoutes.POST("/update")
	}

}
