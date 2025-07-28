package v1

import (
	orgcontrollers "sharpnr-engine/controllers/organisationControllers"
	repositories "sharpnr-engine/core/repositories/organization"
	services "sharpnr-engine/core/services/auth"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func CoreV1Routes(v1 *gin.RouterGroup, db *gorm.DB) {
	orgRepo := repositories.NewOrganizationRepository(db)
	orgService := services.CreateNewOrganization(orgRepo)
	orgController := orgcontrollers.NewOrganisationController(orgService)

	v1.POST("/organisation/register", orgController.Register)
}
