package org_signup

import (
	"sharpnr-engine/app"
	org_signup_controllers "sharpnr-engine/core/org_signup/controllers"
	webMiddleware "sharpnr-engine/middlewares/web"
	"sharpnr-engine/pkg/cookieManager"
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
		cookies,
		jwtmgr,
		orgSignUpsvc,
	)

	publicSignup := v1.Group("/signup")
	{ // this is for starting the SignUp process. after this user will recieve an OTP in their email, and we will set the cokkie with the token for the next step of the signup process
		publicSignup.POST("/initiate", controller.InitiateNewOrgRegistration)
	}

	protectedSignup := v1.Group("/signup")
	protectedSignup.Use(webMiddleware.SignupGuardMiddleware(jwtmgr, cookies))
	{
		protectedSignup.GET("/session/verify", controller.VerifySlug)
		protectedSignup.POST("/verify-otp", controller.VerifyOtpForSignup)
		// Not set up yet.
		protectedSignup.GET("/resend/otp", controller.ResendOtp)
	}
}
