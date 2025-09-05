package v1

import (
	orgcontrollers "sharpnr-engine/core/controllers/organisation_controllers"
	organizationrepositories "sharpnr-engine/core/repositories/organization_repo"
	orgservices "sharpnr-engine/core/services/org_service"
	"sharpnr-engine/shared/middlewares"

	"github.com/gin-gonic/gin"
	"github.com/redis/go-redis/v9"
	"gorm.io/gorm"
)

// These Are the Auth routes for the authentication registration
func AuthRoutes(v1 *gin.RouterGroup, db *gorm.DB, redis *redis.Client) {

	newOrgRepo := organizationrepositories.InitializeNewOrgRegistration(db, redis)
	newOrgAuthService := orgservices.OrgRegService(newOrgRepo)
	newOrgRegController := orgcontrollers.RegisterOrgController(newOrgAuthService)

	//Register
	v1.POST("/register", newOrgRegController.InitOrgReg)
	//Need Middleware
	v1.POST("/otp-verify", middlewares.OrgRegOtpVerificationMiddleware(), newOrgRegController.VerifyOtp)
	v1.POST("/complete-register")
	v1.POST("/login", newOrgRegController.Login)

	v1.POST("/reset-password")

}
