package authenticationEngineRoutes

import (
	"sharpnr-engine/app"
	authenticationEngineControllers "sharpnr-engine/authentication/controllers"
	"sharpnr-engine/pkg/cookieManager"
	jwt_manager "sharpnr-engine/pkg/jwtmanager"

	"github.com/gin-gonic/gin"
)

func RegisterAuthenticationWebRoutes(
	routeGroup *gin.RouterGroup,
	cookie cookieManager.Manager,
	jwtMgr jwt_manager.Manager,
	modules app.ModulesProvider,
) {

	authService := modules.AuthenticationService()
	authController := authenticationEngineControllers.NewAuthenticationController(
		cookie,
		jwtMgr,
		authService,
	)

	idWebRoutes := routeGroup.Group("/identity/web")
	{
		idWebRoutes.POST("/login", authController.WebLogin)

	}
}

func RegisterAuthenticationProtectedWebRoutes(
	routeGroup *gin.RouterGroup,
	cookie cookieManager.Manager,
	jwtMgr jwt_manager.Manager,
	modules app.ModulesProvider,
) {
	authService := modules.AuthenticationService()
	authController := authenticationEngineControllers.NewAuthenticationController(cookie, jwtMgr, authService)

	idProtectedWebRoutes := routeGroup.Group("/identity/web")
	{
		idProtectedWebRoutes.GET("/profile", authController.GetMyProfile)
		idProtectedWebRoutes.GET("/verify-auth", authController.VerifyAuthentication)
		idProtectedWebRoutes.GET("/user", authController.GetUser)
		idProtectedWebRoutes.GET("/logout", authController.LogoutUser)
		idProtectedWebRoutes.POST("/update", authController.UpdateMyProfile)
	}
}
