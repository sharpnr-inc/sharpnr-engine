package menu_routes

import (
	"sharpnr-engine/app"
	menu_controllers "sharpnr-engine/modules/hospitality/menu/controllers"

	"github.com/gin-gonic/gin"
)

func RegisterMenuRoutes(v1 *gin.RouterGroup, modules app.ModulesProvider) {

	ms := modules.Hospitality().MenuService()
	mc := menu_controllers.NewMenuController(ms)

	mr := v1.Group("/:ownersID/menu")

	mr.POST("/addCategory", mc.AddCategoryInMenu)
	mr.GET("/:categoryID/addSubCategory")
}
