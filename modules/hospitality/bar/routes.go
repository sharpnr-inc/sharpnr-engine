package bar

import (
	barcontrollers "sharpnr-engine/modules/hospitality/bar/controllers"
	barrepo "sharpnr-engine/modules/hospitality/bar/repo"
	barservices "sharpnr-engine/modules/hospitality/bar/services"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func BarRoutes(v1 *gin.RouterGroup, db *gorm.DB) {

	newBarRepo := barrepo.NewBarRepository(db)
	newBarService := barservices.NewBarServices(newBarRepo)
	newBarController := barcontrollers.NewBarController(newBarService)

	barRoute := v1.Group("/bar")

	barRoute.GET("/get", newBarController.GetBar)
	barRoute.POST("/add-table")
	barRoute.POST("/update-table")
	barRoute.GET("/table/:tableId")
	barRoute.POST("/start-table-session")

	barRoute.POST("/update", newBarController.UpdateBarDetails)
	barRoute.GET("/getBarConfigDetails", newBarController.GetBarDetailsConfig)
}
