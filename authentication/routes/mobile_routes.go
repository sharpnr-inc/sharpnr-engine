package authenticationEngineRoutes

import (
	"sharpnr-engine/app"
	authenticationEngineControllers "sharpnr-engine/authentication/controllers"
	mobile_middleware "sharpnr-engine/middlewares/mobile"
	"sharpnr-engine/pkg/cookieManager"
	jwt_manager "sharpnr-engine/pkg/jwtmanager"

	"github.com/gin-gonic/gin"
)

// Unprotected routes for the login
// for the mobile
func RegisterAuthenticationMobileRoutes(
	mobileV1 *gin.RouterGroup,
	jwtMgr jwt_manager.Manager,
	modules app.ModulesProvider,
	cookie cookieManager.Manager,

) {
	authService := modules.AuthenticationService()
	authController := authenticationEngineControllers.NewAuthenticationController(cookie, jwtMgr, authService)
	idMobileRoutes := mobileV1.Group("/identity")
	{
		idMobileRoutes.POST("/login", authController.MobileLogin)
		// Refresh endpoint — only needs the refresh token, NOT the access token
		// so it sits outside MobileMiddleware
		idMobileRoutes.POST("/auth/refresh", mobile_middleware.RefreshHandler(jwtMgr))

	}

}

// Procted Routes for the sharpnr Identity service
// These are accessible only after middleware validation
// both Refresh Token and Access Token
func RegisterAuthenticationProtectedMobileRoutes(
	mobileV1 *gin.RouterGroup,
	modules app.ModulesProvider,
	cookie cookieManager.Manager,
	jwtMgr jwt_manager.Manager,
) {
	authService := modules.AuthenticationService()
	authController := authenticationEngineControllers.NewAuthenticationController(cookie, jwtMgr, authService)

	idMobileProtectesRoutes := mobileV1.Group("/identity")
	{
		idMobileProtectesRoutes.GET("/myprofile", authController.GetMyProfile)
	}

}
