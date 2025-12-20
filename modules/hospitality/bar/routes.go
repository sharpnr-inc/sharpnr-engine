package bar

import (
	barcontrollers "sharpnr-engine/modules/hospitality/bar/controllers"
	barrepo "sharpnr-engine/modules/hospitality/bar/repo"
	barservices "sharpnr-engine/modules/hospitality/bar/services"
	kitchenrepo "sharpnr-engine/modules/hospitality/kitchen/kitchenRepo"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func BarRoutes(v1 *gin.RouterGroup, db *gorm.DB) {

	newBarRepo := barrepo.NewBarRepository(db)

	newKitchenRepo := kitchenrepo.NewKitchenRepository(db)
	newBarService := barservices.NewBarServices(newBarRepo, newKitchenRepo)
	newBarController := barcontrollers.NewBarController(newBarService)

	barRoute := v1.Group("/bar")

	barRoute.GET("/get-layout", newBarController.GetBarLayout)
	barRoute.POST("/add-floor", newBarController.AddFloorInBar)
	barRoute.POST("/add-table", newBarController.AddTableInBar)
	barRoute.POST("/update", newBarController.Update)
	barRoute.GET("/getBarConfigDetails", newBarController.GetBarDetailsConfig)

	barOrderRoutes := barRoute.Group("/session")
	{
		barOrderRoutes.POST("/create")
		barOrderRoutes.POST("/addOrders")
		barOrderRoutes.GET("/:id")
		barOrderRoutes.POST("/close")
	}
}
