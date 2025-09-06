package v1

import (
	signupcontrollers "sharpnr-engine/core/controllers/signup_controllers"
	organizationrepositories "sharpnr-engine/core/repositories/organization_repo"
	signupservice "sharpnr-engine/core/services/signup_service"
	"sharpnr-engine/shared/middlewares"

	"github.com/gin-gonic/gin"
	"github.com/redis/go-redis/v9"
	"gorm.io/gorm"
)

// These Are the Auth routes for the authentication registration
func AuthRoutes(v1 *gin.RouterGroup, db *gorm.DB, redis *redis.Client) {

	newOrgRepo := organizationrepositories.NewOrganizationRepo(db, redis)
	newSignupService := signupservice.NewSignupService(newOrgRepo)
	newSignupController := signupcontrollers.NewSignUpController(newSignupService)

	//Register
	v1.POST("/register", newSignupController.RegisterOrg)
	//Need Middleware
	v1.POST("/otp-verify", middlewares.OrgRegOtpVerificationMiddleware(), newSignupController.VerifyRegistrationOTP)
	v1.POST("/complete-register")
	// v1.POST("/login", newOrgRegController.Login)

	v1.POST("/reset-password")

}
