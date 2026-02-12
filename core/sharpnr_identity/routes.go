package sharpnr_identity_auth

import (
	"sharpnr-engine/app"
	"sharpnr-engine/cookieManager"
	sharpnr_identity_auth_controllers "sharpnr-engine/core/sharpnr_identity/controllers"
	jwt_manager "sharpnr-engine/pkg/jwtmanager"

	"github.com/gin-gonic/gin"
	"github.com/redis/go-redis/v9"
	"gorm.io/gorm"
)

func SharpnrIdentityAuthWebRoutes(
	webV1 *gin.RouterGroup,
	cookie cookieManager.Manager,
	jwtMgr jwt_manager.Manager,
	modules app.ModulesProvider,
) {

	sharpIdService := modules.SharpAuthService()
	sharpIdController := sharpnr_identity_auth_controllers.NewSharpnrIdentityAuthController(
		cookie,
		jwtMgr,
		sharpIdService,
	)

	idWebRoutes := webV1.Group("/identity/web")
	{
		idWebRoutes.POST("/login", sharpIdController.WebLogin)
		idWebRoutes.POST("/refresh")
		idWebRoutes.GET("/myprofile")
		idWebRoutes.POST("/renderer")
		idWebRoutes.POST("/update")
	}

}

func SharpnrIdentityAuthMobileRoutes(
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
