package sharpnr_identity_auth

import (
	"sharpnr-engine/cookieManager"
	"sharpnr-engine/sere"

	"github.com/gin-gonic/gin"
	"github.com/redis/go-redis/v9"
	"gorm.io/gorm"
)

func SharpnrIdentityAuthWebRoutes(
	webV1 *gin.RouterGroup,
	db *gorm.DB,
	redis *redis.Client,
	sere *sere.SERE,
	cookie cookieManager.Manager,
) {

	idWebRoutes := webV1.Group("/identity/web")
	{
		idWebRoutes.POST("/login")
		idWebRoutes.POST("/refresh")
		idWebRoutes.GET("/myprofile")
		idWebRoutes.POST("/renderer")
		idWebRoutes.POST("/update")
	}

}

func SharpnrIdentityAuthMobileRoutes(
	mobileV1 *gin.RouterGroup,
	db *gorm.DB,
	redis *redis.Client,
	sere *sere.SERE) {

	idMobileRoutes := mobileV1.Group("/identity/mobile")
	{
		idMobileRoutes.POST("/login")
		idMobileRoutes.POST("/refresh")
		idMobileRoutes.GET("/myprofile")
		idMobileRoutes.POST("/renderer")
		idMobileRoutes.POST("/update")
	}

}
