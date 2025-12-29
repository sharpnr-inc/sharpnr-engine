package v1

import (
	"sharpnr-engine/core/controllers/mobile_auth_controller"
	organizationrepositories "sharpnr-engine/core/repositories/organization_repo"
	authservice "sharpnr-engine/core/services/auth_service"
	hr_employee_repo "sharpnr-engine/modules/hr/employees/repo"
	"sharpnr-engine/sere"

	"github.com/gin-gonic/gin"
	"github.com/redis/go-redis/v9"
	"gorm.io/gorm"
)

func MobileAuthRoutes(v1 *gin.RouterGroup, db *gorm.DB, redis *redis.Client, sere *sere.SERE) {

	newOrgRepo := organizationrepositories.NewOrganizationRepo(db, redis)
	newHrEmployeeRepo := hr_employee_repo.NewHrEmployeeRepository(db)

	authService := authservice.NewAuthService(newOrgRepo, newHrEmployeeRepo)
	ctlr := mobile_auth_controller.NewMobileAuthController(authService)

	mbr := v1.Group("/auth")
	{
		mbr.POST("/login", ctlr.Login)
		mbr.GET("/refresh")

	}
}
