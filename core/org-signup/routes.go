package org_signup

import (
	"sharpnr-engine/app"
	"sharpnr-engine/cookieManager"
	dbtx "sharpnr-engine/core/dbTx"
	org_signup_controllers "sharpnr-engine/core/org-signup/controllers"
	webMiddleware "sharpnr-engine/middlewares/web"
	jwt_manager "sharpnr-engine/pkg/jwtmanager"
	"sharpnr-engine/sere"

	"github.com/gin-gonic/gin"
	"github.com/redis/go-redis/v9"
	"gorm.io/gorm"
)

func OrgSignUpRoute(
	v1 *gin.RouterGroup,
	db *gorm.DB,
	txManager dbtx.TxManager,
	redis *redis.Client,
	sere *sere.SERE,
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

		protectedSignup.POST("/onBoarding", controller.OnBoardOrganization)
	}
}
