package v1

import (
	"github.com/gin-gonic/gin"
	orgcontrollers "github.com/sharpnr-inc/sharpnr-engine/controllers/organisationControllers"
	repositories "github.com/sharpnr-inc/sharpnr-engine/core/repositories/organization"
	services "github.com/sharpnr-inc/sharpnr-engine/core/services/auth"
	"gorm.io/gorm"
)

func CoreV1Routes(v1 *gin.RouterGroup, db *gorm.DB) {
	orgRepo := repositories.NewOrganizationRepository(db)
	orgService := services.CreateNewOrganization(orgRepo)
	orgController := orgcontrollers.NewOrganisationController(orgService)

	v1.POST("/organisation/register", orgController.Register)
}
