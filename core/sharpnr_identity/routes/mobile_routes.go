package sharpnrIdentity_routes

import (
	"github.com/gin-gonic/gin"
)

// Unprotected routes for the login
// for the mobile
func RegisterSharpnrIdentityMobileRoutes(
	mobileV1 *gin.RouterGroup,
) {

	idMobileRoutes := mobileV1.Group("/identity")
	{
		idMobileRoutes.POST("/login")

	}

}

// Procted Routes for the sharpnr Identity service
// These are accessible only after middleware validation
// both Refresh Token and Access Token
func RegisterIdentityProtectedMobileRoutes(
	mobileV1 *gin.RouterGroup,
) {
	idMobileProtectesRoutes := mobileV1.Group("/identity")

	idMobileProtectesRoutes.POST("/refresh")
	idMobileProtectesRoutes.GET("/myprofile")
	idMobileProtectesRoutes.POST("/update")
}
