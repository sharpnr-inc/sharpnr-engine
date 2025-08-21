package v1

import (
	orgcontrollers "sharpnr-engine/controllers/organisationControllers"
	repositories "sharpnr-engine/core/repositories/organization"
	services "sharpnr-engine/core/services/auth"
	"sharpnr-engine/shared/middlewares"

	"github.com/gin-gonic/gin"
	"github.com/redis/go-redis/v9"
	"gorm.io/gorm"
)

func CoreV1Routes(v1 *gin.RouterGroup, db *gorm.DB, redis *redis.Client) {

	newOrgRepo := repositories.InitializeNewOrgRegistration(db, redis)
	newOrgAuthService := services.OrgRegService(newOrgRepo)
	newOrgRegController := orgcontrollers.RegisterOrgController(newOrgAuthService)

	v1.POST("/org/register", newOrgRegController.InitOrgReg)
	v1.POST("/org/verify-email-otp", middlewares.OrgRegOtpVerificationMiddleware(), newOrgRegController.VerifyOtp)
	v1.GET("/org/reset-password", newOrgRegController.ResetOrgPassword)
	v1.POST("/org/update", newOrgRegController.UpdateOrgRegistration)
}
