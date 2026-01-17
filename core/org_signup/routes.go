package org_signup

import (
	"sharpnr-engine/app"
	"sharpnr-engine/cookieManager"
	org_signup_controllers "sharpnr-engine/core/org_signup/controllers"
	webMiddleware "sharpnr-engine/middlewares/web"
	jwt_manager "sharpnr-engine/pkg/jwtmanager"

	"github.com/gin-gonic/gin"
)

func OrgSignUpRoute(
	v1 *gin.RouterGroup,
	cookies cookieManager.Manager,
	jwtmgr jwt_manager.Manager,
	modules app.ModulesProvider,
) {
	orgSignUpsvc := modules.OrgSignUpService()

	controller := org_signup_controllers.NewOrgSignController(
		orgSignUpsvc,
		cookies,
		jwtmgr,
	)

	publicSignup := v1.Group("/signup")
	{
		publicSignup.POST("/initiate", controller.InitiateNewOrgRegistration)
	}

	protectedSignup := v1.Group("/signup")
	protectedSignup.Use(webMiddleware.SignupGuardMiddleware(jwtmgr))
	{
		protectedSignup.POST("/verify-otp", controller.VerifyOtpForSignup)
		protectedSignup.GET("/verify-token", controller.GetSignUpDetails)

		protectedSignup.GET("/resend-otp", controller.ResendOtp)
	}
}
