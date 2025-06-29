// routes/v1/api.go
package v1

import (
	"github.com/gin-gonic/gin"
	orgcontrollers "github.com/sharpnr-inc/sharpnr-engine/controllers/organisationControllers"
	"github.com/sharpnr-inc/sharpnr-engine/repositories"
	orgservices "github.com/sharpnr-inc/sharpnr-engine/services/orgnisation"
	"gorm.io/gorm"
)

func RegisterV1Routes(router *gin.Engine, db *gorm.DB) {
	v1 := router.Group("/api/v1")

	// Set up repository, service, and controller
	orgRepo := repositories.NewOrganizationRepository(db)
	orgService := orgservices.NewOrganizationService(orgRepo)
	orgController := orgcontrollers.NewOrganizationController(orgService)

	// Public routes
	v1.POST("/organizations/register", orgController.Register)
	// v1.POST("/organizations/login", orgController.Login)
}
