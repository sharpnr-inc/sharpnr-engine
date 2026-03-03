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

	mr.GET("", mc.GetMenu)
	mr.POST("/add", mc.AddMenu)
	mr.POST("/addCategory", mc.AddCategoryInMenu)
	mr.POST("/addSubCategory", mc.AddSubCategory)
	mr.POST("/add-menu-item", mc.AddMenuItems)
	mr.POST("/add-variant-group", mc.AddVariantsGroup)
	mr.POST("/add-variant", mc.AddVariants)
	mr.POST("/add-addOns", mc.AddAddons)

	mcr := mr.Group("/configs")
	{
		mcr.POST("/register/category-kitchen")
		mcr.POST("/register/sub-category-kitchen")
		mcr.POST("/register/menu-item-kitchen")
	}

}
