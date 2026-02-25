package bar

import (
	"sharpnr-engine/app"
	barcontrollers "sharpnr-engine/modules/hospitality/bar/controllers"

	"github.com/gin-gonic/gin"
)

func BarRoutes(
	v1 *gin.RouterGroup,
	modules app.ModulesProvider,
) {

	barService := modules.Hospitality().BarService()
	newBarController := barcontrollers.NewBarController(barService)

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
