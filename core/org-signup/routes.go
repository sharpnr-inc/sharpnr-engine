package org_signup

import (
	"sharpnr-engine/cookieManager"
	org_signup_controllers "sharpnr-engine/core/org-signup/controllers"
	org_signup_repositories "sharpnr-engine/core/org-signup/repositories"
	org_signup_services "sharpnr-engine/core/org-signup/services"
	"sharpnr-engine/sere"

	"github.com/gin-gonic/gin"
	"github.com/redis/go-redis/v9"
	"gorm.io/gorm"
)

func OrgSignUpRoute(
	v1 *gin.RouterGroup,
	db *gorm.DB,
	redis *redis.Client,
	sere *sere.SERE,
	cookies cookieManager.Manager,
) {

	tempSignUpRepo := org_signup_repositories.NewTempOrgSignUpRepository(redis)
	mainSignUpRepo := org_signup_repositories.NewOrgMainSignUpRepository(db)

	signUpService := org_signup_services.NewOrgSignUpService(sere, tempSignUpRepo, mainSignUpRepo)
	controller := org_signup_controllers.NewOrgSignController(signUpService, cookies)

	sgnRoutes := v1.Group("/signup")

	{
		sgnRoutes.POST("/initiate", controller.InitiateNewOrgRegistration)
		sgnRoutes.GET("/resend-otp")
		sgnRoutes.POST("/verify-otp")
		sgnRoutes.GET("/verify-token")
		sgnRoutes.POST("/onBoarding")
	}
}
