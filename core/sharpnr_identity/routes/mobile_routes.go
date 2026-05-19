package sharpnrIdentity_routes

import (
	"sharpnr-engine/app"
	sharpnr_identity_auth_controllers "sharpnr-engine/core/sharpnr_identity/controllers"
	mobile_middleware "sharpnr-engine/middlewares/mobile"
	"sharpnr-engine/pkg/cookieManager"
	jwt_manager "sharpnr-engine/pkg/jwtmanager"

	"github.com/gin-gonic/gin"
)

// Unprotected routes for the login
// for the mobile
func RegisterSharpnrIdentityMobileRoutes(
	mobileV1 *gin.RouterGroup,
	jwtMgr jwt_manager.Manager,
	modules app.ModulesProvider,
	cookie cookieManager.Manager,

) {
	sharpIdService := modules.SharpAuthService()
	sharpIdController := sharpnr_identity_auth_controllers.NewSharpnrIdentityAuthController(cookie, jwtMgr, sharpIdService)
	idMobileRoutes := mobileV1.Group("/identity")
	{
		idMobileRoutes.POST("/login", sharpIdController.MobileLogin)
		// Refresh endpoint — only needs the refresh token, NOT the access token
		// so it sits outside MobileMiddleware
		idMobileRoutes.POST("/auth/refresh", mobile_middleware.RefreshHandler(jwtMgr))

	}

}

// Procted Routes for the sharpnr Identity service
// These are accessible only after middleware validation
// both Refresh Token and Access Token
func RegisterIdentityProtectedMobileRoutes(
	mobileV1 *gin.RouterGroup,
	modules app.ModulesProvider,
	cookie cookieManager.Manager,
	jwtMgr jwt_manager.Manager,
) {
	sharpIdService := modules.SharpAuthService()
	sharpIdController := sharpnr_identity_auth_controllers.NewSharpnrIdentityAuthController(cookie, jwtMgr, sharpIdService)

	idMobileProtectesRoutes := mobileV1.Group("/identity")
	{
		idMobileProtectesRoutes.GET("/myprofile", sharpIdController.GetMyProfile)
		idMobileProtectesRoutes.POST("/update")
	}

}
