package v1

import (
	orgcontrollers "sharpnr-engine/core/controllers/organisationControllers"
	organizationrepositories "sharpnr-engine/core/repositories/organization_repo"
	orgservices "sharpnr-engine/core/services/org_service"
	"sharpnr-engine/shared/middlewares"

	"github.com/gin-gonic/gin"
	"github.com/redis/go-redis/v9"
	"gorm.io/gorm"
)

func CoreV1Routes(v1 *gin.RouterGroup, db *gorm.DB, redis *redis.Client) {

	newOrgRepo := organizationrepositories.InitializeNewOrgRegistration(db, redis)
	newOrgAuthService := orgservices.OrgRegService(newOrgRepo)
	newOrgRegController := orgcontrollers.RegisterOrgController(newOrgAuthService)

	v1.POST("/org/register", newOrgRegController.InitOrgReg)
	v1.POST("/org/verify-email-otp", middlewares.OrgRegOtpVerificationMiddleware(), newOrgRegController.VerifyOtp)
	v1.GET("/org/reset-password", newOrgRegController.ResetOrgPassword)
	v1.POST("/org/update", newOrgRegController.UpdateOrgRegistration)
	v1.POST("/org/dashboard", newOrgRegController.GetOrgRootDashBoard)
	v1.POST("/org/login", newOrgRegController.Login)
}
