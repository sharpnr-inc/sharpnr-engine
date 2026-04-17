package sharpnrIdentity_routes

import (
	"sharpnr-engine/app"
	"sharpnr-engine/cookieManager"
	sharpnr_identity_auth_controllers "sharpnr-engine/core/sharpnr_identity/controllers"
	jwt_manager "sharpnr-engine/pkg/jwtmanager"

	"github.com/gin-gonic/gin"
)

func RegisterSharpnrIdentityWebRoutes(
	routeGroup *gin.RouterGroup,
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

	idWebRoutes := routeGroup.Group("/identity/web")
	{
		idWebRoutes.POST("/login", sharpIdController.WebLogin)
		idWebRoutes.POST("/refresh")
	}
}

func RegisterSharpnrIdentityProtectedWebRoutes(
	routeGroup *gin.RouterGroup,
	cookie cookieManager.Manager,
	jwtMgr jwt_manager.Manager,
	modules app.ModulesProvider,
) {
	sharpIdService := modules.SharpAuthService()
	sharpIdController := sharpnr_identity_auth_controllers.NewSharpnrIdentityAuthController(cookie, jwtMgr, sharpIdService)

	idProtectedWebRoutes := routeGroup.Group("/identity/web")
	{
		idProtectedWebRoutes.GET("/profile", sharpIdController.GetMyProfile)
		idProtectedWebRoutes.GET("/verify-auth", sharpIdController.VerifyAuthentication)
		idProtectedWebRoutes.GET("/logout", sharpIdController.LogoutUser)
		idProtectedWebRoutes.POST("/update", sharpIdController.UpdateMyProfile)
	}
}
